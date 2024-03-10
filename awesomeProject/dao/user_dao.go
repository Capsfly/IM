package dao

import (
	"awesomeProject/entity"
	"fmt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Create_User(user entity.User) *gorm.DB {
	fmt.Println("CreateUser")
	return DB.Create(&user)
}

func DeleteUser(user entity.User) *gorm.DB {
	return DB.Delete(&user)
}

func UpdateUser(user entity.User) *gorm.DB {
	return DB.Model(&user).Updates(entity.User{Name: user.Name, Password: user.Password})
}
