package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	backOffice "com.aviebrantz.carshop/pkg/backoffice/controllers"
	carshop "com.aviebrantz.carshop/pkg/common/api"
	"com.aviebrantz.carshop/pkg/common/database"
	"com.aviebrantz.carshop/pkg/common/repository"
	workOrder "com.aviebrantz.carshop/pkg/workorder/controllers"
	"github.com/gobuffalo/packr/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
	"google.golang.org/grpc"
)

func Run() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	restPort := os.Getenv("REST_PORT")
	grpcPort := os.Getenv("GRPC_PORT")
	grpcAddr := fmt.Sprintf("localhost:%s", grpcPort)

	migrate.SetTable("migrations")
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "../../db/migrations"),
	}

	db := database.MustConnectPostgres()
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Failed to apply db migrations: %v\n", err)
	}
	log.Printf("Applied %d migrations!\n", n)

	repo := repository.New(db)

	backOfficeController := backOffice.NewController(backOffice.ControllerDeps{
		DB: repo,
	})
	workOrderController := workOrder.NewController(workOrder.ControllerDeps{
		DB: repo,
	})

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
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
	muxWithCors := allowCORS(mux)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = carshop.RegisterWorkOrderServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v\n", err)
		return err
	}
	err = carshop.RegisterBackOfficeServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v\n", err)
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(fmt.Sprintf(":%s", restPort), muxWithCors)
}
