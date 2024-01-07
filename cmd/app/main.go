package main

import (
	"context"
	"log"

	"github.com/AdrianWR/inspektor/internal/http/rest"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("error: %+v", err)
	}
}

func run(ctx context.Context) error {
	server, err := rest.NewServer()
	if err != nil {
		return err
	}
	err = server.Run(ctx)
	return err
}
