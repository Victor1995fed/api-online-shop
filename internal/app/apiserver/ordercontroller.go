package apiserver

import (
	"api-online-store/internal/app/model"
	"encoding/json"
	"net/http"
)

func (s *server) handleOrderCreate() http.HandlerFunc {

	type request struct {
		Price       float64 `json:"address"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		//req := &request{}
		//if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		//	s.error(w, r, http.StatusBadRequest, err)
		//	return
		//}
		//p := &model.Order{
		//	Price:       req.Price,
		//}
		//
		//if err := s.store.Order().Create(p); err != nil {
		//	s.error(w, r, http.StatusUnprocessableEntity, err)
		//	return
		//}
		//// p.Sanitaze()
		//s.respond(w, r, http.StatusCreated, p)

	}
}

func (s *server) handleOrderFind() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		id := req.ID
		p, err := s.store.Order().Find(id)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// p.Sanitaze()
		s.respond(w, r, http.StatusFound, p)
		// fmt.Fprint(w, "Users!\n")
	}
}

func (s *server) handleOrderList() http.HandlerFunc {
	type request struct {
		Count string `json:"count"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		m := make(map[string]string)
		m["count"] = req.Count
		p, err := s.store.Order().List(m)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, p)

	}
}

func (s *server) handleOrderUpdate() http.HandlerFunc {
	type request struct {
		Description string `json:"description"`
		Price       float64 `json:"price"`
		ImgURL      string `json:"img_url"`
		ID          int    `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		//req := &request{}
		//if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		//	s.error(w, r, http.StatusBadRequest, err)
		//	return
		//}
		//p := &model.Order{
		//	Description: req.Description,
		//	Price:       req.Price,
		//	ImageURL:    req.ImgURL,
		//	ID:          req.ID,
		//}
		//
		//if err := s.store.Order().Update(p); err != nil {
		//	s.error(w, r, http.StatusUnprocessableEntity, err)
		//	return
		//}
		//s.respond(w, r, http.StatusCreated, p)

	}
}

func (s *server) handleOrderDelete() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		p := &model.Order{
			ID: req.ID,
		}

		if err := s.store.Order().Delete(p); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusAccepted, true)
	}
}
