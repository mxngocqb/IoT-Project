package vehicle

import (
	"github.com/go-playground/validator/v10"
	"github.com/mxngocqb/IoT-Project/model"
	"github.com/mxngocqb/IoT-Project/repository/vehicle"
)

type VehicleServiceImpl struct {
	VehicleRepository vehicle.VehicleRepository
	Validate         *validator.Validate
}

func NewVehicleService(VehicleRepository vehicle.VehicleRepository, validate *validator.Validate) *VehicleServiceImpl {
	return &VehicleServiceImpl{VehicleRepository: VehicleRepository, Validate: validate}
}

func (c *VehicleServiceImpl) Create(vehicle *model.Vehicle) (*model.Vehicle, error) {
	err := c.Validate.Struct(vehicle)
	if err != nil {
		return nil, err
	}

	return c.VehicleRepository.Save(vehicle)
}

func (c *VehicleServiceImpl) ReadAll() ([]model.Vehicle, error) {
	return c.VehicleRepository.FindAll()
}

func (c *VehicleServiceImpl) ReadByID(vehicleId string) (*model.Vehicle, error) {
	return c.VehicleRepository.FindByID(vehicleId)
}

func (c *VehicleServiceImpl) Update(vehicle *model.Vehicle) (*model.Vehicle, error) {
	err := c.Validate.Struct(vehicle)
	if err != nil {
		return nil, err
	}
	return c.VehicleRepository.Update(vehicle)
}

func (c *VehicleServiceImpl) Delete(vehicleId string) error {
	return c.VehicleRepository.Delete(vehicleId)
}
