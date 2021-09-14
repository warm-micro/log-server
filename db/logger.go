package db

import "gorm.io/gorm"

type ApiLog struct {
	gorm.Model
	Api     string
	Status  string
	Latency string
	Method  string
}

type ApiCount struct {
	gorm.Model
	Api    string
	Count  uint
	Method string
}
