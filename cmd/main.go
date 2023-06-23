package main

import (
	"fmt"

	"market/config"
	"market/controller"
	"market/models"
	"market/storage/jsondb"
)

func main() {
	cfg := config.Load()
	strg, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)

	// s, err := con.Bonus(&models.UserPrimaryKey{Id: "080b6453-d424-4362-b1e0-18b80caa6fce"})
	// s, err := con.ProductCountSold()
	// if err != nil {
	// 	return
	// }
	// fmt.Println(s)
	// con.ShopCartGetList(&models.ShopCartGetListRequest{})
	// con.SortedShopCartGetList("2022-03-12 12:01:44", "2023-03-12 12:01:44", &models.ShopCartGetListRequest{})
	// f := "2022-03-12 12:01:44"
	// t := "2022-03-12 12:01:44"

	// fmt.Println(strconv.Atoi(f))
	// fmt.Println(strconv.Atoi(t))

	// s, err := con.SortedShopCartGetList(&models.ShopCartGetListRequest{})
	// if err != nil {
	// 	return
	// }
	// for _, v := range s.ShopCarts {
	// 	fmt.Println(v)
	// }

	// s, err := con.Filter(&models.ShopCartGetListRequest{Offset: 0, Limit: 0, FromTime: "2022-09-07 20:16:58", ToTime: "2023-09-07 20:16:58"})
	// for _, v := range s {
	// 	fmt.Println(v)
	// }

	s, err := con.Sort(&models.ShopCartGetListRequest{})
	for _, v := range s.ShopCarts{
		fmt.Println(v)
	}

}
