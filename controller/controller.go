package controller 
import (
	"activityku-go/service"
	"activityku-go/repository"
	"activityku-go/DTO"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"fmt"

)
var userService service.UserService
func Init(db *gorm.DB){
	repo := repository.NewUserRepository(db)
	userService = service.NewUserService(repo)
}

func RegisterUser(c *gin.Context){
	var user DTO.UserSignUp
	if err:= c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	createUser:= userService.RegisterUser(user)
	if createUser != nil {
		c.JSON(500, gin.H{
			"error": createUser,
		})
		return
	}
		c.JSON(200, gin.H{
			"message": "success to create account",
		})
}

func LoginUser(c *gin.Context){
	var user DTO.UserLoginDTO
	if err:= c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	loginUser,err := userService.Login(user.Username,user.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200,gin.H{
		"message": "success to login",
	})
}