package service

import (
	"awesomeProject/dao"
	_ "awesomeProject/docs"
	"awesomeProject/entity"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"strconv"
)

// CreateUser
// @Summary 创建用户
// @Tags 用户模块
// @Produce  json
// @Param name query string true "用户名" maxlength(100)
// @Param password query string true "密码" maxlength(100)
// @Param repassword query string true "确认密码" maxlength(100)
// @Success 200 {string} json{"code","message"}
// @Router /user/create_user [get]
func CreateUser(ctx *gin.Context) {
	user := entity.User{}
	name := ctx.Query("name")
	password := ctx.Query("password")
	repassword := ctx.Query("repassword")
	if password != repassword {
		ctx.JSON(-1, gin.H{"message": "两次密码不一致"})
		return
	}
	user.Name = name
	user.Password = password
	err := dao.Create_User(user).Error
	if err != nil {
		panic("创建用户失败,err=" + err.Error())
	}
	ctx.JSON(200, gin.H{"message": "新增用户成功"})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @Produce  json
// @Param id query uint64 true "id"
// @Param name query string true "用户名" maxlength(100)
// @Param password query string true "密码" maxlength(100)
// @Success 200 {string} json{"code","message"}
// @Failure	 400 {string} json{"code","message"}
// @Router /user/delete_user [post]
func DeleteUser(ctx *gin.Context) {
	id_str := ctx.Query("id")
	id, _ := strconv.Atoi(id_str)
	name := ctx.Query("name")
	password := ctx.Query("password")
	inUser := entity.User{Name: name, Password: password}
	var foundUser entity.User
	if err := dao.DB.Find(&foundUser, "id= ?", id).Error; err != nil {
		panic("删除用户失败,err==" + err.Error())
	}
	if userEqual(inUser, foundUser) {
		dao.DB.Delete(&foundUser)
		ctx.JSON(200, "删除用户成功")
	} else {
		ctx.JSON(400, "删除用户失败")
	}
}

func userEqual(user1 entity.User, user2 entity.User) bool {
	return user1.Name == user2.Name && user1.Password == user2.Password
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @Produce  json
// @Param id query uint64 true "id"
// @Param name query string true "name" maxlength(100)
// @Param password query string true "password" maxlength(100)
// @Param new_name query string true "new name" maxlength(100)
// @Param new_password query string true "new password" maxlength(100)
// @Success 200 {string} json{"code","message"}
// @Failure 400 {string} json{"code","message"}
// @Router /user/update_user [post]
func UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	name := ctx.Query("name")
	password := ctx.Query("password")
	new_name := ctx.Query("new_name")
	new_password := ctx.Query("new_password")
	in_user := entity.User{
		Name:     name,
		Password: password,
	}
	var found_user entity.User
	if err := dao.DB.Find(&found_user, "id=?", id).Error; err != nil {
		panic("更新模块查找用户信息失败")
	}
	if userEqual(found_user, in_user) {
		dao.DB.Model(&found_user).Updates(entity.User{Name: new_name, Password: new_password})
		ctx.JSON(200, "更新用户信息成功")
	} else {
		ctx.JSON(400, "更新用户信息失败")
	}
}
