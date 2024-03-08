package driver

import "github.com/mxngocqb/IoT-Project/model"

type DriverService interface {
	Create(category *model.Driver) (*model.Driver, error)
	ReadAll() ([]model.Driver, error)
	ReadByID(categoryID string) (*model.Driver, error)
	Update(category *model.Driver) (*model.Driver, error)
	Delete(categoryID string) error
}
