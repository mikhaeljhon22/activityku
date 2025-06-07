package model
import (
	"time"
)
type UsersInfo struct {
	ID uint `gorm:"primaryKey"`
	Username string 
	Name string
	Email string 
	Password string 
	CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
} 