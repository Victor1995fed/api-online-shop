package apiserver

import (
	// "database/sql"
	"api-online-store/internal/app/store/sqlstore"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib" // ...
	// _ "github.com/lib/pq"
	// "github.com/gorilla/sessions"
)

// Start ...
func Start(config *Config) error {
	// connectionStr := configToString(config.Database)
	// fmt.Println(connectionStr)
	db, err := newDB(config.Database)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)

	// sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, nil)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello World!")
	// })
	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseConf database) (*sql.DB, error) {
	db, err := sql.Open("pgx", configToString(databaseConf))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func configToString(databaseConf database) string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		databaseConf.Server, databaseConf.Database, databaseConf.User, databaseConf.Password)
}
