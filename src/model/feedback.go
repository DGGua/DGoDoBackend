package model

import (
	"DGoDoBackend/src/db"
	"time"
)

type Feedback struct {
	Id          int       `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Email       string    `gorm:"not null" json:"email"`
	Content     string    `gorm:"not null" json:"content"`
	Create_time time.Time `gorm:"not null" json:"create_time"`
	Status      int       `gorm:"not null;default:0" json:"status"`
}

func (Feedback) TableName() string {
	return "Feedback"
}

const (
	FeedbackActive = 0
	FeedbackSolved = 1
	FeedbackClosed = 2
)

func (fb Feedback) CreateFeedback() (Feedback, error) {

	err := db.MySqlDb.
		Create(&fb).
		Error

	return fb, err
}

func (fb Feedback) UpdateFeedbackStatus(status int) (Feedback, error) {

	fb.Status = status

	err := db.MySqlDb.
		Save(fb).
		Error

	return fb, err
}
