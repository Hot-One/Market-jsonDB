package storage

import "market/models"

type StorageI interface {
	User() UserRepoI
	Product() ProductRepoI
	Category() CategoryRepoI
	ShopCart() ShopCartRepoI
	Branch() BranchRepoI
}

type UserRepoI interface {
	Create(req *models.UserCreate) (*models.User, error)
	GetById(req *models.UserPrimaryKey) (*models.User, error)
	GetList(req *models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(req *models.UserUpdate) (*models.User, error)
	Delete(req *models.UserPrimaryKey) error
}

type ProductRepoI interface {
	Create(req *models.ProductCreate) (*models.Product, error)
	GetById(req *models.ProductPrimaryKey) (*models.Product, error)
	GetList(req *models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(req *models.ProductUpdate) (*models.Product, error)
	Delete(req *models.ProductPrimaryKey) error
}

type CategoryRepoI interface {
	Create(req *models.CategoryCreate) (*models.Category, error)
	GetById(req *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(req *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error)
	Update(req *models.CategoryUpdate) (*models.Category, error)
	Delete(req *models.CategoryPrimaryKey) error
}

type ShopCartRepoI interface {
	Create(req *models.ShopCartCreate) (*models.ShopCart, error)
	GetById(req *models.UserPrimaryKey) (*models.ShopCart, error)
	GetList(req *models.ShopCartGetListRequest) (*models.ShopCartGetListResponse, error)
	SortedGetList(req *models.ShopCartGetListRequest) (*models.ShopCartGetListResponse, error)
}

type BranchRepoI interface {
	Create(req *models.BranchCreate) (*models.Branch, error)
	GetById(req *models.BranchPrimaryKey) (*models.Branch, error)
	GetList(req *models.BranchGetListRequest) (*models.BranchGetListResponse, error)
	Update(req *models.BranchUpdate) (*models.Branch, error)
	Delete(req *models.BranchPrimaryKey) error
}
