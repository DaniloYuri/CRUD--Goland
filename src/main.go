package src

import (
	"awesomeProject1/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) {
	router.GET("/search", controller.GetAllUser)
	router.POST("/save", controller.SaverUser)
	router.GET("/find/:userId", controller.GetUser)
	router.PUT("/update/:userId", controller.UpdateUser)
	router.DELETE("/delete/:userId", controller.DeleteUser)
}
