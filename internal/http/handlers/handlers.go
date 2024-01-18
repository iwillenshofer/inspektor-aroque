package handlers

import (
	inspectionService "inspektor/internal/inspection/service"
	"go.uber.org/zap"
)

type service struct {
	logger            *zap.SugaredLogger
	inspectionService inspectionService.Service
}

func NewHandler(lg *zap.SugaredLogger) service {
	return service{
		logger:            lg,
		inspectionService: inspectionService.NewService(),
	}
}
