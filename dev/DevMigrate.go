package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"

	"blog/api/models"
)

func main() {
	fmt.Println("=== DevMigate START ===");
	Eloquent, err := gorm.Open("mysql", "root:1qaz!QAZ@tcp(localhost:3306)/golang-blog?charset=utf8&parseTime=True&loc=Local");

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}

	// 結束前關閉
	defer Eloquent.Close();

	// 若存在就drop掉
	if Eloquent.HasTable(&models.User{}) {
		Eloquent.DropTable(&models.User{})
	}
	

	// migrate
	Eloquent.AutoMigrate(&models.User{});

	fmt.Println("=== DevMigate END ===");
}