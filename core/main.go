package main

import (
	"log"
	"net"
	"runtime"

	pb "modcore/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Create appropriate listener for the platform
	lis, err := createListener()
	if err != nil {
		log.Fatalf("Failed to create listener: %v", err)
	}
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			log.Fatalf("Failed to close listener: %v", err)
		}
	}(lis)

	// Clean up socket on exit
	defer cleanupSocket()

	// Create gRPC server
	s := grpc.NewServer()

	// Register our service
	pb.RegisterModCoreServer(s, NewModCoreServer())

	// Enable reflection for debugging
	reflection.Register(s)

	log.Printf("ModCore gRPC server starting on %s platform", runtime.GOOS)
	log.Printf("Socket path: %s", getSocketPath())

	// Start serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
