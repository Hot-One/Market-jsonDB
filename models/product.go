package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type Product struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId string `jsom:"category_id"`
}

type ProductCreate struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId string `jsom:"category_id"`
}

type ProductUpdate struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId string `jsom:"category_id"`
}

type ProductGetListRequest struct {
	Offset int
	Limit  int
}

type ProductGetListResponse struct {
	Count    int
	Products []*Product
}
