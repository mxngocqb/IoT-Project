package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mxngocqb/IoT-Project/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(c *Controllers) *gin.Engine {
	router := gin.Default()

	drivers := router.Group("/drivers")
	{
		drivers.POST("", c.DriverController.CreateDriver)
		drivers.GET("", c.DriverController.ReadAllDriver)
		drivers.GET("/:driverID", c.DriverController.ReadDriverByID)
		drivers.PUT("/:driverID", c.DriverController.UpdateDriver)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
