package models

type User struct {
       ID                       int         `gorm:"AUTO_INCREMENT,primary_key"`
       Name                      string          `json:"name"`
       Email                      string         `json:"email"`
       Password                    string        `json:"password"`

}
