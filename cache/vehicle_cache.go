package cache

import (
	"github.com/mxngocqb/IoT-Project/model"
)

type VehicleCache interface {
	Set(key string, value *model.Vehicle)
	Get(key string) *model.Vehicle
	Delete(key string) error
	GetMultiRequest(key string) []model.Vehicle
	SetMultiRequest(key string, value []model.Vehicle) 
}
