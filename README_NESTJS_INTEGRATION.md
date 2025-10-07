# Tích hợp VietQR gRPC với NestJS

Hướng dẫn chi tiết cách tích hợp VietQR gRPC service vào NestJS application.

## 1. Cài đặt Dependencies trong NestJS

```bash
npm install --save @grpc/grpc-js @grpc/proto-loader
npm install --save @nestjs/microservices
```

## 2. Copy Proto File

Copy file `proto/vietqr.proto` từ Go project vào NestJS project:

```bash
# Tạo thư mục proto trong NestJS project
mkdir -p src/proto

# Copy file proto
cp /path/to/vietqr-go/proto/vietqr.proto src/proto/
```

## 3. Tạo Interface cho TypeScript

Tạo file `src/vietqr/interfaces/vietqr.interface.ts`:

```typescript
export interface EncodeRequest {
  bank_code: string;
  bank_no: string;
  amount: number;
  message: string;
}

export interface EncodeResponse {
  qr_code: string;
  success: boolean;
  error: string;
}

export interface DecodeRequest {
  qr_code: string;
}

export interface DecodeResponse {
  bank_code: string;
  bank_no: string;
  amount: number;
  message: string;
  success: boolean;
  error: string;
}

export interface BankInfo {
  code: string;
  name: string;
  bin: string;
}

export interface GetBankListRequest {}

export interface GetBankListResponse {
  banks: BankInfo[];
  success: boolean;
  error: string;
}

export interface VietQRGrpcService {
  encode(data: EncodeRequest): Promise<EncodeResponse>;
  decode(data: DecodeRequest): Promise<DecodeResponse>;
  getBankList(data: GetBankListRequest): Promise<GetBankListResponse>;
}
```

## 4. Tạo VietQR Module

Tạo file `src/vietqr/vietqr.module.ts`:

```typescript
import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { VietQRService } from './vietqr.service';
import { VietQRController } from './vietqr.controller';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'VIETQR_PACKAGE',
        transport: Transport.GRPC,
        options: {
          package: 'vietqr',
          protoPath: join(__dirname, '../proto/vietqr.proto'),
          url: process.env.VIETQR_GRPC_URL || 'localhost:50051',
          loader: {
            keepCase: false,
            longs: String,
            enums: String,
            defaults: true,
            oneofs: true,
          },
        },
      },
    ]),
  ],
  controllers: [VietQRController],
  providers: [VietQRService],
  exports: [VietQRService],
})
export class VietQRModule {}
```

## 5. Tạo VietQR Service

Tạo file `src/vietqr/vietqr.service.ts`:

```typescript
import {
  Injectable,
  Inject,
  OnModuleInit,
  Logger,
  InternalServerErrorException,
} from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { lastValueFrom } from 'rxjs';
import {
  VietQRGrpcService,
  EncodeRequest,
  EncodeResponse,
  DecodeRequest,
  DecodeResponse,
  GetBankListResponse,
} from './interfaces/vietqr.interface';

@Injectable()
export class VietQRService implements OnModuleInit {
  private readonly logger = new Logger(VietQRService.name);
  private vietqrService: VietQRGrpcService;

  constructor(
    @Inject('VIETQR_PACKAGE') private readonly client: ClientGrpc,
  ) {}

  onModuleInit() {
    this.vietqrService = this.client.getService<VietQRGrpcService>('VietQRService');
    this.logger.log('VietQR gRPC client initialized');
  }

  async encodeQR(data: EncodeRequest): Promise<EncodeResponse> {
    try {
      this.logger.log(`Encoding QR for bank: ${data.bank_code}, account: ${data.bank_no}`);
      const response = await lastValueFrom(
        this.vietqrService.encode(data),
      );
      
      if (!response.success) {
        this.logger.error(`Encode failed: ${response.error}`);
      }
      
      return response;
    } catch (error) {
      this.logger.error('Failed to encode QR', error);
      throw new InternalServerErrorException('Không thể tạo mã QR');
    }
  }

  async decodeQR(qrCode: string): Promise<DecodeResponse> {
    try {
      this.logger.log('Decoding QR code');
      const response = await lastValueFrom(
        this.vietqrService.decode({ qr_code: qrCode }),
      );
      
      if (!response.success) {
        this.logger.error(`Decode failed: ${response.error}`);
      }
      
      return response;
    } catch (error) {
      this.logger.error('Failed to decode QR', error);
      throw new InternalServerErrorException('Không thể giải mã QR');
    }
  }

  async getBankList(): Promise<GetBankListResponse> {
    try {
      this.logger.log('Getting bank list');
      const response = await lastValueFrom(
        this.vietqrService.getBankList({}),
      );
      return response;
    } catch (error) {
      this.logger.error('Failed to get bank list', error);
      throw new InternalServerErrorException('Không thể lấy danh sách ngân hàng');
    }
  }
}
```

## 6. Tạo VietQR Controller

Tạo file `src/vietqr/vietqr.controller.ts`:

```typescript
import { Controller, Post, Get, Body, HttpCode, HttpStatus } from '@nestjs/common';
import { VietQRService } from './vietqr.service';
import {
  ApiTags,
  ApiOperation,
  ApiResponse,
  ApiBody,
} from '@nestjs/swagger';

class EncodeQRDto {
  bankCode: string;
  bankNo: string;
  amount?: number;
  message?: string;
}

class DecodeQRDto {
  qrCode: string;
}

@ApiTags('VietQR')
@Controller('vietqr')
export class VietQRController {
  constructor(private readonly vietqrService: VietQRService) {}

  @Post('encode')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Tạo mã QR thanh toán VietQR' })
  @ApiBody({ type: EncodeQRDto })
  @ApiResponse({ status: 200, description: 'Tạo mã QR thành công' })
  @ApiResponse({ status: 400, description: 'Dữ liệu đầu vào không hợp lệ' })
  @ApiResponse({ status: 500, description: 'Lỗi server' })
  async encode(@Body() body: EncodeQRDto) {
    const result = await this.vietqrService.encodeQR({
      bank_code: body.bankCode,
      bank_no: body.bankNo,
      amount: body.amount || 0,
      message: body.message || '',
    });

    return {
      success: result.success,
      data: result.success ? { qrCode: result.qr_code } : null,
      error: result.error || null,
    };
  }

  @Post('decode')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Giải mã QR code VietQR' })
  @ApiBody({ type: DecodeQRDto })
  @ApiResponse({ status: 200, description: 'Giải mã thành công' })
  @ApiResponse({ status: 400, description: 'Mã QR không hợp lệ' })
  @ApiResponse({ status: 500, description: 'Lỗi server' })
  async decode(@Body() body: DecodeQRDto) {
    const result = await this.vietqrService.decodeQR(body.qrCode);

    return {
      success: result.success,
      data: result.success
        ? {
            bankCode: result.bank_code,
            bankNo: result.bank_no,
            amount: result.amount,
            message: result.message,
          }
        : null,
      error: result.error || null,
    };
  }

  @Get('banks')
  @ApiOperation({ summary: 'Lấy danh sách ngân hàng hỗ trợ' })
  @ApiResponse({ status: 200, description: 'Lấy danh sách thành công' })
  @ApiResponse({ status: 500, description: 'Lỗi server' })
  async getBanks() {
    const result = await this.vietqrService.getBankList();

    return {
      success: result.success,
      data: result.success ? result.banks : [],
      error: result.error || null,
    };
  }
}
```

## 7. Import Module vào App

Cập nhật `src/app.module.ts`:

```typescript
import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { VietQRModule } from './vietqr/vietqr.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    VietQRModule,
  ],
})
export class AppModule {}
```

## 8. Cấu hình Environment

Tạo file `.env`:

```env
# VietQR gRPC Configuration
VIETQR_GRPC_URL=localhost:50051

# Nếu deploy production
# VIETQR_GRPC_URL=vietqr-grpc-server:50051
```

## 9. Test API

### Encode QR Code

```bash
curl -X POST http://localhost:3000/vietqr/encode \
  -H "Content-Type: application/json" \
  -d '{
    "bankCode": "TECHCOMBANK",
    "bankNo": "9796868",
    "amount": 50000,
    "message": "Thanh toan don hang"
  }'
```

### Decode QR Code

```bash
curl -X POST http://localhost:3000/vietqr/decode \
  -H "Content-Type: application/json" \
  -d '{
    "qrCode": "00020101021138510010A00000072701210006970407010797968680208QRIBFTTA53037045802VN62240820gen by sunary/vietqr6304BE74"
  }'
```

### Get Bank List

```bash
curl http://localhost:3000/vietqr/banks
```

## 10. Docker Compose cho Full Stack

Tạo `docker-compose.yml` trong NestJS project:

```yaml
version: '3.8'

services:
  # Go gRPC Server
  vietqr-grpc:
    image: vietqr-grpc:latest
    container_name: vietqr-grpc-server
    ports:
      - "50051:50051"
    networks:
      - app-network
    restart: unless-stopped

  # NestJS API
  nestjs-api:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: nestjs-api
    ports:
      - "3000:3000"
    environment:
      - VIETQR_GRPC_URL=vietqr-grpc:50051
    depends_on:
      - vietqr-grpc
    networks:
      - app-network
    restart: unless-stopped

networks:
  app-network:
    driver: bridge
```

## 11. Health Check

Tạo health check endpoint `src/health/health.controller.ts`:

```typescript
import { Controller, Get } from '@nestjs/common';
import { VietQRService } from '../vietqr/vietqr.service';

@Controller('health')
export class HealthController {
  constructor(private readonly vietqrService: VietQRService) {}

  @Get()
  async check() {
    try {
      // Kiểm tra kết nối với gRPC server
      await this.vietqrService.getBankList();
      return {
        status: 'ok',
        vietqr_grpc: 'connected',
        timestamp: new Date().toISOString(),
      };
    } catch (error) {
      return {
        status: 'error',
        vietqr_grpc: 'disconnected',
        error: error.message,
        timestamp: new Date().toISOString(),
      };
    }
  }
}
```

## 12. Error Handling

Tạo custom exception filter `src/filters/grpc-exception.filter.ts`:

```typescript
import {
  ExceptionFilter,
  Catch,
  ArgumentsHost,
  HttpStatus,
  Logger,
} from '@nestjs/common';
import { RpcException } from '@nestjs/microservices';

@Catch(RpcException)
export class GrpcExceptionFilter implements ExceptionFilter {
  private readonly logger = new Logger(GrpcExceptionFilter.name);

  catch(exception: RpcException, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse();
    const error = exception.getError();

    this.logger.error('gRPC Error:', error);

    response.status(HttpStatus.INTERNAL_SERVER_ERROR).json({
      success: false,
      error: 'Lỗi kết nối với VietQR service',
      details: error,
      timestamp: new Date().toISOString(),
    });
  }
}
```

Apply filter trong `main.ts`:

```typescript
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { GrpcExceptionFilter } from './filters/grpc-exception.filter';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.useGlobalFilters(new GrpcExceptionFilter());
  await app.listen(3000);
}
bootstrap();
```

## Troubleshooting

### Lỗi kết nối gRPC

1. Kiểm tra Go gRPC server đang chạy:
   ```bash
   grpcurl -plaintext localhost:50051 list
   ```

2. Kiểm tra URL trong `.env` đúng chưa

3. Kiểm tra network trong Docker Compose

### Lỗi proto file

1. Đảm bảo file proto đã được copy đúng
2. Kiểm tra đường dẫn trong `protoPath`

## Best Practices

1. **Retry Logic**: Implement retry cho gRPC calls
2. **Timeout**: Set timeout phù hợp cho gRPC connections
3. **Circuit Breaker**: Sử dụng circuit breaker pattern
4. **Caching**: Cache bank list để giảm calls
5. **Monitoring**: Log và monitor gRPC calls
6. **Health Checks**: Implement health checks cho cả services

## Kết luận

Sau khi hoàn thành các bước trên, bạn đã có một hệ thống hoàn chỉnh:
- Go gRPC server xử lý VietQR logic
- NestJS API server làm gateway
- Communication qua gRPC protocol
- Dễ dàng scale và maintain

