package main

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name      string `gorm:"primaryKey"`
	Accountid string
	Id        string
	Puuid     string
}
