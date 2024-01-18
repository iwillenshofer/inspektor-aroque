package service

import (
	"context"
	"os"

	"inspektor/internal/inspection/model"
)

func getPodName() string {
	return os.Getenv("HOSTNAME")
}

func getPodNamespace() string {
	namespace, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		return ""
	}
	return string(namespace)
}

func getPodIP() string {
	return os.Getenv("KUBERNETES_SERVICE_HOST")
}

func (s Service) Get(ctx context.Context) (model.App, error) {
	return model.App{
		Name:    "inspektor",
		Version: "v2",
		Pod: model.Pod{
			Name:      getPodName(),
			Namespace: getPodNamespace(),
			IP:        getPodIP(),
		},
	}, nil
}
