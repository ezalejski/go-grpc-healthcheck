package main

import (
	"log"

	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	const httpPort string = ":443"
	grpcServer := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	// Create a server on httpPort
	httpServer := &http.Server{Addr: httpPort, Handler: grpcServer}
	// Start the server with TLS, since we are running HTTP/2 it must be
	// run with TLS.
	log.Printf("Serving on https://0.0.0.0%s", httpPort)
	log.Fatal(httpServer.ListenAndServeTLS("/tls/tls.crt", "/tls/tls.key"))
}
