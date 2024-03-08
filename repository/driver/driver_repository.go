package driver

import "github.com/mxngocqb/IoT-Project/model"

type DriverRepository interface {
	Save(category *model.Driver) (*model.Driver, error)
	FindAll() ([]model.Driver, error)
	FindByID(categoryID string) (*model.Driver, error)
	Update(category *model.Driver) (*model.Driver, error)
	Delete(categoryID string) error
}
