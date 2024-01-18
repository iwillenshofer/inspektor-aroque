package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/zap"

	"inspektor/internal/config"
)

type Server struct {
	logger *zap.SugaredLogger
	router *mux.Router
	config config.Config
}

func NewServer() (*Server, error) {
	cnf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	logger := NewLogger()
	router := mux.NewRouter()
	RegisterRoutes(router, logger)

	return &Server{
		logger: logger,
		router: router,
		config: cnf,
	}, nil
}

func (s *Server) Run(ctx context.Context) error {
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.Server.Port),
		Handler: cors.Default().Handler(s.router),
	}

	stopServer := make(chan os.Signal, 1)
	signal.Notify(stopServer, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(stopServer)

	serverErrors := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		s.logger.Infof("Server listening on port %v", s.config.Server.Port)
		serverErrors <- server.ListenAndServe()
	}(&wg)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %v", err)
	case <-stopServer:
		s.logger.Warn("server received shutdown signal")
		err := server.Shutdown(ctx)
		if err != nil {
			s.logger.Errorf("server shutdown error: %v", err)
		}
		wg.Wait()
		s.logger.Info("server shutdown completed")
	}
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
