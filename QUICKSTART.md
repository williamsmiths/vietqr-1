# Quick Start - VietQR gRPC

Hướng dẫn nhanh để chạy VietQR gRPC server trong 5 phút.

## Yêu cầu trước khi bắt đầu

- Go 1.21.5+
- protoc (Protocol Buffers Compiler)

## Cài đặt nhanh trên Windows

### Bước 1: Cài đặt protoc

Download và cài đặt protoc từ: https://github.com/protocolbuffers/protobuf/releases

Hoặc dùng Chocolatey:
```bash
choco install protoc
```

### Bước 2: Chạy setup script

```powershell
powershell -ExecutionPolicy Bypass -File setup-windows.ps1
```

### Bước 3: Chạy server

```bash
go run cmd/grpc-server/main.go
```

✅ Server đã chạy trên port 50051!

## Cài đặt nhanh trên Linux/macOS

### Một lệnh để setup tất cả:

```bash
# macOS
brew install protobuf && \
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
./generate.sh && \
go mod tidy && \
go run cmd/grpc-server/main.go

# Linux
sudo apt install -y protobuf-compiler && \
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
chmod +x generate.sh && \
./generate.sh && \
go mod tidy && \
go run cmd/grpc-server/main.go
```

## Test nhanh với grpcurl

### Cài đặt grpcurl

```bash
# macOS
brew install grpcurl

# Windows (Chocolatey)
choco install grpcurl

# Go
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
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

### Test Decode

```bash
grpcurl -plaintext -d '{
  "qr_code": "00020101021138510010A00000072701210006970407010797968680208QRIBFTTA53037045802VN62240820gen by sunary/vietqr6304BE74"
}' localhost:50051 vietqr.VietQRService/Decode
```

### Lấy danh sách ngân hàng

```bash
grpcurl -plaintext localhost:50051 vietqr.VietQRService/GetBankList
```

## Chạy với Docker (Không cần cài protoc)

```bash
# Build
docker-compose build

# Run
docker-compose up -d

# Test
grpcurl -plaintext localhost:50051 list
```

## Các lệnh Make hữu ích

```bash
make proto       # Generate protobuf code
make run-server  # Run gRPC server
make build       # Build binary
make clean       # Clean generated files
make help        # Show all commands
```

## Port Configuration

Mặc định: `50051`

Thay đổi port:
```bash
GRPC_PORT=9090 go run cmd/grpc-server/main.go
```

## Tích hợp với NestJS

Xem hướng dẫn chi tiết tại: `README_NESTJS_INTEGRATION.md`

Nhanh:
1. Copy file `proto/vietqr.proto` vào NestJS project
2. Cài đặt: `npm install @grpc/grpc-js @grpc/proto-loader @nestjs/microservices`
3. Tạo gRPC client (xem example trong README_NESTJS_INTEGRATION.md)

## Troubleshooting

### "protoc: command not found"
→ Cài đặt protoc (xem Bước 1)

### "protoc-gen-go: program not found"
→ Chạy: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

### "no matching versions for query latest"
→ Generate proto trước: `./generate.sh` hoặc `generate.bat`

### Port already in use
→ Đổi port: `GRPC_PORT=9090 go run cmd/grpc-server/main.go`

## Danh sách ngân hàng được hỗ trợ

- TECHCOMBANK
- VIETCOMBANK
- BIDV
- VIETINBANK
- ACB
- MBBANK
- VPBANK
- TPBank
- Sacombank
- VIB
- ...và 50+ ngân hàng khác

Xem file `banks.go` để biết danh sách đầy đủ.

## Ví dụ sử dụng trong Go code

```go
package main

import (
    "context"
    "log"
    
    "github.com/sunary/vietqr"
    pb "github.com/sunary/vietqr/proto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func main() {
    // Kết nối tới server
    conn, err := grpc.Dial("localhost:50051", 
        grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    client := pb.NewVietQRServiceClient(conn)
    
    // Encode QR
    encResp, err := client.Encode(context.Background(), &pb.EncodeRequest{
        BankCode: vietqr.TECHCOMBANK,
        BankNo:   "9796868",
        Amount:   50000,
        Message:  "Test payment",
    })
    if err != nil {
        log.Fatal(err)
    }
    log.Println("QR Code:", encResp.QrCode)
    
    // Decode QR
    decResp, err := client.Decode(context.Background(), &pb.DecodeRequest{
        QrCode: encResp.QrCode,
    })
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Bank: %s, Account: %s, Amount: %d\n", 
        decResp.BankCode, decResp.BankNo, decResp.Amount)
}
```

## Next Steps

1. ✅ Server đã chạy
2. 📖 Đọc `README_GRPC.md` để hiểu chi tiết
3. 🔗 Tích hợp với NestJS: `README_NESTJS_INTEGRATION.md`
4. 📝 Xem API docs: `proto/vietqr.proto`

## Support

- GitHub: https://github.com/sunary/vietqr
- Issues: https://github.com/sunary/vietqr/issues

---

**Chúc bạn coding vui vẻ! 🚀**

