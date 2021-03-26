package apiserver

import "net/http"

func (s *server) handleImageUpload() http.HandlerFunc {

	//Just upload images
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

func (s *server) handleFileDownload() http.HandlerFunc {

	//return image
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
