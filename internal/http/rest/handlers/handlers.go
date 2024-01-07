package handlers

import (
	inspectionService "github.com/AdrianWR/inspektor/internal/inspection/service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type service struct {
	logger            *zap.SugaredLogger
	router            *mux.Router
	inspectionService inspectionService.Service
}

func newHandler(lg *zap.SugaredLogger) service {
	return service{
		logger:            lg,
		inspectionService: inspectionService.NewService(),
	}
}
