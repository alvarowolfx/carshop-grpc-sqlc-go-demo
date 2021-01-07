protoc-generate:
	find api -iname "*.proto" | xargs -I@ echo --path=@ | xargs buf generate


sqlc-generate:
	sqlc generate

regenerate: sqlc-generate protoc-generate

