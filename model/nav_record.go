package model

import "time"

type NavRecord struct {
	NavRecordId  uint      `json:"nav_record_id" gorm:"column:nav_record_id;primary_key" validate:"required"`
	DriverId     *uint     `json:"driver_id" gorm:"column:driver_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VehicleId    *uint     `json:"vehicle_id" gorm:"column:vehicle_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	Time         time.Time `json:"time" gorm:"column:time"`
	Latitude     string    `json:"latitude" gorm:"column:latitude"`
	Longitude    string    `json:"longitude" gorm:"column:longitude"`
	LocationInfo float64   `json:"location_info" gorm:"column:location_info"`
	SatelliteNum int       `json:"satellite_num" gorm:"column:satellite_num"`
	Speed        float64   `json:"speed" gorm:"column:speed"`
	Distance     float64   `json:"distance" gorm:"column:distance"`

	Vehicle Vehicle `gorm:"foreignKey:VehicleId"`
	Driver  Driver  `gorm:"foreignKey:DriverId"`
}
