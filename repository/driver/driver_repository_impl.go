package driver

import (
	"github.com/mxngocqb/IoT-Project/model"
	"gorm.io/gorm"
)

type DriverRepositoryImpl struct {
	Db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) *DriverRepositoryImpl {
	db.AutoMigrate(&model.Driver{})
	return &DriverRepositoryImpl{Db: db}
}

func (c *DriverRepositoryImpl) Save(customer *model.Driver) (*model.Driver, error) {
	err := c.Db.Create(customer).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}
