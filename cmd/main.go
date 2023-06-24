package main

import (
	"fmt"

	"market/config"
	"market/controller"
	"market/storage/jsondb"
)

func main() {
	cfg := config.Load()
	strg, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)

	// 1 - Task \\

	// s, err := con.Sort(&models.ShopCartGetListRequest{})
	// for _, v := range s.ShopCarts {
	// 	fmt.Println(v)
	// }

	// 2 - Task \\

	// s, err := con.Filter(&models.ShopCartGetListRequest{Offset: 0, Limit: 0, FromTime: "2022-09-07 20:16:58", ToTime: "2023-09-07 20:16:58"})
	// for _, v := range s {
	// 	fmt.Println(v)
	// }

	// 3 - Task \\

	// s, err := con.HistoryUser(&models.UserPrimaryKey{"27457ac2-74dd-4656-b9b0-0d46b1af10dc"})
	// if err != nil {
	// 	return
	// }
	// fmt.Println(s)
	// for k, v := range s {
	// 	for key, val := range v {
	// 		fmt.Println(k)
	// 		fmt.Println(key, val)
	// 	}
	// }

	// 6 - Task  \\

	// s, err := con.TopProducts()
	// for _, v := range s {
	// 	fmt.Println(v)
	// }

	// 8 - Task \\

	// s, err := con.TopTime()
	// if err != nil{
	// 	return
	// }
	// for _, v := range s{
	// 	fmt.Println(v)
	// }

	// 9 - Task \\
	s, err := con.CategoryHistory()
	if err != nil {
		return
	}
	for _, v := range s {
		fmt.Println(v)
	}
}
