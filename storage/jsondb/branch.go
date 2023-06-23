package jsondb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"

	"market/models"
)

type BranchRepo struct {
	fileName string
	file     *os.File
}

func NewBranchRepo(fileName string, file *os.File) *BranchRepo {
	return &BranchRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *BranchRepo) Create(req *models.BranchCreate) (*models.Branch, error) {
	// Read File \\

	Categories, err := u.read()
	if err != nil {
		return nil, err
	}

	// Create Model of Branch \\
	var (
		id     = uuid.New().String()
		Branch = models.Branch{
			Id:   id,
			Name: req.Name,
		}
	)

	Categories[id] = Branch

	// Write \\

	err = u.write(Categories)
	if err != nil {
		return nil, err
	}

	return &Branch, nil
}

func (u *BranchRepo) GetById(req *models.BranchPrimaryKey) (*models.Branch, error) {
	// Read File \\
	fmt.Println(req)
	Categories, err := u.read()

	if err != nil {
		return nil, err
	}

	// Check Branch Exist \\

	if _, have := Categories[req.Id]; !have {
		return nil, errors.New("Branch not found")
	}

	// Get By Id \\

	Branch := Categories[req.Id]

	return &Branch, nil
}

func (u *BranchRepo) GetList(req *models.BranchGetListRequest) (*models.BranchGetListResponse, error) {

	var resp = &models.BranchGetListResponse{}
	resp.Branches = []*models.Branch{}

	// Read File \\

	BranchMap, err := u.read()
	if err != nil {
		return nil, err
	}

	// Fill the resp  \\

	resp.Count = len(BranchMap)
	for _, val := range BranchMap {
		Branches := val
		resp.Branches = append(resp.Branches, &Branches)
	}

	return resp, nil
}

func (u *BranchRepo) Update(req *models.BranchUpdate) (*models.Branch, error) {
	// Read File \\

	Categories, err := u.read()

	if err != nil {
		return nil, err
	}

	// Check Branch Exist \\

	if _, ok := Categories[req.Id]; !ok {
		return nil, errors.New("Branch not found")
	}

	// Update Branch \\

	Categories[req.Id] = models.Branch{
		Id:   req.Id,
		Name: req.Name,
	}

	// Write Update Branch \\

	err = u.write(Categories)
	if err != nil {
		return nil, err
	}
	Branch := Categories[req.Id]

	return &Branch, nil
}

func (u *BranchRepo) Delete(req *models.BranchPrimaryKey) error {

	// Read File \\

	Categories, err := u.read()

	if err != nil {
		return err
	}

	// Delete Branch \\

	delete(Categories, req.Id)

	// Write Branch \\

	err = u.write(Categories)
	if err != nil {
		return err
	}

	return nil
}

func (u *BranchRepo) read() (map[string]models.Branch, error) {
	var (
		Categories []*models.Branch
		BranchMap  = make(map[string]models.Branch)
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

	for _, Branch := range Categories {
		BranchMap[Branch.Id] = *Branch
	}

	return BranchMap, nil
}

func (u *BranchRepo) write(BranchMap map[string]models.Branch) error {
	var Categories []models.Branch

	for _, value := range BranchMap {
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
