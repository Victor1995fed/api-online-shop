package sqlstore

import (
	"api-online-store/internal/app/store"
	"database/sql"
	// _ "github.com/jackc/pgx/v4/stdlib" // ...
)

//Store ...
type Store struct {
	db *sql.DB
	// userRepository *UserRepository
	productRepository *ProductRepository
	tagRepository     *TagRepository
	orderRepository     *OrderRepository
}

//New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}

	// fmt.Println(ss.db)
	// return &Store{
	// 	db: db,
	// }
}

//GetDb ...
// func (s *Store) GetDb() string {

// 	return "sdf"
// }

//User ...
func (s *Store) User() string {

	return "test"
}

//Product ...
func (s *Store) Product() store.ProductRepository {

	if s.productRepository != nil {
		return s.productRepository
	}

	s.productRepository = &ProductRepository{
		store: s,
	}
	return s.productRepository
}

//Tag ...
func (s *Store) Tag() store.TagRepository {

	if s.tagRepository != nil {
		return s.tagRepository
	}

	s.tagRepository = &TagRepository{
		store: s,
	}
	return s.tagRepository
}

//Order ...
func (s *Store) Order() store.OrderRepository {

	if s.orderRepository != nil {
		return s.orderRepository
	}

	s.orderRepository = &OrderRepository{
		store: s,
	}
	return s.orderRepository
}

// 	s.userRepository = &UserRepository{
// 		store: s,
// 	}
// 	return s.userRepository
// }
