SERVER_PATH="cmd/server/main.go"
PROTO_RELATIVE_PATH="api"
ENV="local"

codegen.proto:
	protoc $(PROTO_RELATIVE_PATH)/*.proto --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false \
	--grpc-gateway_out=. --grpc-gateway_opt paths=source_relative --proto_path=.

run.local:
	ENV=local go run cmd/server/main.go

run.development:
	ENV=development go run cmd/server/main.go

run.staging:
	ENV=staging go run cmd/server/main.go

run.production:
	ENV=production go run cmd/server/main.go