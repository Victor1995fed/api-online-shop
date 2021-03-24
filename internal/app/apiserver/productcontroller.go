package apiserver

import (
	"api-online-store/internal/app/filter"
	"api-online-store/internal/app/model"
	"encoding/json"
	"github.com/gorilla/schema"
	"net/http"
)

func (s *server) handleProductCreate() http.HandlerFunc {

	type request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Price       string `json:"price"`
		Tags        []int  `json:"tags"`
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
			Tags:        nil,
		}

		for _, s := range req.Tags {
			t := model.Tag{}
			t.ID = s
			p.Tags = append(p.Tags, t)
		}
		if err := s.store.Product().Create(p); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// p.Sanitaze()
		s.respond(w, r, http.StatusCreated, p)

	}
}

func (s *server) handleProductFind() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &filter.Product{}
		var decoder = schema.NewDecoder()
		err := decoder.Decode(req, r.URL.Query())
		id := req.Id
		p, err := s.store.Product().Find(id)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusFound, p)
	}
}

func (s *server) handleProductList() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &filter.Product{}
		var decoder = schema.NewDecoder()
		err := decoder.Decode(req, r.URL.Query())
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		p, err := s.store.Product().List(req)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusFound, p)

	}
}

func (s *server) handleProductUpdate() http.HandlerFunc {
	type request struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Price       string `json:"price"`
		Tags        []int  `json:"tags"`
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
			ID:          req.ID,
		}

		for _, s := range req.Tags {
			t := model.Tag{}
			t.ID = s
			p.Tags = append(p.Tags, t)
		}

		if err := s.store.Product().Update(p); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, p)

	}
}

func (s *server) handleProductDelete() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		p := &model.Product{
			ID: req.ID,
		}

		if err := s.store.Product().Delete(p); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusAccepted, true)
	}
}
