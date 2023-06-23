package models

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type UserCreate struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type UserUpdate struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type UserGetListRequest struct {
	Offset int
	Limit  int
}

type UserGetListResponse struct {
	Count int
	Users []*User
}
