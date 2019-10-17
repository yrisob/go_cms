package handlers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yrisob/cms_core/services/dto"

	_ "github.com/go-sql-driver/mysql"
)

// TestFuncName1 is the name of the function that is tested
func TestGetTests(t *testing.T) {

	var udto dto.UserDTO
	udto.Phone = "9216551233"

	test := reflect.ValueOf(udto).Elem()
	typ := test.Type()

	ms := (reflect.New(typ).Elem()).Interface()
	fmt.Println(ms)
	// engine, err := xorm.NewEngine("mysql", "root:gfhjkm1986@/cms_project?parseTime=true")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// assert.Nil(t, err)

	// var dataFromDB = []User{}

	// if err = engine.In("id", 1, 8).Find(&dataFromDB); err != nil {
	// 	fmt.Println(err.Error())
	// }
	// assert.Nil(t, err)
	// fmt.Println("count of records is: ", len(dataFromDB))
	// for _, row := range dataFromDB {
	// 	fmt.Println(row.Email)
	// }

	// var user = &User{ID: 1}
	// if _, err = engine.Get(user); err != nil {
	// 	fmt.Println(err.Error())
	// }

	// assert.Nil(t, err)

	// fmt.Println(fmt.Sprintf("User with id=%v and name=%s", user.ID, user.Name))
}
