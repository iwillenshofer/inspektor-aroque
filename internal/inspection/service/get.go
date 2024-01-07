package service

import (
	"context"
	"os"

	"github.com/AdrianWR/inspektor/internal/inspection/model"
)

func getPodName() string {
	return os.Getenv("POD_NAME")
}

func getPodNamespace() string {
	return os.Getenv("POD_NAMESPACE")
}

func getPodIP() string {
	return os.Getenv("POD_IP")
}

func (s Service) Get(ctx context.Context) (model.App, error) {
	return model.App{
		Name:    "inspektor",
		Version: "0.1.0",
		Pod: model.Pod{
			Name:      getPodName(),
			Namespace: getPodNamespace(),
			IP:        getPodIP(),
		},
	}, nil
}
