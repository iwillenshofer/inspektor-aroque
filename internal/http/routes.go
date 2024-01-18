package http

import (
	"net/http"

	"inspektor/internal/http/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func RegisterRoutes(r *mux.Router, lg *zap.SugaredLogger) {
	handler := handlers.NewHandler(lg)

	// register root route am calll inspect
    r.HandleFunc("/", handler.Get()).Methods(http.MethodGet)

	s := r.PathPrefix("/v1").Subrouter()
	s.HandleFunc("/healthz", handler.Health()).Methods(http.MethodGet)
	s.HandleFunc("/inspect", handler.Get()).Methods(http.MethodGet)
}
