package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/yrisob/cms_core/services"
)

type BindingFunc func(c echo.Context) (interface{}, error)

type CrudServer struct {
	echo.Echo
}

func New() *CrudServer {
	crud := &CrudServer{}
	crud.Echo = *echo.New()
	return crud
}

func (cc *CrudServer) AddCrudRoute(prefix string, crudservice services.CrudService, bindingFunc BindingFunc) {
	cc.GET(prefix, func(c echo.Context) error {
		models, err := (crudservice).FindAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, models)

	})
	cc.GET(fmt.Sprintf("%s/:id", prefix), func(c echo.Context) error {
		var id int
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		}
		model, err := crudservice.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, model)
	})

	cc.POST(prefix, func(c echo.Context) error {
		userDTO, err := bindingFunc(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		}
		model, err := crudservice.Insert(userDTO)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, model)

	})

	cc.PUT(fmt.Sprintf("%s/:id", prefix), func(c echo.Context) error {
		var id int
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		}
		userDTO, err := bindingFunc(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		}
		model, err := crudservice.Update(id, userDTO)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, model)

	})

	cc.DELETE(fmt.Sprintf("%s/:id", prefix), func(c echo.Context) error {
		var id int
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		}
		if err = crudservice.Delete(id); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"id": id})
	})
}
