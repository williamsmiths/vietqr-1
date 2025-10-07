# Stage 1: Build stage
FROM golang:1.24-alpine AS builder

# Cài đặt các dependencies cần thiết
RUN apk add --no-cache git ca-certificates tzdata protobuf protobuf-dev

# Tạo thư mục làm việc
WORKDIR /app

# Copy go mod files trước để cache dependencies
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Install protobuf plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Copy toàn bộ source code
COPY . .

# Generate protobuf code
RUN protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/vietqr.proto

# Build gRPC server binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o grpc-server cmd/grpc-server/main.go

# Stage 2: Runtime stage
FROM alpine:latest

# Cài đặt ca-certificates
RUN apk --no-cache add ca-certificates tzdata

# Tạo user non-root
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Tạo thư mục làm việc
WORKDIR /app

# Copy binary từ builder stage
COPY --from=builder /app/grpc-server .

# Chuyển ownership
RUN chown -R appuser:appgroup /app

# Chuyển sang user non-root
USER appuser

# Expose gRPC port
EXPOSE 9090

# Chạy gRPC server
CMD ["./grpc-server"]
