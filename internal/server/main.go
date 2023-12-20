package server

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	appv1 "github.com/AdrianWR/inspektor/gen/go/inspection/v1"
)

func Run(ctx context.Context, network string, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			grpclog.Errorf("failed to close listener %s %s: %v", network, address, err)
		}
	}()

	s := grpc.NewServer()
	appv1.RegisterInspectionServiceServer(s, newInspectionServer())

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()

	grpclog.Infof("starting gRPC server at %s %s", network, address)
	return s.Serve(l)
}
