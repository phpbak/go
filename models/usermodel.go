package models

import (
	"github.com/astaxie/beego"
)

type User struct {
	Id         int
	LoginName  string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (a *User) TableName() string {
	return TableName("uc_admin")
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}