package models

import (
	//"github.com/jinzhu/gorm"
	"time"
)

type Region struct  {
	//gorm.Model
	ID                 int          `gorm:"AUTO_INCREMENT,primary_key"`
	LocationName         string
	CitizenID            int64
	Longtitude            string
	Latitude              string
        Createdat          time.Time
	Updatedat        time.Time
}
