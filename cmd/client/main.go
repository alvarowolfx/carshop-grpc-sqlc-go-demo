package main

import (
	"context"
	"fmt"
	"log"
	"os"

	carshop "com.aviebrantz.carshop/api"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	grpcPort := os.Getenv("GRPC_PORT")
	grpcAddr := fmt.Sprintf(":%s", grpcPort)

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	ctx := context.Background()
	client := carshop.NewWorkOrderServiceClient(conn)
	response, err := client.GetRunningWorkOrders(ctx, &carshop.RunningWorkOrdersQuery{})
	if err != nil {
		log.Fatalf("failed to fetch running work orders: %v", err)
	}

	for _, wo := range response.WorkOrder {
		log.Println(wo.LicensePlate, wo.Status, wo.PreviousStatus)
	}
}
