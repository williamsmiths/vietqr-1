.PHONY: proto run-server clean help

# Generate protobuf code
proto:
	@echo "Generating protobuf code..."
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/vietqr.proto
	@echo "Protobuf code generated successfully!"

# Run gRPC server
run-server:
	@echo "Starting gRPC server..."
	go run cmd/grpc-server/main.go

# Build gRPC server
build:
	@echo "Building gRPC server..."
	go build -o bin/grpc-server cmd/grpc-server/main.go
	@echo "Build complete: bin/grpc-server"

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	rm -f proto/*.pb.go
	rm -rf bin/
	@echo "Cleaned!"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	@echo "Dependencies installed!"

# Help
help:
	@echo "Available commands:"
	@echo "  make proto       - Generate protobuf code"
	@echo "  make run-server  - Run gRPC server"
	@echo "  make build       - Build gRPC server binary"
	@echo "  make deps        - Install dependencies"
	@echo "  make clean       - Clean generated files"
	@echo "  make help        - Show this help message"

