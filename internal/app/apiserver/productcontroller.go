package apiserver

import (
	"api-online-store/internal/app/model"
	"encoding/json"
	"net/http"
)

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

		if err := s.store.Product().Create(p); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// p.Sanitaze()
		s.respond(w, r, http.StatusCreated, p)

	}
}

func (s *server) handleProductFind() http.HandlerFunc {
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
		p, err := s.store.Product().Find(id)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// p.Sanitaze()
		s.respond(w, r, http.StatusFound, p)
		// fmt.Fprint(w, "Users!\n")
	}
}

func (s *server) handleProductList() http.HandlerFunc {
	type request struct {
		Count string `json:"count"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		// id, _ := strconv.Atoi(req.ID)
		var m map[string]string
		m["count"] = req.Count
		if _, err := s.store.Product().List(m); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// p.Sanitaze()
		s.respond(w, r, http.StatusCreated, m)

	}
}

func (s *server) handleProductUpdate() http.HandlerFunc {
	type request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Price       string `json:"price"`
		ImgURL      string `json:"img_url"`
		ID          int    `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		// id, _ := strconv.Atoi(req.ID)
		p := &model.Product{
			Title:       req.Title,
			Description: req.Description,
			Price:       req.Price,
			ImageURL:    req.ImgURL,
			ID:          req.ID,
		}

		if err := s.store.Product().Update(p); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// p.Sanitaze()
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
		// id, _ := strconv.Atoi(req.ID)
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
