package entity

import "gorm.io/gorm"

// 消息
type Message struct {
	gorm.Model
	FromId   uint
	TargetId uint
	Type     string
	Media    int
	Content  string
	Pic      string
	Url      string
	Desc     string
	Amount   int
}
