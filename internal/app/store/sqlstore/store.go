package sqlstore

import (
	"database/sql"
	// _ "github.com/jackc/pgx/v4/stdlib" // ...
)

//Store ...
type Store struct {
	db *sql.DB
	// userRepository *UserRepository
}

//New ...
func New(db *sql.DB) Store {
	return Store{
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

// 	s.userRepository = &UserRepository{
// 		store: s,
// 	}
// 	return s.userRepository
// }
