package handlers

import (
	"net/http"

	"github.com/yrisob/cms_core/services/dto"

	"github.com/labstack/echo"
)

//GetUsers  method for testing routing
// func GetUsers(c echo.Context) (err error) {
// 	db, err := models.GetDBEngine()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
// 	}
// 	if users, err := services.CreateUserService(db).FindAll(); err == nil {
// 		return c.JSON(http.StatusOK, users)
// 	}
// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
// }

//GetUser  method for checking work with params
// func GetUser(c echo.Context) (err error) {
// 	var id int
// 	if id, err = strconv.Atoi(c.Param("id")); err == nil {
// 		if user, err := models.GetUserById(id); err == nil {
// 			return c.JSON(http.StatusOK, user)
// 		}
// 	}
// 	return c.JSON(http.StatusNotFound, map[string]interface{}{"message": err.Error()})
// }

//PostTest insert test data into
func PostTest(c echo.Context) (err error) {
	user := new(dto.UserDTO)
	if err = c.Bind(user); err != nil {
		return
	}
	if err = c.Validate(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}
