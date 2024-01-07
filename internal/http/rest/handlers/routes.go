package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func Register(r *mux.Router, lg *zap.SugaredLogger) {
	handler := newHandler(lg)
	r.HandleFunc("/healthz", handler.Health()).Methods(http.MethodGet)
	r.HandleFunc("/inspect", handler.Get()).Methods(http.MethodGet)
}
