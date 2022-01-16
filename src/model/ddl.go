package model

import (
	"DGoDoBackend/src/db"
	"time"
)

type DDL struct {
	Id          int       `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Content     string    `gorm:"not null" json:"content"`
	Create_time time.Time `gorm:"not null" json:"create_time"`
	Expect_time time.Time `gorm:"not null;default:null" json:"expect_time"`
	Finish_time time.Time `gorm:"default:null" json:"finish_time"`
	Status      int       `gorm:"not null;default:0" json:"status"`
}

func (DDL) TableName() string {
	return "DDL"
}

const (
	DDLActive  = 0
	DDLDone    = 1
	DDLAbandon = 2
)

func GetAllDDLByName(name string) ([]DDL, error) {

	var ans []DDL

	err := db.MySqlDb.
		Model(&DDL{}).
		Where("name = ?", name).
		Find(&ans).
		Error

	return ans, err
}

func GetDDLByID(id int) (DDL, error) {

	var ans DDL

	err := db.MySqlDb.
		Model(&DDL{}).
		First(&ans, id).
		Error
	return ans, err
}

func (ddl DDL) UpdateDDLStatus(status int) (DDL, error) {

	ddl.Status = status
	ddl.Finish_time = time.Now()

	err := db.MySqlDb.
		Save(ddl).
		Error

	return ddl, err
}

func (ddl DDL) CreateDDL() (DDL, error) {

	err := db.MySqlDb.
		Create(&ddl).
		Error

	return ddl, err
}
