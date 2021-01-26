package apiserver

import (
	"fmt"
	"net/http"
)

// Start ...
func Start(config *Config) error {
	// db, err := newDB(config.DatabaseURL)
	// if err != nil {
	// 	return err
	// }

	// defer db.Close()
	// store := sqlstore.New(db)
	// sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	// srv := newServer(store, sessionStore)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	return http.ListenAndServe(config.BindAddr, nil)
}

// func newDB(databaseURL string) (*sql.DB, error) {
// 	db, err := sql.Open("pgx", databaseURL)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
