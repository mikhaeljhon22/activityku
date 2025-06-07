package service
import (
	"activityku-go/model"
	"activityku-go/repository"
	"activityku-go/util"
	"activityku-go/DTO"
)

type UserService interface{
	RegisterUser(userDTO DTO.UserSignUp) error
    Login(username string, password string) (*model.UsersInfo,error) 
}

type userService struct {
	repo repository.UserRepository
}
func NewUserService(repo repository.UserRepository) UserService{
	return &userService{repo}
}

func (s *userService) RegisterUser(userDTO DTO.UserSignUp) error {
	HashPw:= util.HashPassword(userDTO.Password)
	user := model.UsersInfo{
		Name: userDTO.Name,
		Username: userDTO.Username,
		Email: userDTO.Email,
		Password: HashPw,
	}

	return s.repo.CreateUser(user)
}

func (s *userService) Login(username string, password string) (*model.UsersInfo,error){
	HashPw:= util.HashPassword(password)
	return s.repo.Login(username,HashPw)
}
