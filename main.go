package main

import (
	"github.com/tetzng/golang-blog/controller"
	"github.com/tetzng/golang-blog/db"
	"github.com/tetzng/golang-blog/repository"
	"github.com/tetzng/golang-blog/router"
	"github.com/tetzng/golang-blog/usecase"
	"github.com/tetzng/golang-blog/validator"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userValidator := validator.NewUserValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
