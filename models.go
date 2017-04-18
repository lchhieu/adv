package main

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

func (u *User) TableName() string {
	return "users"
}

type Profile struct {
	Id      int
	Age     int16
	Address string
	User    *User `orm:"reverse(one)"` // Reverse relationship (optional)
}

func (p *Profile) TableName() string {
	return "profiles"
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User), new(Profile))
}
