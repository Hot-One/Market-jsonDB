package models

type ShopCart struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}

type ShopCartCreate struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}

type ShopCartUpdate struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}

type ShopCartGetListRequest struct {
	Offset   int
	Limit    int
	FromTime string
	ToTime   string
}

type ShopCartGetListResponse struct {
	Count     int
	ShopCarts []*ShopCart
}
