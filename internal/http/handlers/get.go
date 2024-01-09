package handlers

import (
	"net/http"

	"github.com/AdrianWR/inspektor/internal/inspection/model"
)

func (s service) Get() http.HandlerFunc {
	type response struct {
		model.App
	}

	return func(w http.ResponseWriter, r *http.Request) {
		getResponse, err := s.inspectionService.Get(r.Context())
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		s.respond(w, response{
			App: getResponse,
		}, http.StatusOK)

	}
}
