# Hướng dẫn Setup VietQR gRPC Server

## Bước 1: Cài đặt Protocol Buffers Compiler (protoc)

### Windows

**Cách 1: Download trực tiếp**
1. Tải protoc từ: https://github.com/protocolbuffers/protobuf/releases
2. Tải file `protoc-XX.X-win64.zip` (phiên bản mới nhất)
3. Giải nén vào thư mục (ví dụ: `C:\protoc`)
4. Thêm `C:\protoc\bin` vào PATH:
   - Mở "Environment Variables"
   - Chỉnh sửa biến "Path"
   - Thêm đường dẫn `C:\protoc\bin`
5. Mở terminal mới và kiểm tra: `protoc --version`

**Cách 2: Sử dụng Chocolatey**
```bash
choco install protoc
```

**Cách 3: Sử dụng Scoop**
```bash
scoop install protobuf
```

### macOS

```bash
brew install protobuf
```

### Linux (Ubuntu/Debian)

```bash
# Cài đặt từ apt
sudo apt update
sudo apt install -y protobuf-compiler

# Hoặc download binary mới nhất
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v25.1/protoc-25.1-linux-x86_64.zip
unzip protoc-25.1-linux-x86_64.zip -d $HOME/.local
export PATH="$PATH:$HOME/.local/bin"
```

## Bước 2: Cài đặt Go plugins cho protoc

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

**Lưu ý:** Đảm bảo `$GOPATH/bin` hoặc `$HOME/go/bin` đã được thêm vào PATH.

Kiểm tra:
```bash
which protoc-gen-go
which protoc-gen-go-grpc
```

## Bước 3: Generate Protobuf Code

### Windows
```bash
generate.bat
```

### Linux/macOS
```bash
chmod +x generate.sh
./generate.sh
```

### Hoặc sử dụng Makefile
```bash
make proto
```

## Bước 4: Cài đặt Go Dependencies

```bash
go mod tidy
```

## Bước 5: Chạy gRPC Server

### Chạy trực tiếp
```bash
go run cmd/grpc-server/main.go
```

### Build và chạy
```bash
make build
./bin/grpc-server
```

### Sử dụng Docker
```bash
docker-compose up -d
```

## Kiểm tra

### Kiểm tra server đang chạy

```bash
# Sử dụng grpcurl (cần cài đặt trước)
grpcurl -plaintext localhost:50051 list
```

### Test Encode
```bash
grpcurl -plaintext -d '{
  "bank_code": "TECHCOMBANK",
  "bank_no": "9796868",
  "amount": 50000,
  "message": "Test payment"
}' localhost:50051 vietqr.VietQRService/Encode
```

## Troubleshooting

### Lỗi: protoc: command not found
- Đảm bảo protoc đã được cài đặt và thêm vào PATH
- Mở terminal mới sau khi cài đặt

### Lỗi: protoc-gen-go: program not found or is not executable
- Chạy: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
- Kiểm tra PATH đã có `$GOPATH/bin`
- Windows: Thêm `%USERPROFILE%\go\bin` vào PATH
- Linux/Mac: Thêm `$HOME/go/bin` vào PATH

### Lỗi: no matching versions for query "latest"
- Generate proto code trước: `generate.bat` hoặc `./generate.sh`
- Sau đó chạy: `go mod tidy`

### Port 50051 already in use
- Thay đổi port: `GRPC_PORT=9090 go run cmd/grpc-server/main.go`
- Hoặc kill process đang sử dụng port đó

## Cấu trúc Files sau khi setup

```
vietqr/
├── proto/
│   ├── vietqr.proto          # Proto definition
│   ├── vietqr.pb.go          # Generated (sau khi chạy generate)
│   └── vietqr_grpc.pb.go     # Generated (sau khi chạy generate)
├── cmd/grpc-server/
│   └── main.go               # gRPC server
├── go.mod
├── go.sum                    # Generated sau go mod tidy
└── ...
```

## Tích hợp với NestJS

Xem file `README_GRPC.md` để biết chi tiết cách tích hợp với NestJS.

## Liên hệ

Nếu gặp vấn đề, vui lòng tạo issue tại: https://github.com/sunary/vietqr/issues

