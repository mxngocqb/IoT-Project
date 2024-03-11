package vehicle

import "github.com/mxngocqb/IoT-Project/model"

type VehicleRepository interface {
	Save(category *model.Vehicle) (*model.Vehicle, error)
	FindAll() ([]model.Vehicle, error)
	FindByID(categoryID string) (*model.Vehicle, error)
	Update(category *model.Vehicle) (*model.Vehicle, error)
	Delete(categoryID string) error
}
