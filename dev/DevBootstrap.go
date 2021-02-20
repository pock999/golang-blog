package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"

	"blog/api/models"
)

func main() {
	fmt.Println("=== DevBootstrap START ===");

	Eloquent, err := gorm.Open("mysql", "root:1qaz!QAZ@tcp(localhost:3306)/golang-blog?charset=utf8&parseTime=True&loc=Local");
	
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}

	// 結束前關閉
	defer Eloquent.Close();

	var user models.User;

	user.Account = "account0001";
	user.Password = "password001";

	result := Eloquent.Create(&user)

	if result.Error != nil {
		err = result.Error
		fmt.Println(err);
		return
	}

	fmt.Println("=== DevBootstrap END ===");

	return
}