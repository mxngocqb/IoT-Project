package vehicle

import "github.com/mxngocqb/IoT-Project/model"

type VehicleService interface {
	Create(category *model.Vehicle) (*model.Vehicle, error)
	ReadAll() ([]model.Vehicle, error)
	ReadByID(categoryID string) (*model.Vehicle, error)
	Update(category *model.Vehicle) (*model.Vehicle, error)
	Delete(categoryID string) error
}
