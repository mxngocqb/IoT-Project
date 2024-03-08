package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mxngocqb/IoT-Project/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(c *Controllers) *gin.Engine {
	router := gin.Default()

	customers := router.Group("/drivers")
	{
		customers.POST("", c.DriverController.CreateDriver)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
