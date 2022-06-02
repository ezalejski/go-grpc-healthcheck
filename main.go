package main

import (
	"fmt"
	"log"
	"net"

	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	const grpcPort string = ":50051"
	const httpPort string = ":8080"
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal("failed-to-listen-tcp-port: " + err.Error())
	}
	// Create a GRPC server
	grpcServer := grpc.NewServer()
	// Register healthcheck
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	// Start a GRPC server on grpcPort
	go grpcServer.Serve(listener)

	// Create a http server on httpPort
	httpServer := &http.Server{Addr: httpPort, Handler: grpcServer}
	// Start the server with TLS, since we are running HTTP/2 it must be
	// run with TLS.
	log.Printf("Serving on https://0.0.0.0%s", httpPort)
	log.Fatal(httpServer.ListenAndServeTLS("/tls/tls.crt", "/tls/tls.key"))
}
