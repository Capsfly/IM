package service

import (
	"awesomeProject/dao"
	_ "awesomeProject/docs"
	"awesomeProject/entity"
	"errors"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// CreateUser
// @Summary 创建用户
// @Tags 用户模块
// @Produce  json
// @Param UID query string true "UID" maxlength(100)
// @Param name query string true "用户名" maxlength(100)
// @Param password query string true "密码" maxlength(100)
// @Param repassword query string true "确认密码" maxlength(100)
// @Param email query string true "email" maxlength(100)
// @Param phone query string true "phone" maxlength(100)
// @Success 200 {string} json{"code","message"}
// @Failure 400 {string} json{"code","message"}
// @Router /user/create_user [post]
func CreateUser(ctx *gin.Context) {
	uid := ctx.Query("UID")
	name := ctx.Query("name")
	password := ctx.Query("password")
	repassword := ctx.Query("repassword")
	email := ctx.Query("email")
	phone := ctx.Query("phone")
	if password != repassword {
		ctx.JSON(400, gin.H{"message": "两次密码不一致"})
		return
	}
	user := entity.User{
		UID:      uid,
		Name:     name,
		Password: password,
		Email:    email,
		Phone:    phone,
	}
	var foundUser entity.User
	dao.DB.Where("UID=?", uid).First(&foundUser)
	if foundUser.UID == uid {
		ctx.JSON(400, "UID重复，请重试")
		return
	}

	requestValid, _ := govalidator.ValidateStruct(user)
	if !requestValid {
		ctx.JSON(400, "邮箱或者手机号不合法，请检查输入")
		return
	}

	if err := dao.DB.Model(&entity.User{}).Create(&user).Error; err != nil {
		panic("创建用户失败,err==" + err.Error())
	}
	ctx.JSON(200, gin.H{"message": "新增用户成功"})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @Produce  json
// @Param UID query string true "UID"
// @Param name query string true "用户名" maxlength(100)
// @Param password query string true "密码" maxlength(100)
// @Success 200 {string} json{"code","message"}
// @Failure	 400 {string} json{"code","message"}
// @Router /user/delete_user [post]
func DeleteUser(ctx *gin.Context) {
	uid := ctx.Query("UID")
	name := ctx.Query("name")
	password := ctx.Query("password")

	var user entity.User
	if err := dao.DB.Model(&entity.User{}).Where("UID=? and name=? and password=?", uid, name, password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(400, "删除的用户没找到")
		}
		panic("删除模块查找用户失败,err=" + err.Error())
	}
	dao.DB.Delete(&user)
	ctx.JSON(200, "删除用户成功")
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @Produce  json
// @Param UID query string true "UID"
// @Param name query string true "name" maxlength(100)
// @Param password query string true "password" maxlength(100)
// @Param new_name query string true "new name" maxlength(100)
// @Param new_password query string true "new password" maxlength(100)
// @Success 200 {string} json{"code","message"}
// @Failure 400 {string} json{"code","message"}
// @Router /user/update_user [post]
func UpdateUser(ctx *gin.Context) {
	uid := ctx.Query("UID")
	name := ctx.Query("name")
	password := ctx.Query("password")

	newName := ctx.Query("new_name")
	newPassword := ctx.Query("new_password")
	newUser := entity.User{
		Name:     newName,
		Password: newPassword,
	}
	var oldUser entity.User
	if err := dao.DB.Where("UID=? and name=? and password=?", uid, name, password).First(&oldUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(400, "想要修改用户的信息不存在")
			return
		}
		panic("更新模块查找用户失败,err=" + err.Error())
	}

	if err := dao.DB.Model(&oldUser).Updates(newUser).Error; err != nil {
		panic("更新模块查找用户失败,err=" + err.Error())
	}
	ctx.JSON(200, "更新用户信息成功")
}
