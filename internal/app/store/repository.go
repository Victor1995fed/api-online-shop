package store

import "api-online-store/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

// ProductRepository ...
type ProductRepository interface {
	Create(*model.Product) error
	Find(int) (*model.Product, error)
	Update(int) (*model.Product, error)
	Delete(int) (bool, error)
}
