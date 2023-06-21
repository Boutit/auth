SERVER_PATH="cmd/server/main.go"
PROTO_RELATIVE_PATH="api/protos/boutit/auth"
ENV="local"

codegen.proto:
	protoc $(PROTO_RELATIVE_PATH)/*.proto --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false \
	--grpc-gateway_out=. --grpc-gateway_opt paths=source_relative --proto_path=.

grpcui:
	grpcui -plaintext localhost:8080

local:
	ENV=local go run cmd/server/main.go

development:
	ENV=development go run cmd/server/main.go

staging:
	ENV=staging go run cmd/server/main.go

production:
	ENV=production go run cmd/server/main.go