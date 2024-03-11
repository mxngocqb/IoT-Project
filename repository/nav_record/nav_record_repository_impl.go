package nav_record

import (
	"github.com/mxngocqb/IoT-Project/model"
	"gorm.io/gorm"
)

type NavRecordRepositoryImpl struct {
	Db *gorm.DB
}

func NewNavRecordRepository(db *gorm.DB) *NavRecordRepositoryImpl {
	db.AutoMigrate(&model.NavRecord{})
	return &NavRecordRepositoryImpl{Db: db}
}