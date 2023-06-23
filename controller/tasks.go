package controller

import (
	"fmt"
	"log"
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
	// for _, v := range orderDateFilter {
	// 	fmt.Println(v)
	// }
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
func (c *Controller) HistoryUser(req *models.UserPrimaryKey) (*models.History, error) {
	// productsinfo := []int{}
	// product := make(map[string][]int)
	orderMap := make(map[string]interface{})

	// Order \\

	getorder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		log.Printf("Error while HistoryUser => GetByIdOrder: %+v", err)
		return nil, err
	}

	// User \\
	user := ""
	for _, v := range getorder.ShopCarts {
		if v.UserId == req.Id {
			if v.Status == true {
				getuser, err := c.UserGetById(&models.UserPrimaryKey{
					Id: v.UserId,
				})
				if err != nil {
					log.Printf("Error while getuser => UserGetById: %+v", err)
					return nil, err
				}
				user = getuser.Name
				getproduct, err := c.ProductGetById(&models.ProductPrimaryKey{
					Id: v.ProductId,
				})
				if err != nil {
					log.Printf("Error while getuser => UserGetById: %+v", err)
					return nil, err
				}
				orderMap["Product"] = getproduct.Name
				orderMap["Price"] = getproduct.Price
				orderMap["Count"] = v.Count
				orderMap["Time"] = v.Time

			}
		}
	}

	return &models.History{
		Name:  user,
		Order: orderMap,
	}, nil
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

// 6 - Task \\ Done. Work 50/50
// Top 10 ta sotilayotgan mahsulotlarni royxati.
func (c *Controller) TopProducts() ([]string, []int, error) {
	prodctsMap := make(map[string]int)
	productName := []string{}
	productCount := []int{}
	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, nil, err
	}

	for _, value := range getOrder.ShopCarts {
		if value.Status == true {
			prodctsMap[value.ProductId] += value.Count
		}
	}
	keys := make([]string, 0, len(prodctsMap))
	for key := range prodctsMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return prodctsMap[keys[i]] > prodctsMap[keys[j]]
	})

	for _, k := range keys {
		getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{
			Id: k,
		})
		if err != nil {
			return nil, nil, err
		}
		// fmt.Println(i, getProduct.Name, prodctsMap[k])
		productName = append(productName, getProduct.Name)
		productCount = append(productCount, prodctsMap[k])
	}

	return productName, productCount, nil
}

// 7 - Task \\ Done. Work 50/50
// Top 10 ta sotilayotgan mahsulotlarni royxati.
func (c *Controller) FailureProducts() ([]string, []int, error) {
	prodctsMap := make(map[string]int)
	productName := []string{}
	productCount := []int{}
	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, nil, err
	}

	for _, value := range getOrder.ShopCarts {
		if value.Status == true {
			prodctsMap[value.ProductId] += value.Count
		}
	}
	keys := make([]string, 0, len(prodctsMap))
	for key := range prodctsMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return prodctsMap[keys[i]] < prodctsMap[keys[j]]
	})

	for _, k := range keys {
		getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{
			Id: k,
		})
		if err != nil {
			return nil, nil, err
		}
		// fmt.Println(i, getProduct.Name, prodctsMap[k])
		productName = append(productName, getProduct.Name)
		productCount = append(productCount, prodctsMap[k])
	}

	return productName, productCount, nil
}

// 8 - Task \\ Done. Work 50/50
// Qaysi Sanada eng kop mahsulot sotilganligi boyicha jadval
func (c *Controller) TopTime() (*models.DateHistory, error) {
	getorder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	var (
		countcalc = 0
		data      = ""
		count     = []int{}
		date      = []string{}
	)

	for _, value := range getorder.ShopCarts {
		if value.Status == true {
			if value.Count >= countcalc {
				countcalc = value.Count
				data = value.Time
			}
		}
	}
	for _, value := range getorder.ShopCarts {
		if value.Status == true {
			if value.Count == countcalc {
				count = append(count, countcalc)
				date = append(date, data)
				// getuser, err := c.UserGetById(&models.UserPrimaryKey{Id: value.UserId})
				// if err != nil {
				// 	return nil, err
				// }
				// fmt.Println(getuser.Name)
			}
		}
	}
	result := models.DateHistory{
		Count: count,
		Date:  date,
	}

	return &result, nil
}

// 9 - Task \\ Not Done
// Qaysi category larda qancha mahsulot sotilgan boyicha jadval F
func (c *Controller) CategoryHistory() (map[string]int, error) {
	getorder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}
	category := make(map[string]int)
	for _, v := range getorder.ShopCarts {
		if v.Status == true {
			getproduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: v.ProductId})
			if err != nil {
				return nil, err
			}
			getcategory, err := c.CategoryGetList(&models.CategoryGetListRequest{})
			if err != nil {
				return nil, err
			}
			for _, value := range getcategory.Categories {
				fmt.Println(getproduct)
				if value.ParentId == getproduct.CategoryId {
					fmt.Println(value)
					getbyidP, err := c.CategoryGetById(&models.CategoryPrimaryKey{Id: value.ParentId})
					if err != nil {
						return nil, err
					}
					category[getbyidP.Name] += v.Count
				}
			}
		}
	}
	return category, nil
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
