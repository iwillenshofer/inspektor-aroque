package main

import (
	"context"
	"flag"

	"github.com/AdrianWR/inspektor/internal/gateway"
	"google.golang.org/grpc/grpclog"
)

var (
	endpoint = flag.String("endpoint", "localhost:9090", "gRPC server endpoint")
	network  = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent with -endpoint`)
)

func main() {
	flag.Parse()

	ctx := context.Background()
	opts := gateway.Options{
		Addr: ":8080",
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
	}
	if err := gateway.Run(ctx, opts); err != nil {
		grpclog.Fatal(err)
	}
}
