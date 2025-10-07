@echo off
REM Script để generate protobuf code trên Windows

echo Generating protobuf code for VietQR gRPC...

REM Kiểm tra protoc đã được cài đặt chưa
where protoc >nul 2>nul
if %errorlevel% neq 0 (
    echo protoc chua duoc cai dat. Vui long cai dat Protocol Buffers compiler.
    echo Tai tai: https://grpc.io/docs/protoc-installation/
    exit /b 1
)

REM Kiểm tra Go protobuf plugins
where protoc-gen-go >nul 2>nul
if %errorlevel% neq 0 (
    echo Installing protoc-gen-go...
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
)

where protoc-gen-go-grpc >nul 2>nul
if %errorlevel% neq 0 (
    echo Installing protoc-gen-go-grpc...
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
)

REM Generate proto files
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/vietqr.proto

if %errorlevel% equ 0 (
    echo √ Protobuf code generated successfully!
) else (
    echo × Failed to generate protobuf code
    exit /b 1
)

