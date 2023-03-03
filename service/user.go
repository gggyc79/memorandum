package service

import (
	"beiwanglu/model"
	"beiwanglu/pkg/util"
	"beiwanglu/serializer"

	"gorm.io/gorm"
)

// UserService 用户注册服务
type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

func (service *UserService) Register() *serializer.Response {
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	// 表单验证
	if count == 1 {

		return &serializer.Response{
			Status: 400,
			Msg:    "已有该账号",
		}
	}
	user.UserName = service.UserName
	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		code := 400
		return &serializer.Response{
			Status: code,
			Msg:    err.Error(),
		}
	}
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {

		return &serializer.Response{
			Status: 500,
			Msg:    "数据库问题",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg:    "成功",
	}
}

// Login 用户登陆函数
func (service *UserService) Login() serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		// 如果查询不到，返回相应的错误
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return serializer.Response{
				Status: 400,
				Msg:    "密码错误",
			}
		}
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	if !user.CheckPassword(service.Password) {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "yes",
	}
}
