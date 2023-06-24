package controller

import (
	"fmt"
	"sort"

	"market/models"
)

// 1 - Task \\ Done. Work Fully
// Shop cart boyicha default holati time sort bolishi kerak. DESC
func (c *Controller) Sort(req *models.ShopCartGetListRequest) (*models.ShopCartGetListResponse, error) {
	var resp = &models.ShopCartGetListResponse{}
	var orderDateFilter []*models.ShopCart
	getorder, err := c.ShopCartGetList(req)
	if err != nil {
		return nil, err
	}
	for _, ord := range getorder.ShopCarts {
		orderDateFilter = append(orderDateFilter, ord)

	}
	sort.Slice(orderDateFilter, func(i, j int) bool {
		return orderDateFilter[i].Time > orderDateFilter[j].Time
	})
	resp.Count = len(orderDateFilter)
	resp.ShopCarts = orderDateFilter
	return resp, nil
}

// 2 - Task \\ Done. Work Fully
func (c *Controller) Filter(req *models.ShopCartGetListRequest) ([]*models.ShopCart, error) {
	var orderDateFilter []*models.ShopCart
	getorder, err := c.ShopCartGetList(req)
	if err != nil {
		return nil, err
	}
	for _, ord := range getorder.ShopCarts {
		if ord.Time >= req.FromTime && ord.Time < req.ToTime {
			orderDateFilter = append(orderDateFilter, ord)
		}
	}

	return orderDateFilter, nil
}

// 3 Task \\ Done. Work Fully
// Client history chiqish kerak. Ya'ni sotib olgan mahsulotlari korsatish kerak \\
func (c *Controller) HistoryUser(req *models.UserPrimaryKey) (map[string][]models.History, error) {
	var (
		orders   = []models.History{}
		orderMap = make(map[string][]models.History)
	)
	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	getUser, err := c.UserGetById(&models.UserPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}

	for _, v := range getOrder.ShopCarts {
		getproduct, err := c.ProductGetById(&models.ProductPrimaryKey{v.ProductId})
		if err != nil {
			return nil, err
		}

		if v.UserId == req.Id {
			if v.Status == true {
				orders = append(orders, models.History{
					ProductName: getproduct.Name,
					Count:       v.Count,
					Total:       v.Count * getproduct.Price,
					Time:        v.Time,
				})
			}
		}
	}
	orderMap[getUser.Name] = orders
	return orderMap, nil
}

// 4 - Task \\ Done. Work Fully
// Client qancha pul mahsulot sotib olganligi haqida hisobot.
func (c *Controller) UserCash(req *models.UserPrimaryKey) (map[string]int, error) {
	user := make(map[string]int)

	getorder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	getuser, err := c.UserGetById(req)

	for _, value := range getorder.ShopCarts {
		if value.UserId == req.Id {
			if value.Status == true {
				getproduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
				if err != nil {
					return nil, err
				}
				user[getuser.Name] += value.Count * getproduct.Price
			}
		}
	}
	return user, nil
}

// 5 - Task \\ Done. Work Fully
// Productlarni Qancha sotilgan boyicha hisobot
func (c *Controller) ProductCountSold() (map[string]int, error) {
	product := make(map[string]int)

	getorder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	for _, value := range getorder.ShopCarts {
		getproduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
		if err != nil {
			return nil, err
		}
		if value.Status == true {
			product[getproduct.Name] += value.Count
		}

	}
	return product, nil
}

// 6 - Task \\ Done. Work Fully
// Top 10 ta sotilayotgan mahsulotlarni royxati.
func (c *Controller) TopProducts() ([]*models.ProductsHistory, error) {
	var (
		prodctsMap = make(map[string]int)
		products   []*models.ProductsHistory
	)

	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	for _, value := range getOrder.ShopCarts {
		getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
		if err != nil {
			return nil, err
		}
		if value.Status == true {
			prodctsMap[getProduct.Name] += value.Count
		}
	}
	for k, v := range prodctsMap {
		products = append(products, &models.ProductsHistory{
			Name:  k,
			Count: v,
		})
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Count > products[j].Count
	})

	return products, nil
}

// 7 - Task \\ Done. Work Fully
// Top 10 ta sotilayotgan mahsulotlarni royxati.
func (c *Controller) FailureProducts() ([]*models.ProductsHistory, error) {
	var (
		prodctsMap = make(map[string]int)
		products   []*models.ProductsHistory
	)

	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	for _, value := range getOrder.ShopCarts {
		getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
		if err != nil {
			return nil, err
		}
		if value.Status == true {
			prodctsMap[getProduct.Name] += value.Count
		}
	}
	for k, v := range prodctsMap {
		products = append(products, &models.ProductsHistory{
			Name:  k,
			Count: v,
		})
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Count < products[j].Count
	})

	return products, nil
}

// 8 - Task \\ Done. Fully
// Qaysi Sanada eng kop mahsulot sotilganligi boyicha jadval
func (c *Controller) TopTime() ([]*models.DateHistory, error) {
	var (
		toptimes = make(map[string]int)
		result   []*models.DateHistory
	)

	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	for _, value := range getOrder.ShopCarts {
		if value.Status == true {
			toptimes[value.Time] += value.Count
		}
	}

	for k, v := range toptimes {
		result = append(result, &models.DateHistory{
			Date:  k,
			Count: v,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	return result, nil
}

// 9 - Task \\ Not Done
// Qaysi category larda qancha mahsulot sotilgan boyicha jadval F
func (c *Controller) CategoryHistory() ([]*models.CategoryHistory, error) {
	// var (
	// 	categoryMap = make(map[string]int)
	// 	category    []*models.CategoryHistory
	// )

	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}
	for _, value := range getOrder.ShopCarts {
		getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
		if err != nil {
			return nil, err
		}
		fmt.Println(getProduct.CategoryId)
	}
	return nil, nil
}

// 10 - Task \\ Done. Work Fully
// Qaysi Client eng Active xaridor. Bitta ma'lumot chiqsa yetarli.
func (c *Controller) ActiveUser() (string, error) {
	users := make(map[string]int)
	getorder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return "", err
	}
	for _, value := range getorder.ShopCarts {
		if value.Status == true {
			getproduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
			if err != nil {
				return "", err
			}
			users[value.UserId] += value.Count * getproduct.Price
		}
	}
	user, sum := "", 0
	for key, value := range users {
		if sum < value {
			user = key
			sum = value
		}
	}
	getuser, err := c.UserGetById(&models.UserPrimaryKey{
		Id: user,
	})
	if err != nil {
		return "", err
	}
	return getuser.Name, nil
}

// 11 - Task \\ Done. Work Fully
// Agar client 9 dan katta mahuslot sotib olgan bolsa,
// 1 tasi tekinga beriladi va 9 ta uchun pul hisoblanadi.
// 1 tasi eng arzon mahsulotni pulini hisoblamaysiz.
// Yangi korzinka qoshib tekshirib koring.
func (c *Controller) Bonus(req *models.UserPrimaryKey) (int, error) {
	getorder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return 0, err
	}
	getuser, err := c.UserGetById(&models.UserPrimaryKey{Id: req.Id})
	if err != nil {
		return 0, err
	}
	ProductPrices := []int{}
	sum := 0
	for _, value := range getorder.ShopCarts {
		if getuser.Id == value.UserId {
			if value.Status == true {
				getproduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
				if err != nil {
					return 0, err
				}
				if value.Count > 9 {
					sum = value.Count * getproduct.Price
					ProductPrices = append(ProductPrices, getproduct.Price)
				}
			}
		}
	}
	sort.Ints(ProductPrices)
	sum -= ProductPrices[0]
	return sum, nil
}
