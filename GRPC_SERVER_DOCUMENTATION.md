# VietQR gRPC Server Documentation

## Tổng quan

VietQR gRPC Server là một microservice cung cấp API để tạo và giải mã mã QR chuyển khoản ngân hàng theo chuẩn VietQR của Việt Nam. Server được xây dựng bằng Go và sử dụng gRPC protocol để giao tiếp với các ứng dụng client như NestJS.

## Kiến trúc hệ thống

### Cấu trúc thư mục
```
vietqr-1/
├── cmd/grpc-server/
│   └── main.go              # Entry point của gRPC server
├── proto/
│   ├── vietqr.proto         # Protocol buffer definitions
│   ├── vietqr.pb.go         # Generated Go code từ proto
│   └── vietqr_grpc.pb.go    # Generated gRPC service code
├── banks_helper.go          # Danh sách ngân hàng hỗ trợ
├── const.go                 # Constants cho QR encoding
├── encode.go                # Logic encode QR code
├── decode.go                # Logic decode QR code
├── transferinfo.go          # Struct thông tin chuyển khoản
├── utils.go                 # Utility functions
├── crc16.go                 # CRC16 checksum implementation
└── test/
    └── test_client.go       # Test client
```

### Dependencies chính
- `google.golang.org/grpc` - gRPC framework
- `github.com/sunary/vietqr` - VietQR library
- Protocol Buffers cho serialization
- CRC16 checksum cho validation

## API Services

### 1. Encode Service
**Mục đích**: Tạo mã QR từ thông tin chuyển khoản

**Request**:
```protobuf
message EncodeRequest {
    string bank_code = 1;  // Mã ngân hàng (bắt buộc)
    string bank_no = 2;    // Số tài khoản (bắt buộc)
    int64 amount = 3;      // Số tiền (tùy chọn, 0 nếu không có)
    string message = 4;    // Nội dung chuyển khoản (tùy chọn)
}
```

**Response**:
```protobuf
message EncodeResponse {
    bool success = 1;      // Trạng thái thành công
    string qr_code = 2;    // Mã QR được tạo
    string error = 3;      // Thông báo lỗi nếu có
}
```

**Validation**:
- `bank_code` không được để trống
- `bank_no` không được để trống
- `amount` và `message` là optional

**Ví dụ Request**:
```json
{
    "bank_code": "TECHCOMBANK",
    "bank_no": "9796868",
    "amount": 50000,
    "message": "Thanh toan don hang"
}
```

**Ví dụ Response**:
```json
{
    "success": true,
    "qr_code": "00020101021138510010A00000072701210006970407010797968680208QRIBFTTA53037045802VN62240820gen by sunary/vietqr6304BE74",
    "error": ""
}
```

### 2. Decode Service
**Mục đích**: Giải mã thông tin từ mã QR

**Request**:
```protobuf
message DecodeRequest {
    string qr_code = 1;    // Mã QR cần giải mã
}
```

**Response**:
```protobuf
message DecodeResponse {
    bool success = 1;       // Trạng thái thành công
    string bank_code = 2;  // Mã ngân hàng
    string bank_no = 3;    // Số tài khoản
    int64 amount = 4;      // Số tiền
    string message = 5;    // Nội dung
    string error = 6;      // Thông báo lỗi nếu có
}
```

**Validation**:
- `qr_code` không được để trống

**Ví dụ Request**:
```json
{
    "qr_code": "00020101021138510010A00000072701210006970407010797968680208QRIBFTTA53037045802VN62240820gen by sunary/vietqr6304BE74"
}
```

**Ví dụ Response**:
```json
{
    "success": true,
    "bank_code": "TECHCOMBANK",
    "bank_no": "9796868",
    "amount": 50000,
    "message": "Thanh toan don hang",
    "error": ""
}
```

### 3. GetBankList Service
**Mục đích**: Lấy danh sách các ngân hàng được hỗ trợ

**Request**:
```protobuf
message GetBankListRequest {
    // Không có parameters
}
```

**Response**:
```protobuf
message GetBankListResponse {
    bool success = 1;
    repeated BankInfo banks = 2;
    string error = 3;
}

message BankInfo {
    string code = 1;       // Mã ngân hàng
    string name = 2;       // Tên ngân hàng
    string bin = 3;        // BIN code
}
```

**Ví dụ Response**:
```json
{
    "success": true,
    "banks": [
        {
            "code": "TECHCOMBANK",
            "name": "Ngân hàng TMCP Kỹ Thương Việt Nam",
            "bin": "970407"
        },
        {
            "code": "VIETCOMBANK",
            "name": "Ngân hàng TMCP Ngoại thương Việt Nam",
            "bin": "970436"
        }
    ],
    "error": ""
}
```

### 4. HealthCheck Service
**Mục đích**: Kiểm tra trạng thái hoạt động của server

**Request**:
```protobuf
message HealthCheckRequest {
    // Không có parameters
}
```

**Response**:
```protobuf
message HealthCheckResponse {
    bool healthy = 1;           // Server có khỏe mạnh không
    string status = 2;          // Trạng thái: "ok", "error"
    string message = 3;         // Thông báo chi tiết
    int64 timestamp = 4;        // Unix timestamp
    string version = 5;         // Phiên bản server
}
```

**Ví dụ Response**:
```json
{
    "healthy": true,
    "status": "ok",
    "message": "VietQR gRPC Server đang hoạt động bình thường",
    "timestamp": 1759902306,
    "version": "1.0.0"
}
```

## Cấu hình Server

### Environment Variables
- `GRPC_PORT`: Port để server lắng nghe (mặc định: 9090)

### Network Configuration
- **Protocol**: TCP
- **Address**: 0.0.0.0 (listen trên tất cả interfaces)
- **Port**: 9090 (có thể thay đổi qua environment variable)

### Features
- **gRPC Reflection**: Được enable để hỗ trợ debugging với grpcurl
- **Graceful Shutdown**: Hỗ trợ tắt server an toàn với SIGINT/SIGTERM
- **Logging**: Chi tiết log cho mỗi request và error
- **CRC16 Validation**: Kiểm tra tính toàn vẹn của QR code

## Cách chạy Server

### 1. Chạy trực tiếp
```bash
cd cmd/grpc-server
go run main.go
```

### 2. Build và chạy
```bash
go build -o vietqr-server main.go
./vietqr-server
```

### 3. Với custom port
```bash
GRPC_PORT=8080 go run main.go
```

### 4. Sử dụng Docker
```bash
docker-compose up vietqr-server
```

## Testing với grpcurl

### 1. Test Encode
```bash
grpcurl -plaintext -d '{
  "bank_code": "TECHCOMBANK",
  "bank_no": "9796868",
  "amount": 50000,
  "message": "Test thanh toan"
}' localhost:9090 vietqr.VietQRService/Encode
```

### 2. Test Decode
```bash
grpcurl -plaintext -d '{
  "qr_code": "00020101021138510010A00000072701210006970407010797968680208QRIBFTTA53037045802VN62240820gen by sunary/vietqr6304BE74"
}' localhost:9090 vietqr.VietQRService/Decode
```

### 3. Test GetBankList
```bash
grpcurl -plaintext localhost:9090 vietqr.VietQRService/GetBankList
```

### 4. Test HealthCheck
```bash
grpcurl -plaintext localhost:9090 vietqr.VietQRService/HealthCheck
```

## QR Code Format

### Cấu trúc QR Code VietQR
QR code được tạo theo chuẩn VietQR với cấu trúc:

```
00 - Version (01)
01 - Init Method (11: không có số tiền, 12: có số tiền)
38 - VietQR Provider Data
  00 - Provider GUID (A000000727)
  01 - Provider Data (Bank BIN + Account Number)
  02 - Provider Service (QRIBFTTA)
52 - Category (trống)
53 - Currency (704 - VND)
54 - Amount (nếu có)
55 - Tip and Fee Type (trống)
56 - Tip and Fee Amount (trống)
57 - Tip and Fee Percent (trống)
58 - Nation (VN)
59 - Merchant Name (trống)
60 - City (trống)
61 - ZIP Code (trống)
62 - Additional Data
  08 - Purpose of Transaction (nội dung chuyển khoản)
63 - CRC (4 ký tự hex)
```

### Ví dụ QR Code
```
00020101021138510010A00000072701210006970407010797968680208QRIBFTTA53037045802VN62240820gen by sunary/vietqr6304BE74
```

**Giải thích**:
- `00020101021138510010A00000072701210006970407010797968680208QRIBFTTA53037045802VN62240820gen by sunary/vietqr6304BE74` - Nội dung QR
- `6304BE74` - CRC16 checksum

## Error Handling

### Validation Errors
- **bank_code empty**: "bank_code không được để trống"
- **bank_no empty**: "bank_no không được để trống"
- **qr_code empty**: "qr_code không được để trống"

### Decode Errors
- **Invalid QR format**: "Lỗi giải mã QR: [chi tiết lỗi]"
- **Invalid CRC**: "invalid CRC"
- **Invalid VIETQR code**: "invalid VIETQR code"

### Server Errors
- **Port binding failed**: "Failed to listen: [error]"
- **Server start failed**: "Failed to serve: [error]"

## Logging

Server ghi log chi tiết cho:
- Mỗi request đến (với parameters)
- Errors trong quá trình xử lý
- Server lifecycle events (start, stop)
- Graceful shutdown process

**Ví dụ log**:
```
2024/01/15 10:30:15 Encode request: BankCode=TECHCOMBANK, BankNo=9796868, Amount=50000, Message=Test thanh toan
2024/01/15 10:30:15 gRPC server listening on port 9090
2024/01/15 10:30:15 Server ready to accept requests...
```

## Performance & Scalability

### Concurrent Handling
- gRPC server tự động handle concurrent requests
- Mỗi request được xử lý trong goroutine riêng biệt

### Memory Usage
- Sử dụng protocol buffers cho serialization (hiệu quả hơn JSON)
- Graceful shutdown để tránh memory leaks

## Security Considerations

### Input Validation
- Validate tất cả input parameters
- Sanitize error messages để tránh information disclosure

### Network Security
- Có thể thêm TLS/SSL cho production
- Có thể thêm authentication/authorization middleware

## Monitoring & Debugging

### Health Checks
- Server có thể được monitor qua gRPC health check protocol
- Logs cung cấp thông tin chi tiết về server status

### Debugging Tools
- **grpcurl**: Test API endpoints
- **grpc reflection**: Inspect service definitions
- **Go debugger**: Debug source code

## Tích hợp với NestJS

### 1. Cài đặt Dependencies
```bash
npm install --save @grpc/grpc-js @grpc/proto-loader
npm install --save @nestjs/microservices
```

### 2. Cấu hình gRPC Client
```typescript
ClientsModule.register([
  {
    name: 'VIETQR_PACKAGE',
    transport: Transport.GRPC,
    options: {
      package: 'vietqr',
      protoPath: join(__dirname, '../proto/vietqr.proto'),
      url: process.env.VIETQR_GRPC_URL || 'localhost:9090',
    },
  },
])
```

### 3. Sử dụng trong Service
```typescript
async encodeQR(data: EncodeRequest): Promise<EncodeResponse> {
  const response = await lastValueFrom(
    this.vietqrService.encode(data),
  );
  return response;
}
```

## Deployment

### Docker Support
- Có Dockerfile và docker-compose.yml
- Support cho cả local và production deployment

### Environment Configuration
- Flexible port configuration
- Environment-based configuration

## Troubleshooting

### Common Issues
1. **Port already in use**: Thay đổi GRPC_PORT
2. **Connection refused**: Kiểm tra server có đang chạy không
3. **Invalid QR code**: Kiểm tra format của QR code input

### Debug Steps
1. Kiểm tra logs của server
2. Test với grpcurl để verify API
3. Kiểm tra network connectivity
4. Verify protocol buffer definitions

## Danh sách ngân hàng hỗ trợ

Server hỗ trợ hơn 70 ngân hàng tại Việt Nam, bao gồm:

### Ngân hàng lớn
- **TECHCOMBANK** - Ngân hàng TMCP Kỹ Thương Việt Nam
- **VIETCOMBANK** - Ngân hàng TMCP Ngoại thương Việt Nam
- **BIDV** - Ngân hàng TMCP Đầu tư và Phát triển Việt Nam
- **VIETINBANK** - Ngân hàng TMCP Công thương Việt Nam
- **AGRIBANK** - Ngân hàng Nông nghiệp và Phát triển Nông thôn Việt Nam

### Ngân hàng thương mại
- **ACB** - Ngân hàng TMCP Á Châu
- **SACOMBANK** - Ngân hàng TMCP Sài Gòn Thương Tín
- **VPBANK** - Ngân hàng TMCP Việt Nam Thịnh Vượng
- **MBBANK** - Ngân hàng TMCP Quân đội
- **HDBANK** - Ngân hàng TMCP Phát triển TP. Hồ Chí Minh

### Ngân hàng quốc tế
- **HSBC** - Ngân hàng TNHH MTV HSBC Việt Nam
- **SHINHAN_BANK** - Ngân hàng TNHH MTV Shinhan Việt Nam
- **STANDARD_CHARTERED_BANK** - Ngân hàng TNHH MTV Standard Chartered Bank Việt Nam

## Future Enhancements

### Potential Improvements
- Thêm authentication/authorization
- Implement rate limiting
- Add metrics và monitoring
- Support cho multiple protocols (HTTP/REST)
- Caching cho bank list
- Batch operations support
- Webhook notifications
- QR code image generation

## Kết luận

VietQR gRPC Server cung cấp một giải pháp hoàn chỉnh để tạo và giải mã mã QR chuyển khoản theo chuẩn VietQR. Với kiến trúc microservice, server có thể dễ dàng tích hợp với các ứng dụng khác như NestJS, Node.js, hoặc bất kỳ ngôn ngữ nào hỗ trợ gRPC.

Server được thiết kế để:
- Xử lý concurrent requests hiệu quả
- Cung cấp API đơn giản và dễ sử dụng
- Hỗ trợ đầy đủ chuẩn VietQR của Việt Nam
- Dễ dàng deploy và scale
- Tích hợp tốt với các hệ thống hiện có