package models

type History struct {
	Name  string
	Order map[string]interface{}
}

type DateHistory struct {
	Date  []string
	Count []int
}
