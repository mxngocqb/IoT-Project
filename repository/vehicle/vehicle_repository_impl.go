package vehicle

import (
	"github.com/mxngocqb/IoT-Project/model"
	"gorm.io/gorm"
)

type VehicleRepositoryImpl struct {
	Db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) *VehicleRepositoryImpl {
	db.AutoMigrate(&model.Vehicle{})
	return &VehicleRepositoryImpl{Db: db}
}

func (c *VehicleRepositoryImpl) Save(vehicle *model.Vehicle) (*model.Vehicle, error) {
	err := c.Db.Create(vehicle).Error
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}

func (c *VehicleRepositoryImpl) FindAll() ([]model.Vehicle, error) {
	var categories []model.Vehicle
	err := c.Db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *VehicleRepositoryImpl) FindByID(categoryID string) (*model.Vehicle, error) {
	var category model.Vehicle
	err := c.Db.Where("vehicle_id = ?", categoryID).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *VehicleRepositoryImpl) Update(category *model.Vehicle) (*model.Vehicle, error) {
	err := c.Db.Save(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *VehicleRepositoryImpl) Delete(categoryID string) error {
	var category model.Vehicle
	err := c.Db.Where("vehicle_id = ?", categoryID).Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}