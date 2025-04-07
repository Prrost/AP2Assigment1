package Storage

import "inventory-service/domain"

type Storage interface {
	CreateObject(object domain.Object) (domain.Object, error)
	UpdateObjectByID(id int, object domain.Object) (domain.Object, error)
	GetObjectByID(id int) (domain.Object, error)
	DeleteObjectByID(id int) (domain.Object, error)
	IsProductExists(name string) (bool, error)
	ListProducts(name string, limit int, offset int) ([]domain.Object, error)
}
