package models

import (
	//"github.com/jinzhu/gorm"
	"time"
)

type Citizen struct  {
	//gorm.Model
	ID                 int        `gorm:"AUTO_INCREMENT,primary_key"`
	FName              string
	LName              string
	UserID             int64
        Createdat          time.Time
	Updatedat        time.Time
}
