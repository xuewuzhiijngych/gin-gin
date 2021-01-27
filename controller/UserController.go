package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"ych.hub/common"
	"ych.hub/dto"
	"ych.hub/model"
	"ych.hub/response"
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
		response.Fail(ctx, nil, "手机号格式不正确")
		return
	}
	if len(password) < 6 {
		response.Fail(ctx, nil, "密码长度不得少于六位")
		return
	}
	// 没有传入用户名，给予一个10位的随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	// 判断手机号是否存在
	if telephoneExist(DB, telephone) {
		response.Fail(ctx, nil, "用户已存在")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(ctx, nil, "系统异常")
		return
	}
	// 创建用户
	user := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&user)

	// 返回结果
	response.Success(ctx, nil, "注册成功~")

}

// Login 用户登录
func Login(ctx *gin.Context) {
	DB := common.GetDB()

	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		response.Fail(ctx, nil, "手机号格式不正确")
		return
	}
	if len(password) < 6 {
		response.Fail(ctx, nil, "密码长度不得少于六位")
		return
	}

	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)

	if user.ID == 0 {
		response.Fail(ctx, nil, "用户不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(ctx, nil, "密码错误")
		return
	}

	// 发放 token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Fail(ctx, nil, "系统异常")
		log.Printf("token 生成异常 %v", err)
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{
		"token": token,
	}, "登录成功")
	return
}

// Info 用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{
		"user": dto.ToUserDto(user.(model.User)),
	}, "")
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
