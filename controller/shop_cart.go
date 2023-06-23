package controller

import (
	"errors"
	"log"

	"market/models"
)

func (c *Controller) ShopCartCreate(req *models.ShopCartCreate) (*models.ShopCart, error) {

	log.Printf("ShopCart create req: %+v\n", req)

	resp, err := c.Strg.ShopCart().Create(req)
	if err != nil {
		log.Printf("error while ShopCart Create: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) ShopCartGetById(req *models.UserPrimaryKey) (*models.ShopCart, error) {

	resp, err := c.Strg.ShopCart().GetById(req)
	if err != nil {
		log.Printf("error while ShopCart GetById: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) ShopCartGetList(req *models.ShopCartGetListRequest) (*models.ShopCartGetListResponse, error) {

	resp, err := c.Strg.ShopCart().GetList(req)
	if err != nil {
		log.Printf("error while ShopCart GetList: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) SortedShopCartGetList(req *models.ShopCartGetListRequest) (*models.ShopCartGetListResponse, error) {

	resp, err := c.Strg.ShopCart().GetList(req)
	if err != nil {
		log.Printf("error while ShopCart GetList: %+v\n", err)
		return nil, err
	}

	return resp, nil
}
