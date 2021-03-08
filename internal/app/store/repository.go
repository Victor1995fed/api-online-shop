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
	Update(*model.Product) error
	List(map[string]string) ([]model.Product, error)
	Delete(*model.Product) error
}

// TagRepository ...
type TagRepository interface {
	Create(*model.Tag) error
	Find(int) (*model.Tag, error)
	Update(*model.Tag) error
	List(map[string]string) ([]model.Tag, error)
	Delete(*model.Tag) error
}

// OrderRepository ...
type OrderRepository interface {
	Create(*model.Order) error
	Find(int) (*model.Order, error)
	Update(*model.Order) error
	List(map[string]string) ([]model.Order, error)
	Delete(*model.Order) error
}

