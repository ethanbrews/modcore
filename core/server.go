package main

import (
	"context"
	pb "modcore/proto/gen"
)

// ModCoreServer implements the ModCore service
type ModCoreServer struct {
	pb.UnimplementedModCoreServer
}

// CoreInfo returns build information about the core service
func (s *ModCoreServer) CoreInfo(ctx context.Context, req *pb.Empty) (*pb.BuildInfoResponse, error) {
	return &pb.BuildInfoResponse{
		CoreVersion: "1.0.0",
		Build:       "development",
		ApiVersion:  "v1",
	}, nil
}

// NewModCoreServer creates a new instance of the ModCore server
func NewModCoreServer() *ModCoreServer {
	return &ModCoreServer{}
}
