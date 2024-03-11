package model

type Vehicle struct {
	VehicleId    uint   `json:"vehicle_id" gorm:"column:vehicle_id;primary_key" validate:"required"`
	VehiclePlate string `json:"vehicle_plate" gorm:"column:vehicle_plate;not null;type:varchar(15);default:null" validate:"required"`
	VehicleName  string `json:"vehicle_name" gorm:"column:vehicle_name"`
}
