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

type ProductRepo struct {
	fileName string
	file     *os.File
}

func NewProductRepo(fileName string, file *os.File) *ProductRepo {
	return &ProductRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *ProductRepo) Create(req *models.ProductCreate) (*models.Product, error) {
	// Read File \\

	products, err := u.read()
	if err != nil {
		return nil, err
	}

	// Create Model of Product \\
	var (
		id      = uuid.New().String()
		product = models.Product{
			Id:         id,
			Name:       req.Name,
			Price:      req.Price,
			CategoryId: req.CategoryId,
		}
	)

	products[id] = product

	// Write \\

	err = u.write(products)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (u *ProductRepo) GetById(req *models.ProductPrimaryKey) (*models.Product, error) {
	// Read File \\

	products, err := u.read()

	if err != nil {
		return nil, err
	}

	// Check Product Exist \\

	if _, have := products[req.Id]; !have {
		return nil, errors.New("Product not found")
	}

	// Get By Id \\

	product := products[req.Id]

	return &product, nil
}

func (u *ProductRepo) GetList(req *models.ProductGetListRequest) (*models.ProductGetListResponse, error) {

	var resp = &models.ProductGetListResponse{}
	resp.Products = []*models.Product{}

	// Read File \\

	ProductMap, err := u.read()
	if err != nil {
		return nil, err
	}

	// Fill the resp  \\

	resp.Count = len(ProductMap)
	for _, val := range ProductMap {
		Products := val
		resp.Products = append(resp.Products, &Products)
	}

	return resp, nil
}

func (u *ProductRepo) Update(req *models.ProductUpdate) (*models.Product, error) {
	// Read File \\

	Products, err := u.read()

	if err != nil {
		return nil, err
	}

	// Check Product Exist \\

	if _, ok := Products[req.Id]; !ok {
		return nil, errors.New("Product not found")
	}

	// Update Product \\

	Products[req.Id] = models.Product{
		Id:         req.Id,
		Name:       req.Name,
		Price:      req.Price,
		CategoryId: req.CategoryId,
	}

	// Write Update Product \\

	err = u.write(Products)
	if err != nil {
		return nil, err
	}
	Product := Products[req.Id]

	return &Product, nil
}

func (u *ProductRepo) Delete(req *models.ProductPrimaryKey) error {

	// Read File \\

	products, err := u.read()

	if err != nil {
		return err
	}

	// Delete Product \\

	delete(products, req.Id)

	// Write Product \\

	err = u.write(products)
	if err != nil {
		return err
	}

	return nil
}

func (u *ProductRepo) read() (map[string]models.Product, error) {
	var (
		Products   []*models.Product
		ProductMap = make(map[string]models.Product)
	)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &Products)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, Product := range Products {
		ProductMap[Product.Id] = *Product
	}

	return ProductMap, nil
}

func (u *ProductRepo) write(ProductMap map[string]models.Product) error {
	var Products []models.Product

	for _, value := range ProductMap {
		Products = append(Products, value)
	}

	body, err := json.MarshalIndent(Products, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
