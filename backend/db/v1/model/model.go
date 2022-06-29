package model

type Users struct {
	Userid string `gorm:"column:userid"`
	Name   string `gorm:"column:name"`
}
