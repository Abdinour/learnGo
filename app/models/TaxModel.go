package models

import (
	//"github.com/jinzhu/gorm"
	"time"
)

type Tax struct  {
	//gorm.Model
	ID                     int        `gorm:"AUTO_INCREMENT,primary_key"`
	CitizenID              int64
	CitizenName            string
	PropertyID             int64
	Money                   string
        Createdat               time.Time
	Updatedat               time.Time
}
