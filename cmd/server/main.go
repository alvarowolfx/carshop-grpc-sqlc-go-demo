package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	carshop "com.aviebrantz.carshop/api"
	"com.aviebrantz.carshop/pkg/controllers"
	"com.aviebrantz.carshop/pkg/database"
	"com.aviebrantz.carshop/pkg/repository"
	"github.com/gobuffalo/packr/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	restPort := os.Getenv("REST_PORT")
	grpcPort := os.Getenv("GRPC_PORT")
	grpcAddr := fmt.Sprintf("localhost:%s", grpcPort)

	migrate.SetTable("migrations")
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "../../assets/migrations"),
	}

	db := database.MustConnectPostgres()
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Failed to apply db migrations: %v\n", err)
	}
	log.Printf("Applied %d migrations!\n", n)

	repo := repository.New(db)

	backOfficeController := controllers.NewBackOfficeController(controllers.BackOfficeControllerDeps{
		DB: repo,
	})
	workOrderController := controllers.NewWorkOrderController(controllers.WorkOrderControllerDeps{
		DB: repo,
	})

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	carshop.RegisterWorkOrderServiceServer(s, workOrderController)
	carshop.RegisterBackOfficeServiceServer(s, backOfficeController)

	go func() {
		if err := s.Serve(conn); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = carshop.RegisterWorkOrderServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v\n", err)
		return
	}
	err = carshop.RegisterBackOfficeServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v\n", err)
		return
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", restPort), mux))
}
