protoc-generate:
	protoc -I ./api \
  --go_out ./api --go_opt paths=source_relative \
  --go-grpc_out ./api --go-grpc_opt paths=source_relative --go-grpc_opt require_unimplemented_servers=false \
  --grpc-gateway_out ./api \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  --openapiv2_out ./api \
  --openapiv2_opt logtostderr=true \
  ./api/carshop.proto

sqlc-generate:
	sqlc generate

regenerate: sqlc-generate protoc-generate

