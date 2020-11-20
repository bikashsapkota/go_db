package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Email        	   string `json:"email"`
	FirstName     	   string `json:"first_name"`
	MiddleName   	   string `json:"middle_name"`
	LastName     	   string `json:"last_name"`
	Address      	   string `json:"address"`
	MobileNumber 	   string `json:"mobile_number"`
	EmailNotification  bool   `json:"email_notification"`
	SmsNotification    bool   `json:"sms_notification"`
	PhoneNotification  bool   `json:"phone_notification"`
	PaymentMethod	   int
}

func (u *Users) FullName() string {
	var fullName string
	if u.MiddleName != "" {
		fullName = fmt.Sprintf("%s %s %s", u.FirstName, u.MiddleName, u.LastName)
	} else {
		fullName = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	}
	return fullName
}
