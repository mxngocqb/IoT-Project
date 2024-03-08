package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mxngocqb/IoT-Project/model"
	"github.com/mxngocqb/IoT-Project/service/driver"
	"github.com/rs/zerolog/log"
)

type DriverController struct {
	driverService driver.DriverService
}

func NewDriverController(driverService driver.DriverService) *DriverController {
	return &DriverController{driverService: driverService}
}

// Createdriver
// @Tags driver
// @Summary Create a new driver
// @Description Create a new driver
// @Accept json
// @Produce json
// @Param driver body model.Driver true "driver object that needs to be added"
// @Success 200
// @Router /drivers [post]
func (c *DriverController) CreateDriver(ctx *gin.Context) {
	log.Info().Msg("Creating Driver")
	createDriver := model.Driver{}
	err := ctx.ShouldBindJSON(&createDriver)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	driver, er := c.driverService.Create(&createDriver)

	if er != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	serverReponse := gin.H{
		"code": 200,
		"status": "Created",
		"data": driver,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverReponse)

}
