package models

import (
	orm "blog/api/config"
)

type User struct {
	Id uint `gorm:"AUTO_INCREMENT"`
	Account string `json:"account"`
	Password string `json:"password"`
}


func (user *User) FindAll() (users []User, err error) {
	err = orm.Eloquent.Find(&users).Error;
	return users, err;
}