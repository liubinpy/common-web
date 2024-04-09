package system

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey"json:"id"`
	Name      string         `json:"name"`
	Account   string         `gorm:"unique"json:"account"`
	Password  string         `json:"password"`
	Email     *string        `json:"email"`
	Age       uint8          `json:"age"`
	Birthday  time.Time      `json:"birthday"`
	CreatedAt time.Time      `gorm:"autoUpdateTime"json:"createAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "sys_user"
}
