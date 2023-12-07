package main

import "gorm.io/gorm"

type Metas struct {
	gorm.Model
	Author      string `json:"author" gorm:"column:author"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	KeyWord     string `json:"keyword" gorm:"column:keyword"`
	Url         string `json:"url" gorm:"column:url"`
	Image       string `json:"image" gorm:"column:image"`
	Favicon     string `json:"favicon" gorm:"column:favicon"`
	Owner       uint   `json:"owner" gorm:"column:owner"`
}
