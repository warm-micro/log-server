package db

import "gorm.io/gorm"

// 단일 로그 저장
type ApiLog struct {
	gorm.Model
	Api     string
	Status  string
	Latency string
	Method  string
}

// 서비스 호출 횟수 저장
type ApiCount struct {
	gorm.Model
	Api    string
	Count  uint
	Method string
}
