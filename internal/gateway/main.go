package gateway

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/grpclog"
)

// Endpoint is a network address of a gRPC server
type Endpoint struct {
	Network, Addr string
}

// Options is a gateway server configuration
type Options struct {
	// Addr is a network address of the gateway server
	Addr string
	// GRPCServer defines a gRPC server endpoint
	GRPCServer Endpoint
	// Mux is a list of options to configure the gRPC gateway multiplexer
	Mux []runtime.ServeMuxOption
}

// Run starts a HTTP server and blocks until stop signal is received.
// The server will be stopped gracefully when the context is canceled.
func Run(ctx context.Context, opts Options) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := dial(ctx, opts.GRPCServer.Network, opts.GRPCServer.Addr)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			grpclog.Infof("failed to close connection: %v", err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthzServer(conn))

	gw, err := newGateway(ctx, conn, opts.Mux)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	srv := &http.Server{
		Addr:    opts.Addr,
		Handler: mux,
	}
	go func() {
		<-ctx.Done()
		grpclog.Infof("shutting down gateway server at %s", opts.Addr)
		if err := srv.Shutdown(context.Background()); err != nil {
			grpclog.Infof("failed to shutdown gateway server: %v", err)
		}
	}()

	grpclog.Infof("starting gateway server at %s", opts.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		grpclog.Errorf("failed to listen and serve: %v", err)
		return err
	}

	return nil
}
