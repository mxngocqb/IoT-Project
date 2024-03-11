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
		drivers.DELETE("/:driverID", c.DriverController.DeleteDriver)
	}

	vehicles := router.Group("/vehicles")
	{
		vehicles.POST("", c.VehicleController.CreateVehicle)
		vehicles.GET("", c.VehicleController.ReadAllVehicle)
		vehicles.GET("/:vehicleID", c.VehicleController.ReadVehicleByID)
		vehicles.PUT("/:vehicleID", c.VehicleController.UpdateVehicle)
		vehicles.DELETE("/:vehicleID", c.VehicleController.DeleteVehicle)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
