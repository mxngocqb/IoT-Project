package driver

import (
	"github.com/go-playground/validator/v10"
	"github.com/mxngocqb/IoT-Project/model"
	"github.com/mxngocqb/IoT-Project/repository/driver"
)

type DriverServiceImpl struct {
	DriverRepository driver.DriverRepository
	Validate         *validator.Validate
}

func NewDriverService(driverRepository driver.DriverRepository, validate *validator.Validate) *DriverServiceImpl {
	return &DriverServiceImpl{DriverRepository: driverRepository, Validate: validate}
}

func (c *DriverServiceImpl) Create(driver *model.Driver) (*model.Driver, error) {
	err := c.Validate.Struct(driver)
	if err != nil {
		return nil, err
	}
	return c.DriverRepository.Save(driver)
}

func (c *DriverServiceImpl) ReadAll() ([]model.Driver, error) {
	return c.DriverRepository.FindAll()
}

func (c *DriverServiceImpl) ReadByID(driverID string) (*model.Driver, error) {
	return c.DriverRepository.FindByID(driverID)
}

func (c *DriverServiceImpl) Update(driver *model.Driver) (*model.Driver, error) {
	err := c.Validate.Struct(driver)
	if err != nil {
		return nil, err
	}
	return c.DriverRepository.Update(driver)
}

func (c *DriverServiceImpl) Delete(driverID string) error {
	return c.DriverRepository.Delete(driverID)
}
