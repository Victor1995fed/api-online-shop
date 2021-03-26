package apiserver

import "github.com/gorilla/handlers"

func (s *server) configureRouter() {
	// s.router.Use(s.setRequestID)
	// s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	// userController := *controller.User
	// Product
	s.router.HandleFunc("/product", s.handleProductFind()).Methods("GET")
	s.router.HandleFunc("/product-list", s.handleProductList()).Methods("GET")
	s.router.HandleFunc("/product", s.handleProductCreate()).Methods("POST")
	s.router.HandleFunc("/product", s.handleProductUpdate()).Methods("PUT")
	s.router.HandleFunc("/product", s.handleProductDelete()).Methods("DELETE")
	s.router.HandleFunc("/product/{id:[0-9]+}/add/image", s.handleProductAddImage()).Methods("PUT")
	//s.router.HandleFunc("/product/remove/image", s.handleProductRemoveImage()).Methods("DELETE")

	//Tag
	s.router.HandleFunc("/tag", s.handleTagFind()).Methods("GET")
	s.router.HandleFunc("/tag-list", s.handleTagList()).Methods("GET")
	s.router.HandleFunc("/tag", s.handleTagCreate()).Methods("POST")
	s.router.HandleFunc("/tag", s.handleTagUpdate()).Methods("PUT")
	s.router.HandleFunc("/tag", s.handleTagDelete()).Methods("DELETE")

	//Order
	s.router.HandleFunc("/order", s.handleOrderFind()).Methods("GET")
	s.router.HandleFunc("/order-list", s.handleOrderList()).Methods("GET")
	s.router.HandleFunc("/order", s.handleOrderCreate()).Methods("POST")
	s.router.HandleFunc("/order", s.handleOrderUpdate()).Methods("PUT")
	s.router.HandleFunc("/order", s.handleOrderDelete()).Methods("DELETE")

	// println(controller.User)
	// s.router.HandleFunc("/users", s.listUsers()).Methods("GET")
	s.router.HandleFunc("/users", s.ListUsers()).Methods("GET")
	// s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")

	// private := s.router.PathPrefix("/private").Subrouter()
	// private.Use(s.authenticateUser)
	// private.HandleFunc("/whoami", s.handlerWhoami()).Methods("GET")
}
