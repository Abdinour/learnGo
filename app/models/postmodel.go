package models

import "time"

//import "github.com/jinzhu/gorm"

type  Post struct  {
	//gorm.Model
	ID             int                `gorm:"AUTO_INCREMENT,primary_key"`
	Title          string              `json:"title"`
	Body           string              `json:"body"`
	UserID          int64
	User              User               `json:"user"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt       *time.Time           `sql:"index"`



}
