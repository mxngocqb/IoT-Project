package cache

import (
	"github.com/mxngocqb/IoT-Project/model"
)

type DriverCache interface {
	Set(key string, value *model.Driver)
	Get(key string) *model.Driver
	Delete(key string) error
}
