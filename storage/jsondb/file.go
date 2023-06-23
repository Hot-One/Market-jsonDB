package jsondb

import (
	"os"

	"market/config"
	"market/storage"
)

type StoreJSON struct {
	user     *UserRepo
	product  *ProductRepo
	category *CategoryRepo
	shopcart *ShopCartRepo
	branch   *BranchRepo
}

func NewConnectionJSON(cfg *config.Config) (storage.StorageI, error) {

	userfile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	productfile, err := os.Open(cfg.Path + cfg.ProductFileName)
	if err != nil {
		return nil, err
	}

	categoryfile, err := os.Open(cfg.Path + cfg.CategoryFileName)
	if err != nil {
		return nil, err
	}

	shopcartfile, err := os.Open(cfg.Path + cfg.ShopCartFileName)
	if err != nil {
		return nil, err
	}

	branchfile, err := os.Open(cfg.Path + cfg.BranchFileName)
	if err != nil {
		return nil, err
	}

	return &StoreJSON{
		user:     NewUserRepo(cfg.Path+cfg.UserFileName, userfile),
		product:  NewProductRepo(cfg.Path+cfg.ProductFileName, productfile),
		category: NewCategoryRepo(cfg.Path+cfg.CategoryFileName, categoryfile),
		shopcart: NewShopCartRepo(cfg.Path+cfg.ShopCartFileName, shopcartfile),
		branch:   NewBranchRepo(cfg.Path+cfg.BranchFileName, branchfile),
	}, nil
}

func (u *StoreJSON) User() storage.UserRepoI {
	return u.user
}

func (p *StoreJSON) Product() storage.ProductRepoI {
	return p.product
}

func (c *StoreJSON) Category() storage.CategoryRepoI {
	return c.category
}

func (s *StoreJSON) ShopCart() storage.ShopCartRepoI {
	return s.shopcart
}

func (b *StoreJSON) Branch() storage.BranchRepoI {
	return b.branch
}
