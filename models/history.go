package models

type History struct {
	ProductName string
	Count       int
	Total       int
	Time        string
}

type ProductsHistory struct {
	Name  string
	Count int
}

type DateHistory struct {
	Date  string
	Count int
}

type CategoryHistory struct {
	Name  string
	Count int
}
