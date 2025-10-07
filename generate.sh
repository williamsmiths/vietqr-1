#!/bin/bash

# Script để generate protobuf code

echo "Generating protobuf code for VietQR gRPC..."

# Kiểm tra protoc đã được cài đặt chưa
if ! command -v protoc &> /dev/null
then
    echo "protoc chưa được cài đặt. Vui lòng cài đặt Protocol Buffers compiler."
    echo "Tải tại: https://grpc.io/docs/protoc-installation/"
    exit 1
fi

# Kiểm tra Go protobuf plugins
if ! command -v protoc-gen-go &> /dev/null
then
    echo "Installing protoc-gen-go..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
fi

if ! command -v protoc-gen-go-grpc &> /dev/null
then
    echo "Installing protoc-gen-go-grpc..."
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
fi

# Generate proto files
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/vietqr.proto

if [ $? -eq 0 ]; then
    echo "✓ Protobuf code generated successfully!"
else
    echo "✗ Failed to generate protobuf code"
    exit 1
fi

