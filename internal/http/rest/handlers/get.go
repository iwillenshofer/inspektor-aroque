package handlers

import "net/http"

func (s service) Get() http.HandlerFunc {
	type response struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Pod     struct {
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
			IP        string `json:"ip"`
		} `json:"pod"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		getResponse, err := s.inspectionService.Get(r.Context())
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		s.respond(w, response{
			Name:    getResponse.Name,
			Version: getResponse.Version,
			Pod: struct {
				Name      string `json:"name"`
				Namespace string `json:"namespace"`
				IP        string `json:"ip"`
			}{
				Name:      getResponse.Pod.Name,
				Namespace: getResponse.Pod.Namespace,
				IP:        getResponse.Pod.IP,
			},
		}, http.StatusOK)

	}
}
