package service

import (
	"beiwanglu/model"
	"beiwanglu/pkg/util"
	"beiwanglu/serializer"
	"time"
)

// 创建任务的服务
type CreateTaskService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}
	code := 200
	err := model.DB.Create(&task).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = 400
		return serializer.Response{
			Status: code,
			Msg:    "失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "成功",
	}
}
