package main

import (
	"context"
	"log"
	"modcore/cli/ipc"
	"time"

	pb "modcore/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		ipc.getSocketPath(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(ipc.SocketDialer()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}(conn)

	client := pb.NewModCoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
 
	resp, err := client.CoreInfo(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response: %s ; %s ; %s", resp.ApiVersion, resp.CoreVersion, resp.Build)

}
