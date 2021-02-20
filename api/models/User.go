package models

import (
	orm "blog/api/config"
)

type User struct {
	Id uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Account string `json:"account"`
	Password string `json:"password"`
	Articles []Article `gorm:"foreignKey:UserId"`
}


func (user *User) FindAll() (users []User, err error) {
	err = orm.Eloquent.Find(&users).Error;
	return users, err;
}