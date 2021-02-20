package models

import (
	// orm "blog/api/config"
)

type Article struct {
	Id uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Title string `json:"title"`
	Content string `json:"content"`
	UserId uint `gorm:"column:UserId"`
}