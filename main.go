package main
import (
	"activityku-go/config"
	"activityku-go/model"
	"activityku-go/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	db,err := config.ConnectTOPostgresql()
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("success connect to postgresql", db)
	}
	db.AutoMigrate(&model.UsersInfo{},&model.Todolist{})

	r := gin.Default()
	
	controller.Init(db)
	/*
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	*/
	api := r.Group("/api/v1")
	api.POST("/signUp", controller.RegisterUser)
	api.POST("/login", controller.LoginUser)
	r.Run(":8080")
}