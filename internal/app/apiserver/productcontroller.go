package apiserver

import (
	"api-online-store/internal/app/filter"
	"api-online-store/internal/app/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
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

func (s *server) handleProductAddImage() http.HandlerFunc  {

	//TODO:: Перенести сохранение файлов в filerepository
	//TODO:: Вынести указание на папку для файлов в конфиги
	return func(w http.ResponseWriter, r *http.Request) {
		//file, _, err := r.FormFile("image")
		// 32 MB is the default used by FormFile()
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		files := r.MultipartForm.File["images"]
		if files == nil {
			s.respond(w, r, http.StatusBadRequest, "Field for files must call is 'images'")
			return
		}
		fmt.Println(files)
		for _, fileHeader := range files {

			//if fileHeader.Size > MAX_UPLOAD_SIZE {
			//	http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			//	return
			//}

			// Open the file
			file, err := fileHeader.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			defer file.Close()

			buff := make([]byte, 512)
			_, err = file.Read(buff)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			filetype := http.DetectContentType(buff)
			if filetype != "image/jpeg" && filetype != "image/png" {
				http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
				return
			}

			_, err = file.Seek(0, io.SeekStart)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = os.MkdirAll("./uploads", os.ModePerm)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			defer f.Close()

			_, err = io.Copy(f, file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}


		s.respond(w, r, http.StatusAccepted, true)
	}

}

//func (s *server) handleProductRemoveImage() http.HandlerFunc {
//	// unlink images by id
//	//...
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		s.respond(w, r, http.StatusAccepted, true)
//	}
//}