package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:100" json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	// Token    string `json:"token" form:"token"`
}
