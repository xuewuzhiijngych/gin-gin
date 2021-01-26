package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"ych.hub/common"
	"ych.hub/model"
	"ych.hub/utils"
)

// Register 用户注册
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(422, gin.H{
			"code": "-1",
			"msg":  "手机号格式不正确",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(422, gin.H{
			"code": "-1",
			"msg":  "密码长度不得少于六位",
		})
		return
	}
	// 没有传入用户名，给予一个10位的随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	// 判断手机号是否存在
	if telephoneExist(DB, telephone) {
		ctx.JSON(422, gin.H{
			"code": "-1",
			"msg":  "用户已存在",
		})
		return
	}

	// 创建用户
	user := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&user)

	// 返回结果
	ctx.JSON(200, gin.H{
		"code": "1",
		"msg":  "注册成功~",
	})
}

// telephoneExist 判断手机号是否存在
func telephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
