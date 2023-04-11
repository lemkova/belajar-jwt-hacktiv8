package models

import (
	"jwt-h8/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~You must enter your full name"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~You must enter your email address,email~You must enter a valid email address"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~You must enter your password,minstringlength(6)~Your password must be at least 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)

	if err != nil {
		return err
	}

	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return
}
