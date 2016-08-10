package models

import (
//"github.com/jinzhu/gorm"
"time"
)

type Property struct  {
	//gorm.Model
	ID                   int        `gorm:"AUTO_INCREMENT,primary_key"`
	ItemNumber           string
	Description         string
	TaxID                 int64
	Createdat             time.Time
	Updatedat              time.Time
}

