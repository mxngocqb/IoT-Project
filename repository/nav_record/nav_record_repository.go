package nav_record

import "github.com/mxngocqb/IoT-Project/model"

type NavRecordRepository interface {
	Save(category *model.NavRecord) (*model.NavRecord, error)
	FindAll() ([]model.NavRecord, error)
	FindByID(categoryID string) (*model.NavRecord, error)
	Update(category *model.NavRecord) (*model.NavRecord, error)
	Delete(categoryID string) error
}
