package server

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"com.aviebrantz.carshop/pkg/auth"
	backOffice "com.aviebrantz.carshop/pkg/backoffice/controllers"
	carshop "com.aviebrantz.carshop/pkg/common/api"
	"com.aviebrantz.carshop/pkg/common/database"
	"com.aviebrantz.carshop/pkg/common/repository"
	workOrder "com.aviebrantz.carshop/pkg/workorder/controllers"
	"github.com/gobuffalo/packr/v2"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func runMigrations(db *sql.DB) {
	migrate.SetTable("migrations")
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "../../db/migrations"),
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Failed to apply db migrations: %v\n", err)
	}
	log.Printf("Applied %d migrations!\n", n)
}

func Run() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	db := database.MustConnectPostgres()
	runMigrations(db)
	repo := repository.New(db)

	grpcServerOpts := []grpc.ServerOption{}
	authProviderType := os.Getenv("AUTH_PROVIDER")
	var authProvider auth.AuthProvider
	if authProviderType == "auth0" {
		identifier := os.Getenv("AUTH0_IDENTIFIER")
		issuer := os.Getenv("AUTH0_ISSUER")
		jwksURI := os.Getenv("AUTH0_JWKS_URI")
		authProvider = &auth.Auth0Provider{
			Identifier: identifier,
			Issuer:     issuer,
			JwksURI:    jwksURI,
		}
		authFunc := auth.VerifyFuncForProvider(authProvider)
		grpcServerOpts = append(
			grpcServerOpts,
			grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authFunc)),
		)
	}

	backOfficeController := backOffice.NewController(backOffice.ControllerDeps{
		DB: repo,
	})
	workOrderController := workOrder.NewController(workOrder.ControllerDeps{
		DB:           repo,
		AuthProvider: authProvider,
	})

	grpcServer := grpc.NewServer(grpcServerOpts...)
	carshop.RegisterWorkOrderServiceServer(grpcServer, workOrderController)
	carshop.RegisterBackOfficeServiceServer(grpcServer, backOfficeController)

	// Start HTTP server that serves both Rest Gateway and GRPC
	mux := http.NewServeMux()

	// Register gRPC server endpoint
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	ctx := context.Background()
	err = carshop.RegisterWorkOrderServiceHandlerFromEndpoint(ctx, gwmux, addr, opts)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v\n", err)
		return err
	}
	err = carshop.RegisterBackOfficeServiceHandlerFromEndpoint(ctx, gwmux, addr, opts)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v\n", err)
		return err
	}

	mux.Handle("/", gwmux)
	swaggerBox := packr.New("swagger.json", "../../api")
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		swagger, err := swaggerBox.FindString("carshop.swagger.json")
		if err != nil {
			w.WriteHeader(404)
			io.WriteString(w, "Swagger file not found")
			return
		}
		io.Copy(w, strings.NewReader(swagger))
	})
	muxWithCors := allowCORS(mux)

	return http.ListenAndServe(
		addr,
		grpcHandlerFunc(grpcServer, muxWithCors),
	)
}
