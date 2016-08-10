package models
import (
	//"github.com/jinzhu/gorm"
)

type Likes struct  {
	//gorm.Model
	ID                 int        `gorm:"AUTO_INCREMENT,primary_key"`
	PostID             int
	UserID             int64
}
