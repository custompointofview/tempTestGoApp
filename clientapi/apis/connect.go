package apis

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func GRPCConnect(ctx context.Context, endpoint string) (*grpc.ClientConn, error) {
	done := time.After(30 * time.Second)
	tick := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil, nil
		case <-done:
			return nil, fmt.Errorf("connection to endpoint was not established: %v", endpoint)
		case <-tick.C:
			conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
			if err != nil {
				break
			}
			log.Printf("Connected to endpoint @: %s", endpoint)
			return conn, nil
		}
	}
}
