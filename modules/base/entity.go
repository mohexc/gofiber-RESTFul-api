package base

import "time"

type Model struct {
	ID       uint      `gorm:"primary_key"`
	CreateAt time.Time `gorm:"created_at"`
	UpdateAt time.Time `gorm:"updated_at"`
	DeleteAt time.Time `gorm:"deleted_at"`
}
