# VietQR gRPC Service

Service gRPC để tạo và giải mã mã QR thanh toán VietQR, hỗ trợ giao tiếp với NestJS và các service khác.

## Yêu cầu

- Go 1.21.5 trở lên
- Protocol Buffers compiler (`protoc`)
- Go plugins cho protoc:
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`

## Cài đặt

### 1. Cài đặt Protocol Buffers Compiler

**Windows:**
```bash
# Tải từ: https://github.com/protocolbuffers/protobuf/releases
# Giải nén và thêm vào PATH
```

**macOS:**
```bash
brew install protobuf
```

**Linux:**
```bash
sudo apt install -y protobuf-compiler
```

### 2. Cài đặt Go plugins

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 3. Cài đặt dependencies

```bash
go mod tidy
```

## Generate Protobuf Code

**Windows:**
```bash
generate.bat
```

**Linux/macOS:**
```bash
chmod +x generate.sh
./generate.sh
```

Hoặc sử dụng Makefile:
```bash
make proto
```

## Chạy gRPC Server

```bash
# Sử dụng go run
go run cmd/grpc-server/main.go

# Hoặc sử dụng Makefile
make run-server

# Hoặc build và chạy
make build
./bin/grpc-server
```

Server sẽ chạy trên port `9090` (mặc định). Có thể thay đổi bằng biến môi trường:

```bash
GRPC_PORT=9090 go run cmd/grpc-server/main.go
```

## API Methods

### 1. Encode - Tạo mã QR

**Request:**
```protobuf
message EncodeRequest {
  string bank_code = 1;    // Mã ngân hàng (TECHCOMBANK, VIETCOMBANK, ...)
  string bank_no = 2;      // Số tài khoản
  int64 amount = 3;        // Số tiền (optional, 0 nếu không có)
  string message = 4;      // Nội dung chuyển khoản (optional)
}
```

**Response:**
```protobuf
message EncodeResponse {
  string qr_code = 1;      // Chuỗi mã QR
  bool success = 2;        // Trạng thái thành công
  string error = 3;        // Thông báo lỗi (nếu có)
}
```

### 2. Decode - Giải mã QR

**Request:**
```protobuf
message DecodeRequest {
  string qr_code = 1;      // Chuỗi mã QR cần decode
}
```

**Response:**
```protobuf
message DecodeResponse {
  string bank_code = 1;    // Mã ngân hàng
  string bank_no = 2;      // Số tài khoản
  int64 amount = 3;        // Số tiền
  string message = 4;      // Nội dung chuyển khoản
  bool success = 5;        // Trạng thái thành công
  string error = 6;        // Thông báo lỗi (nếu có)
}
```

### 3. GetBankList - Lấy danh sách ngân hàng

**Request:**
```protobuf
message GetBankListRequest {}
```

**Response:**
```protobuf
message GetBankListResponse {
  repeated BankInfo banks = 1;
  bool success = 2;
  string error = 3;
}

message BankInfo {
  string code = 1;         // Mã ngân hàng
  string name = 2;         // Tên ngân hàng
  string bin = 3;          // BIN của ngân hàng
}
```

## Test với grpcurl

### Cài đặt grpcurl

**macOS:**
```bash
brew install grpcurl
```

**Linux:**
```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

**Windows:**
```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

### Test các API

**1. List services:**
```bash
grpcurl -plaintext localhost:9090 list
```

**2. Encode QR code:**
```bash
grpcurl -plaintext -d '{
  "bank_code": "TECHCOMBANK",
  "bank_no": "9796868",
  "amount": 50000,
  "message": "Thanh toan don hang"
}' localhost:9090 vietqr.VietQRService/Encode
```

**3. Decode QR code:**
```bash
grpcurl -plaintext -d '{
  "qr_code": "00020101021138510010A00000072701210006970407010797968680208QRIBFTTA53037045802VN62240820gen by sunary/vietqr6304BE74"
}' localhost:9090 vietqr.VietQRService/Decode
```

**4. Get bank list:**
```bash
grpcurl -plaintext localhost:9090 vietqr.VietQRService/GetBankList
```

## Tích hợp với NestJS

### 1. Cài đặt packages trong NestJS

```bash
npm install @grpc/grpc-js @grpc/proto-loader
```

### 2. Copy file proto vào NestJS project

```bash
# Copy file proto/vietqr.proto vào thư mục proto/ trong NestJS project
cp proto/vietqr.proto /path/to/nestjs-project/proto/
```

### 3. Cấu hình gRPC client trong NestJS

**app.module.ts:**
```typescript
import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'VIETQR_PACKAGE',
        transport: Transport.GRPC,
        options: {
          package: 'vietqr',
          protoPath: join(__dirname, '../proto/vietqr.proto'),
          url: 'localhost:9090',
        },
      },
    ]),
  ],
})
export class AppModule {}
```

### 4. Sử dụng trong service

**vietqr.service.ts:**
```typescript
import { Injectable, Inject, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { Observable } from 'rxjs';

interface EncodeRequest {
  bank_code: string;
  bank_no: string;
  amount: number;
  message: string;
}

interface EncodeResponse {
  qr_code: string;
  success: boolean;
  error: string;
}

interface DecodeRequest {
  qr_code: string;
}

interface DecodeResponse {
  bank_code: string;
  bank_no: string;
  amount: number;
  message: string;
  success: boolean;
  error: string;
}

interface VietQRService {
  encode(data: EncodeRequest): Observable<EncodeResponse>;
  decode(data: DecodeRequest): Observable<DecodeResponse>;
  getBankList(data: {}): Observable<any>;
}

@Injectable()
export class VietQRClientService implements OnModuleInit {
  private vietqrService: VietQRService;

  constructor(
    @Inject('VIETQR_PACKAGE') private client: ClientGrpc,
  ) {}

  onModuleInit() {
    this.vietqrService = this.client.getService<VietQRService>('VietQRService');
  }

  async encodeQR(data: EncodeRequest): Promise<EncodeResponse> {
    return new Promise((resolve, reject) => {
      this.vietqrService.encode(data).subscribe({
        next: (response) => resolve(response),
        error: (error) => reject(error),
      });
    });
  }

  async decodeQR(qrCode: string): Promise<DecodeResponse> {
    return new Promise((resolve, reject) => {
      this.vietqrService.decode({ qr_code: qrCode }).subscribe({
        next: (response) => resolve(response),
        error: (error) => reject(error),
      });
    });
  }

  async getBankList(): Promise<any> {
    return new Promise((resolve, reject) => {
      this.vietqrService.getBankList({}).subscribe({
        next: (response) => resolve(response),
        error: (error) => reject(error),
      });
    });
  }
}
```

### 5. Sử dụng trong controller

**vietqr.controller.ts:**
```typescript
import { Controller, Post, Body, Get } from '@nestjs/common';
import { VietQRClientService } from './vietqr.service';

@Controller('vietqr')
export class VietQRController {
  constructor(private readonly vietqrService: VietQRClientService) {}

  @Post('encode')
  async encode(@Body() body: any) {
    return await this.vietqrService.encodeQR({
      bank_code: body.bankCode,
      bank_no: body.bankNo,
      amount: body.amount || 0,
      message: body.message || '',
    });
  }

  @Post('decode')
  async decode(@Body() body: { qrCode: string }) {
    return await this.vietqrService.decodeQR(body.qrCode);
  }

  @Get('banks')
  async getBanks() {
    return await this.vietqrService.getBankList();
  }
}
```

## Docker Support

### Build và chạy với Docker

```bash
# Build image
docker build -t vietqr-grpc .

# Run container
docker run -p 9090:9090 vietqr-grpc
```

### Sử dụng docker-compose

```bash
docker-compose up -d
```

## Cấu trúc Project

```
vietqr/
├── api/                    # API documentation
├── proto/                  # Protocol Buffer definitions
│   └── vietqr.proto       # gRPC service definition
├── cmd/
│   └── grpc-server/       # gRPC server implementation
│       └── main.go
├── banks.go               # Bank constants
├── banks_helper.go        # Bank helper functions
├── encode.go              # Encode functions
├── decode.go              # Decode functions
├── transferinfo.go        # Transfer info struct
├── generate.sh            # Proto generation script (Linux/Mac)
├── generate.bat           # Proto generation script (Windows)
├── Makefile              # Make commands
└── go.mod                # Go dependencies
```

## Các lệnh Makefile

```bash
make proto       # Generate protobuf code
make run-server  # Run gRPC server
make build       # Build gRPC server binary
make deps        # Install dependencies
make clean       # Clean generated files
make help        # Show help
```

## Biến môi trường

| Biến | Mô tả | Giá trị mặc định |
|------|-------|------------------|
| GRPC_PORT | Port của gRPC server | 9090 |

## Danh sách ngân hàng hỗ trợ

Xem file `banks.go` để biết danh sách đầy đủ các ngân hàng được hỗ trợ, bao gồm:
- TECHCOMBANK
- VIETCOMBANK
- BIDV
- VIETINBANK
- ACB
- MBBANK
- ...và nhiều ngân hàng khác

## Liên hệ & Hỗ trợ

- GitHub: https://github.com/sunary/vietqr
- Issues: https://github.com/sunary/vietqr/issues

## License

MIT License - xem file LICENSE để biết thêm chi tiết.

