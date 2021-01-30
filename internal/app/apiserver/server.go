package apiserver

import (
	"api-online-store/internal/app/store"
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// "context"
// "encoding/json"
// "errors"
// "net/http"
// "test-golang-api/internal/app/models"

// "time"

// "github.com/google/uuid"
// "github.com/gorilla/handlers"
// "github.com/gorilla/mux"
// "github.com/gorilla/sessions"
// "github.com/sirupsen/logrus"

// const (
// 	sessionName        = "testapi"
// 	ctxKeyUser  ctxKey = iota
// 	ctxKeyRequestID
// )

// var (
// 	errIncorrectEmailOrPassword = errors.New("imcorrect email or password")
// 	errNotAuthanticated         = errors.New("not authanticated")
// )

// type ctxKey int8
type server struct {
	router *mux.Router
	// logger       *logrus.Logger
	store store.Store
	// controllers
	// sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		// logger:       logrus.New(),
		store: store,
		// sessionStore: sessionStore,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	// s.router.Use(s.setRequestID)
	// s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	// userController := *controller.User
	s.router.HandleFunc("/product", s.handleProductCreate()).Methods("POST")
	// println(controller.User)
	// s.router.HandleFunc("/users", s.listUsers()).Methods("GET")
	s.router.HandleFunc("/users", s.ListUsers()).Methods("GET")
	// s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")

	// private := s.router.PathPrefix("/private").Subrouter()
	// private.Use(s.authenticateUser)
	// private.HandleFunc("/whoami", s.handlerWhoami()).Methods("GET")
}

// func (s *server) listUsers() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// req := &request{}
// 		// if err := json.NewDecoder(r.Body).Decode(req); err != nil {
// 		// 	s.error(w, r, http.StatusBadRequest, err)
// 		// 	return
// 		// }
// 		// u := &models.User{
// 		// 	Email:    req.Email,
// 		// 	Password: req.Password,
// 		// }

// 		// if err := s.store.User().Create(u); err != nil {
// 		// 	s.error(w, r, http.StatusUnprocessableEntity, err)
// 		// 	return
// 		// }
// 		// u.Sanitaze()
// 		fmt.Fprint(w, "Users!\n")
// 		// s.respond(w, r, http.StatusCreated, nil)

// 	}
// }

// func (s *server) setRequestID(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		id := uuid.New().String()
// 		w.Header().Set("X-Request-ID", id)
// 		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))

// 	})
// }

// func (s *server) logRequest(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		logger := s.logger.WithFields(logrus.Fields{
// 			"remote_addr": r.RemoteAddr,
// 			"request_id":  r.Context().Value(ctxKeyRequestID),
// 		})

// 		logger.Infof("started %s %s", r.Method, r.RequestURI)
// 		start := time.Now()
// 		rw := &responseWriter{w, http.StatusOK}
// 		next.ServeHTTP(rw, r)
// 		logger.Info(
// 			"completed with %d %s in %v",
// 			rw.code,
// 			http.StatusText(rw.code),
// 			time.Now().Sub(start),
// 		)
// 	})
// }

// func (s *server) authenticateUser(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		session, err := s.sessionStore.Get(r, sessionName)
// 		if err != nil {
// 			s.error(w, r, http.StatusInternalServerError, err)
// 			return
// 		}
// 		id, ok := session.Values["user_id"]
// 		if !ok {
// 			s.error(w, r, http.StatusUnauthorized, errNotAuthanticated)
// 			return
// 		}

// 		u, err := s.store.User().Find(id.(int))
// 		if err != nil {
// 			s.error(w, r, http.StatusUnauthorized, errNotAuthanticated)
// 			return
// 		}

// 		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))

// 	})
// }

// func (s *server) handlerWhoami() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*models.User))
// 	}
// }

// func (s *server) handleUsersCreate() http.HandlerFunc {

// 	type request struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		req := &request{}
// 		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
// 			s.error(w, r, http.StatusBadRequest, err)
// 			return
// 		}
// 		u := &models.User{
// 			Email:    req.Email,
// 			Password: req.Password,
// 		}

// 		if err := s.store.User().Create(u); err != nil {
// 			s.error(w, r, http.StatusUnprocessableEntity, err)
// 			return
// 		}
// 		u.Sanitaze()
// 		s.respond(w, r, http.StatusCreated, u)

// 	}
// }

// func (s *server) handleSessionsCreate() http.HandlerFunc {

// 	type request struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		req := &request{}
// 		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
// 			s.error(w, r, http.StatusBadRequest, err)
// 			return
// 		}
// 		u, err := s.store.User().FindByEmail(req.Email)
// 		if err != nil || !u.ComparePassword(req.Password) {
// 			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
// 			return
// 		}
// 		session, err := s.sessionStore.Get(r, sessionName)
// 		if err != nil {
// 			s.error(w, r, http.StatusInternalServerError, err)
// 			return
// 		}
// 		session.Values["user_id"] = u.ID
// 		s.sessionStore.Save(r, w, session)
// 		if err != nil {
// 			s.error(w, r, http.StatusInternalServerError, err)
// 			return
// 		}
// 		s.respond(w, r, http.StatusOK, nil)
// 	}
// }

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
