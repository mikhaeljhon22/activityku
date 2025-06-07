package model
import (
	"time"
)
type Todolist struct {
  ID uint `gorm:"primaryKey"`
  UserID uint
  UsersInfo UsersInfo `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
  Task string 
  CreatedAt time.Time `gorm:"autoCreateTime"`

}