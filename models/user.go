package models

import (
	"time"
)

type BasicModel interface{}

type User struct {
	ID          int       `xorm:"pk serial 'id'" json:"id"`
	CreatedDate time.Time `xorm:"created 'createdDate'" json:"createdDate"`
	UpdatedDate time.Time `xorm:"updated 'updatedDate'" json:"updatedDate"`
	Name        string    `xorm:"name" json:"name"`
	Email       string    `xorm:"email" json:"email"`
	Phone       string    `xorm:"phone" json:"phone"`
	Role        int       `xorm:"role" json:"role"`
	Password    string    `xorm:"password" json:"-"`
}
