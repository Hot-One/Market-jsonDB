package models

type BranchPrimaryKey struct {
	Id string `json:"id"`
}

type Branch struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type BranchCreate struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type BranchUpdate struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type BranchGetListRequest struct {
	Offset int
	Limit  int
}

type BranchGetListResponse struct {
	Count    int
	Branches []*Branch
}
