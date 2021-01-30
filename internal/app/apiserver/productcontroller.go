package apiserver

import (
	"api-online-store/internal/app/model"
	"encoding/json"
	"net/http"
)

// func (s *server) handleProductCreate() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Users!\n")
// 	}
// }

func (s *server) handleProductCreate() http.HandlerFunc {

	type request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Price       string `json:"price"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		p := &model.Product{
			Title:       req.Title,
			Description: req.Description,
			Price:       req.Price,
		}

		// if err := s.store.User().Create(u); err != nil {
		// 	s.error(w, r, http.StatusUnprocessableEntity, err)
		// 	return
		// }
		// p.Sanitaze()
		s.respond(w, r, http.StatusCreated, p)

	}
}
