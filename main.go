package main

import (
	"github.com/labstack/echo"

	"github.com/yrisob/cms_core/custom_validator"
	"github.com/yrisob/cms_core/handlers"
	"github.com/yrisob/cms_core/models"
	"github.com/yrisob/cms_core/services"
	"github.com/yrisob/cms_core/services/dto"
)

func main() {
	server := handlers.New()
	server.Validator = custom_validator.NewValidator()
	db, err := models.GetDBEngine()
	if err != nil {
		println(err.Error())
		panic(err)
	}

	server.AddCrudRoute("user", services.CreateUserService(db), func(c echo.Context) (interface{}, error) {
		user := &dto.UserDTO{}
		err := c.Bind(user)
		if err != nil {
			return nil, err
		}
		if err = server.Validator.Validate(user); err != nil {
			return nil, err
		}
		return *user, nil
	})
	server.Logger.Fatal(server.Start(":8080"))
}
