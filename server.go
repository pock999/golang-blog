package main

import (
	"fmt"

	"blog/api/routers"

	orm "blog/api/config"
)

func main() {
	// 結束時，關閉資料庫連線
	defer orm.Eloquent.Close();

	fmt.Println("=== Main START ===");

	router := routers.InitRouter();
  router.Run(":8080")

}