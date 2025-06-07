package repository
import (
	"activityku-go/model"
	"gorm.io/gorm"
)

type UserRepository interface{
	CreateUser(user model.UsersInfo) error
	Login(username string, password string) (*model.UsersInfo,error) 
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user model.UsersInfo) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) Login(username string, password string) (*model.UsersInfo,error)   {
	var user model.UsersInfo
	finding:= r.db.First(&user, "username = ? AND password = ?", username, password).Error
	return &user, finding
}

