package jsondb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"

	"market/models"
)

type CategoryRepo struct {
	fileName string
	file     *os.File
}

func NewCategoryRepo(fileName string, file *os.File) *CategoryRepo {
	return &CategoryRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *CategoryRepo) Create(req *models.CategoryCreate) (*models.Category, error) {
	// Read File \\

	Categories, err := u.read()
	if err != nil {
		return nil, err
	}

	// Create Model of Category \\
	var (
		id       = uuid.New().String()
		Category = models.Category{
			Id:       id,
			Name:     req.Name,
			ParentId: req.ParentId,
		}
	)

	Categories[id] = Category

	// Write \\

	err = u.write(Categories)
	if err != nil {
		return nil, err
	}

	return &Category, nil
}

func (u *CategoryRepo) GetById(req *models.CategoryPrimaryKey) (*models.Category, error) {
	// Read File \\
	Categories, err := u.read()

	if err != nil {
		return nil, err
	}

	// Check Category Exist \\

	if _, have := Categories[req.Id]; !have {
		return nil, errors.New("Category not found")
	}

	// Get By Id \\

	Category := Categories[req.Id]

	return &Category, nil
}

func (u *CategoryRepo) GetList(req *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error) {

	var resp = &models.CategoryGetListResponse{}
	resp.Categories = []*models.Category{}

	// Read File \\

	CategoryMap, err := u.read()
	if err != nil {
		return nil, err
	}

	// Fill the resp  \\

	resp.Count = len(CategoryMap)
	for _, val := range CategoryMap {
		Categories := val
		resp.Categories = append(resp.Categories, &Categories)
	}

	return resp, nil
}

func (u *CategoryRepo) Update(req *models.CategoryUpdate) (*models.Category, error) {
	// Read File \\

	Categories, err := u.read()

	if err != nil {
		return nil, err
	}

	// Check Category Exist \\

	if _, ok := Categories[req.Id]; !ok {
		return nil, errors.New("Category not found")
	}

	// Update Category \\

	Categories[req.Id] = models.Category{
		Id:       req.Id,
		Name:     req.Name,
		ParentId: req.ParentId,
	}

	// Write Update Category \\

	err = u.write(Categories)
	if err != nil {
		return nil, err
	}
	Category := Categories[req.Id]

	return &Category, nil
}

func (u *CategoryRepo) Delete(req *models.CategoryPrimaryKey) error {

	// Read File \\

	Categories, err := u.read()

	if err != nil {
		return err
	}

	// Delete Category \\

	delete(Categories, req.Id)

	// Write Category \\

	err = u.write(Categories)
	if err != nil {
		return err
	}

	return nil
}

func (u *CategoryRepo) read() (map[string]models.Category, error) {
	var (
		Categories  []*models.Category
		CategoryMap = make(map[string]models.Category)
	)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &Categories)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, Category := range Categories {
		CategoryMap[Category.Id] = *Category
	}

	return CategoryMap, nil
}

func (u *CategoryRepo) write(CategoryMap map[string]models.Category) error {
	var Categories []models.Category

	for _, value := range CategoryMap {
		Categories = append(Categories, value)
	}

	body, err := json.MarshalIndent(Categories, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
