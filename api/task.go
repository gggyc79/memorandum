package api

import (
	"beiwanglu/pkg/util"
	"beiwanglu/service"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	createService := service.CreateTaskService{}
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err.Error())
		util.LogrusObj.Info(err)
	}
}
