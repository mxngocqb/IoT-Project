package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mxngocqb/IoT-Project/model"
	"github.com/mxngocqb/IoT-Project/service/vehicle"
	"github.com/rs/zerolog/log"
)

type VehicleController struct {
	vehicleService vehicle.VehicleService
}

func NewVehicleController(vehicleService vehicle.VehicleService) *VehicleController {
	return &VehicleController{vehicleService: vehicleService}
}

type VehicleServerResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// CreateVehicle handles the creation of a new vehicle
// @Summary Create a new vehicle
// @Description Create a new vehicle with the provided information
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param vehicle body model.Vehicle true "Vehicle information"
// @Success 200 {object} VehicleServerResponse "Successful response"
// @Failure 400 {object} VehicleServerResponse "Bad request"
// @Failure 500 {object} VehicleServerResponse "Internal server error"
// @Router /vehicles [post]
func (c *VehicleController) CreateVehicle(ctx *gin.Context) {
	log.Info().Msg("Creating Vehicle")
	createVehicle := model.Vehicle{}
	err := ctx.ShouldBindJSON(&createVehicle)

	if err != nil {
		ctx.JSON(400, VehicleServerResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}
	vehicle, er := c.vehicleService.Create(&createVehicle)

	if er != nil {
		ctx.JSON(500, VehicleServerResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
		return
	}

	serverReponse := VehicleServerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   vehicle,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverReponse)
}

// ReadAllVehicle handles the retrieval of all vehicles
// @Summary Retrieve all vehicles
// @Description Retrieve all vehicles stored in the database
// @Tags Vehicle
// @Produce json
// @Success 200 {object} VehicleServerResponse "Successful response"
// @Failure 500 {object} VehicleServerResponse "Internal server error"
// @Router /vehicles [get]
func (c *VehicleController) ReadAllVehicle(ctx *gin.Context) {
	log.Info().Msg("Reading All Vehicle")
	vehicles, err := c.vehicleService.ReadAll()

	if err != nil {
		ctx.JSON(500, VehicleServerResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	serverResponse := VehicleServerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   vehicles,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}

// ReadVehicleByID handles the retrieval of a vehicle by its ID
// @Summary Retrieve a vehicle by its ID
// @Description Retrieve a vehicle from the database by its unique identifier
// @Tags Vehicle
// @Produce json
// @Param vehicleID path string true "Vehicle ID"
// @Success 200 {object} VehicleServerResponse "Successful response"
// @Failure 500 {object} VehicleServerResponse "Internal server error"
// @Router /vehicles/{vehicleID} [get]
func (c *VehicleController) ReadVehicleByID(ctx *gin.Context) {
	log.Info().Msg("Reading Category By ID")
	vehicleID := ctx.Param("vehicleID")
	vehicle, err := c.vehicleService.ReadByID(vehicleID)
	if err != nil {
		ctx.JSON(500, VehicleServerResponse{Code: http.StatusInternalServerError, Status: "Internal Server Error", Data: err.Error()})
		return
	}
	serverResponse := VehicleServerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   vehicle,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}

// UpdateVehicle handles the updating of a vehicle
// @Summary Update a vehicle
// @Description Update an existing vehicle with the provided information
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param vehicle body model.Vehicle true "Vehicle information"
// @Success 200 {object} VehicleServerResponse "Successful response"
// @Failure 400 {object} VehicleServerResponse "Bad request"
// @Failure 500 {object} VehicleServerResponse "Internal server error"
// @Router /vehicles [put]
func (c *VehicleController) UpdateVehicle(ctx *gin.Context) {
	log.Info().Msg("Updating Vehicle")
	updateVehicle := model.Vehicle{}
	err := ctx.ShouldBindJSON(&updateVehicle)
	if err != nil {
		ctx.JSON(400, VehicleServerResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}
	vehicle, er := c.vehicleService.Update(&updateVehicle)
	if er != nil {
		ctx.JSON(500, VehicleServerResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Erro",
			Data:   err.Error(),
		})
		return
	}
	serverResponse := VehicleServerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   vehicle,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}

// DeleteVehicle handles the deletion of a vehicle
// @Summary Delete a vehicle
// @Description Delete a vehicle from the database by its ID
// @Tags Vehicle
// @Param vehicleID path string true "Vehicle ID"
// @Produce json
// @Success 200 {object} VehicleServerResponse "Successful response"
// @Failure 500 {object} VehicleServerResponse "Internal server error"
// @Router /vehicles/{vehicleID} [delete]
func (c *VehicleController) DeleteVehicle(ctx *gin.Context) {
	log.Info().Msg("Deleting Vehicle")
	VehicleID := ctx.Param("vehicleID")
	err := c.vehicleService.Delete(VehicleID)
	if err != nil {
		ctx.JSON(500, VehicleServerResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
		return
	}
	serverResponse := VehicleServerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   VehicleID,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}
