package server

import (
	"context"

	pb "github.com/AdrianWR/inspektor/gen/go/inspection/v1"
	"google.golang.org/grpc/grpclog"
)

type inspectionServer struct {
	pb.UnimplementedInspectionServiceServer
}

func newInspectionServer() pb.InspectionServiceServer {
	return new(inspectionServer)
}

func (s *inspectionServer) GetInspection(ctx context.Context, _ *pb.GetInspectionRequest) (*pb.GetInspectionResponse, error) {
	grpclog.Infof("GetInspection")
	return &pb.GetInspectionResponse{
		App: &pb.App{Name: "test", Version: "1.0.0", Pod: &pb.Pod{Name: "test", Namespace: "test", Ip: "0.0.0.0"}},
	}, nil
}
