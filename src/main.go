package main

import (
	"controller"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"interactor"
	"model"
	"repository"
	"request"
)

func initDb() *gorm.DB {
	db, err := gorm.Open("mysql", "master:gagagigu123@tcp(localhost:3306)/clean_architecture_practice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Task{}, &model.User{})
	return db
}

func initRepository(db *gorm.DB) repository.UserRepository {
	return repository.DbUserRepository{Db: db}
}

func Init(db *gorm.DB) controller.UserControllerInterface {
	userRepository := initRepository(db)
	interactor := interactor.UserSignUpInteractor{UserRepository: userRepository}
	return controller.UserController{UserSignUpCase: interactor}
}

func main() {
	db := initDb()
	defer db.Close()
	controller := Init(db)
	request := request.UserCreateRequest{Name: "kawaguchi", Email: "kawaguchi.yasutaka@gmail.com"}
	response, _ := controller.SignUp(request)
	fmt.Println(response)
	fmt.Println(repository.UsersStore)
}
