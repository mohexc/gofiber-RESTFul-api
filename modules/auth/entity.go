package auth

import "restful-api/modules/base"

type user struct {
	base.Model
	email    string `gorm:"email"`
	password string `gorm:"password"`
}
