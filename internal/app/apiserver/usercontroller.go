package apiserver

import (
	"fmt"
	"net/http"
)

func (s *server) ListUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// req := &request{}
		// if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		// 	s.error(w, r, http.StatusBadRequest, err)
		// 	return
		// }
		// u := &models.User{
		// 	Email:    req.Email,
		// 	Password: req.Password,
		// }

		// if err := s.store.User().Create(u); err != nil {
		// 	s.error(w, r, http.StatusUnprocessableEntity, err)
		// 	return
		// }
		// u.Sanitaze()
		fmt.Fprint(w, "Users!\n")
		s.respond(w, r, http.StatusCreated, nil)

	}
}
