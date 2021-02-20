package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

var Eloquent *gorm.DB

func init() {
	var err error;
	Eloquent, err = gorm.Open("mysql", "root:1qaz!QAZ@tcp(localhost:3306)/golang-blog?charset=utf8&parseTime=True&loc=Local");

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}