package services

import (
	"github.com/go-xorm/xorm"
	"github.com/yrisob/cms_core/models"
	"github.com/yrisob/cms_core/services/dto"
)

//UserService  operates with user data model
type UserService struct {
	db *xorm.Engine
}

//CreateUserService method for initialyse user service
func CreateUserService(db *xorm.Engine) *UserService {
	return &UserService{db: db}
}

//FindAll method to find all records in user table
func (us *UserService) FindAll() (*[]interface{}, error) {
	users := []models.User{}
	err := us.db.Find(&users)
	if err != nil {
		return nil, err
	}
	b := make([]interface{}, len(users))
	for i := range users {
		b[i] = users[i]
	}
	return &b, err

}

//GetByID get user record by it id
func (us *UserService) GetByID(id int) (*interface{}, error) {
	user := &models.User{ID: id}
	_, err := us.db.Get(user)
	if err != nil {
		return nil, err
	}
	var result interface{}
	result = *user
	return &result, nil
}

//Insert user to database
func (us *UserService) Insert(data interface{}) (*interface{}, error) {
	userDTO := data.(dto.UserDTO)
	user := &models.User{}
	user.Name = userDTO.Name
	user.Password = userDTO.Password
	user.Role = userDTO.Role
	user.Email = userDTO.Email
	user.Phone = userDTO.Phone
	_, err := us.db.Insert(user)
	if err != nil {
		return nil, err
	}

	var result interface{}
	result = *user
	return &result, nil
}

//Update data by user
func (us *UserService) Update(id int, data interface{}) (*interface{}, error) {
	userDTO := data.(dto.UserDTO)
	user := &models.User{}
	_, err := us.db.ID(id).Get(user)
	if err != nil {
		return nil, err
	}
	user.Name = userDTO.Name
	user.Password = userDTO.Password
	user.Role = userDTO.Role
	user.Email = userDTO.Email
	user.Phone = userDTO.Phone

	_, err = us.db.Id(id).Update(user)
	if err != nil {
		return nil, err
	}
	var result interface{}
	result = *user
	return &result, nil
}

//Delete user data by id
func (us *UserService) Delete(id int) (err error) {
	user := &models.User{}
	_, err = us.db.ID(id).Delete(user)
	if err != nil {
		return err
	}
	return nil
}
