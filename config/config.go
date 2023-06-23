package config

type Config struct {
	Path string

	UserFileName     string
	ProductFileName  string
	CategoryFileName string
	ShopCartFileName string
	BranchFileName   string
}

func Load() Config {
	cfg := Config{}

	cfg.Path = "./data"

	cfg.UserFileName = "/user.json"

	cfg.ProductFileName = "/product.json"

	cfg.CategoryFileName = "/category.json"

	cfg.ShopCartFileName = "/shop_cart.json"

	cfg.BranchFileName = "/branch.json"

	return cfg
}
