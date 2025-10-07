# vietqr

This Go-based library enables encoding and decoding for VietQR, a standardized QR code solution for payment and transfer services across Vietnam's NAPAS network. Once encoded, the QR code can be further customized using any QR code generator that converts text to QR format.

VietQR serves as a unified brand identity for QR-based payments and transfers, seamlessly processed through the NAPAS network, member banks, payment intermediaries, and domestic and international partners. It complies with the EMV Co. QR payment standards and adheres to foundational QR code standards established by the State Bank of Vietnam.

## Sample

```go
package main

import (
	"github.com/sunary/vietqr"
)

func main() {
	println("encode", vietqr.Encode(vietqr.TransferInfo{
		BankCode: vietqr.TECHCOMBANK,
		BankNo:   "9796868",
		Message:  "gen by sunary/vietqr",
	}))

	info, err := vietqr.Decode("00020101021138510010A00000072701210006970407010797968680208QRIBFTTA53037045802VN62240820gen by sunary/vietqr6304BE74")
	if err != nil {
		println("err", err.Error())
		return
	}

	println("bank_code", info.BankCode)
	println("bank_no", info.BankNo)
	println("amount", info.Amount)
	println("message", info.Message)
}
```

Run online: [go.dev/play](https://go.dev/play/p/g9gCWmI9iRl)

## gRPC Server

This project now includes a gRPC server for easy integration with other services like NestJS, Python, Java, etc.

### Quick Start

**Windows:**
```bash
# 1. Install protoc (one-time setup)
choco install protoc
# or download from: https://github.com/protocolbuffers/protobuf/releases

# 2. Run setup script
powershell -ExecutionPolicy Bypass -File setup-windows.ps1

# 3. Start gRPC server
go run cmd/grpc-server/main.go
```

**Linux/macOS:**
```bash
# 1. Install protoc (one-time setup)
brew install protobuf  # macOS
# or
sudo apt install protobuf-compiler  # Linux

# 2. Setup and run
./generate.sh && go mod tidy && go run cmd/grpc-server/main.go
```

**Docker:**
```bash
docker-compose up -d
```

Server will run on `localhost:50051`

### gRPC API Methods

**1. Encode** - Generate VietQR code
```protobuf
rpc Encode(EncodeRequest) returns (EncodeResponse);
```

**2. Decode** - Parse VietQR code
```protobuf
rpc Decode(DecodeRequest) returns (DecodeResponse);
```

**3. GetBankList** - Get supported banks
```protobuf
rpc GetBankList(GetBankListRequest) returns (GetBankListResponse);
```

### Test with grpcurl

```bash
# Install grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# Test Encode
grpcurl -plaintext -d '{
  "bank_code": "TECHCOMBANK",
  "bank_no": "9796868",
  "amount": 50000,
  "message": "Payment"
}' localhost:50051 vietqr.VietQRService/Encode

# Test Decode
grpcurl -plaintext -d '{
  "qr_code": "00020101021138510010A00000072701210006970407010797968680208QRIBFTTA53037045802VN62240820gen by sunary/vietqr6304BE74"
}' localhost:50051 vietqr.VietQRService/Decode

# Get bank list
grpcurl -plaintext localhost:50051 vietqr.VietQRService/GetBankList
```

### NestJS Integration

Full NestJS integration guide available in [`README_NESTJS_INTEGRATION.md`](README_NESTJS_INTEGRATION.md)

Quick example:
```typescript
import { ClientsModule, Transport } from '@nestjs/microservices';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'VIETQR_PACKAGE',
        transport: Transport.GRPC,
        options: {
          package: 'vietqr',
          protoPath: join(__dirname, './proto/vietqr.proto'),
          url: 'localhost:50051',
        },
      },
    ]),
  ],
})
export class AppModule {}
```

### Documentation

- 📖 **[QUICKSTART.md](QUICKSTART.md)** - Get started in 5 minutes
- 🔧 **[SETUP.md](SETUP.md)** - Detailed setup instructions
- 📡 **[README_GRPC.md](README_GRPC.md)** - Complete gRPC documentation
- 🔗 **[README_NESTJS_INTEGRATION.md](README_NESTJS_INTEGRATION.md)** - NestJS integration guide
- 💻 **[examples/client](examples/client)** - Go client example

### Features

✅ High-performance gRPC server  
✅ Support all Vietnamese banks (67+ banks)  
✅ Encode & Decode VietQR codes  
✅ Dynamic QR (with or without amount)  
✅ Easy integration with any language  
✅ Docker support  
✅ Production-ready  
✅ Comprehensive documentation