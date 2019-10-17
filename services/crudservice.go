package services

//CrudService parent interface for all crud services
type CrudService interface {
	FindAll() (*[]interface{}, error)
	GetByID(id int) (*interface{}, error)
	Insert(data interface{}) (*interface{}, error)
	Update(id int, data interface{}) (*interface{}, error)
	Delete(id int) error
}
