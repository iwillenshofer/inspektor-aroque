package gateway

import (
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

func healthzServer(conn *grpc.ClientConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if s := conn.GetState(); s != connectivity.Ready {
			http.Error(w, `{"status": "unhealthy"}`, http.StatusBadGateway)
			return
		}
		w.Write([]byte(`{"status": "healthy"}`))
	}
}
