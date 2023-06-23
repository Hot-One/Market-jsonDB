package jsondb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"market/models"
)

type ShopCartRepo struct {
	fileName string
	file     *os.File
}

func NewShopCartRepo(fileName string, file *os.File) *ShopCartRepo {
	return &ShopCartRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *ShopCartRepo) Create(req *models.ShopCartCreate) (*models.ShopCart, error) {
	// Read File \\

	ShopCarts, err := u.read()
	if err != nil {
		return nil, err
	}

	// Create Model of ShopCart \\
	var (
		ShopCart = models.ShopCart{
			ProductId: req.ProductId,
			UserId:    req.UserId,
			Count:     req.Count,
			Status:    req.Status,
			Time:      time.Now().Format("2006-01-02 15:04:05"),
		}
	)

	// Write \\
	ShopCarts[req.UserId] = ShopCart
	err = u.write(ShopCarts)
	if err != nil {
		return nil, err
	}

	return &ShopCart, nil
}

func (u *ShopCartRepo) GetById(req *models.UserPrimaryKey) (*models.ShopCart, error) {
	// Read File \\

	shopcarts, err := u.read()

	if err != nil {
		return nil, err
	}

	// Check ShopCart Exist \\

	if _, have := shopcarts[req.Id]; !have {
		return nil, errors.New("ShopCart not found")
	}

	// Get By Id \\

	shopcart := shopcarts[req.Id]

	return &shopcart, nil
}

func (u *ShopCartRepo) SortedGetList(req *models.ShopCartGetListRequest) (*models.ShopCartGetListResponse, error) {
	var resp = &models.ShopCartGetListResponse{}
	resp.ShopCarts = []*models.ShopCart{}

	ShopCartMap, err := u.read()
	if err != nil {
		return nil, err
	}

	resp.Count = len(ShopCartMap)

	for _, value := range ShopCartMap {
		shopcarts := value
		resp.ShopCarts = append(resp.ShopCarts, &shopcarts)
	}

	// sort.Slice(ShopCartMap, func(i, j int) bool {
	// 	return ShopCartMap.[i] > ShopCartMap[j]

	// })
	return resp, nil
}

func (u *ShopCartRepo) GetList(req *models.ShopCartGetListRequest) (*models.ShopCartGetListResponse, error) {
	var resp = &models.ShopCartGetListResponse{}
	resp.ShopCarts = []*models.ShopCart{}

	// Read File \\

	ShopCartMap, err := u.read()
	if err != nil {
		return nil, err
	}

	// Fill the resp  \\

	resp.Count = len(ShopCartMap)
	for _, val := range ShopCartMap {
		shopcarts := val
		resp.ShopCarts = append(resp.ShopCarts, &shopcarts)
	}

	return resp, nil
}

func (u *ShopCartRepo) read() (map[string]models.ShopCart, error) {
	var (
		ShopCarts   []*models.ShopCart
		ShopCartMap = make(map[string]models.ShopCart)
	)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &ShopCarts)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, ShopCart := range ShopCarts {
		ShopCartMap[ShopCart.UserId] = *ShopCart
	}

	return ShopCartMap, nil
}

func (u *ShopCartRepo) write(ShopCartMap map[string]models.ShopCart) error {
	var ShopCarts []models.ShopCart

	for _, value := range ShopCartMap {
		ShopCarts = append(ShopCarts, value)
	}

	body, err := json.MarshalIndent(ShopCarts, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
