package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sunary/vietqr"
	pb "github.com/sunary/vietqr/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server implements VietQRServiceServer
type server struct {
	pb.UnimplementedVietQRServiceServer
}

// Encode tạo mã QR từ thông tin chuyển khoản
func (s *server) Encode(ctx context.Context, req *pb.EncodeRequest) (*pb.EncodeResponse, error) {
	log.Printf("Encode request: BankCode=%s, BankNo=%s, Amount=%d, Message=%s",
		req.BankCode, req.BankNo, req.Amount, req.Message)

	// Validate input
	if req.BankCode == "" {
		return &pb.EncodeResponse{
			Success: false,
			Error:   "bank_code không được để trống",
		}, nil
	}
	if req.BankNo == "" {
		return &pb.EncodeResponse{
			Success: false,
			Error:   "bank_no không được để trống",
		}, nil
	}

	// Encode QR code
	qrCode := vietqr.Encode(vietqr.TransferInfo{
		BankCode: req.BankCode,
		BankNo:   req.BankNo,
		Amount:   req.Amount,
		Message:  req.Message,
	})

	return &pb.EncodeResponse{
		QrCode:  qrCode,
		Success: true,
	}, nil
}

// Decode giải mã thông tin từ mã QR
func (s *server) Decode(ctx context.Context, req *pb.DecodeRequest) (*pb.DecodeResponse, error) {
	log.Printf("Decode request: QRCode=%s", req.QrCode)

	// Validate input
	if req.QrCode == "" {
		return &pb.DecodeResponse{
			Success: false,
			Error:   "qr_code không được để trống",
		}, nil
	}

	// Decode QR code
	info, err := vietqr.Decode(req.QrCode)
	if err != nil {
		log.Printf("Decode error: %v", err)
		return &pb.DecodeResponse{
			Success: false,
			Error:   fmt.Sprintf("Lỗi giải mã QR: %v", err),
		}, nil
	}

	return &pb.DecodeResponse{
		BankCode: info.BankCode,
		BankNo:   info.BankNo,
		Amount:   info.Amount,
		Message:  info.Message,
		Success:  true,
	}, nil
}

// GetBankList lấy danh sách các ngân hàng hỗ trợ
func (s *server) GetBankList(ctx context.Context, req *pb.GetBankListRequest) (*pb.GetBankListResponse, error) {
	log.Println("GetBankList request")

	banks := vietqr.GetBankList()
	pbBanks := make([]*pb.BankInfo, len(banks))

	for i, bank := range banks {
		pbBanks[i] = &pb.BankInfo{
			Code: bank.Code,
			Name: bank.Name,
			Bin:  bank.Bin,
		}
	}

	return &pb.GetBankListResponse{
		Banks:   pbBanks,
		Success: true,
	}, nil
}

// HealthCheck kiểm tra trạng thái server
func (s *server) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Println("HealthCheck request")

	return &pb.HealthCheckResponse{
		Healthy:   true,
		Status:    "ok",
		Message:   "VietQR gRPC Server đang hoạt động bình thường",
		Timestamp: time.Now().Unix(),
		Version:   "1.0.0",
	}, nil
}

func main() {
	// Lấy port từ biến môi trường, mặc định là 9090
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "9090"
	}

	// Tạo listener
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Tạo gRPC server
	s := grpc.NewServer()
	pb.RegisterVietQRServiceServer(s, &server{})

	// Enable reflection để có thể sử dụng grpcurl
	reflection.Register(s)

	log.Printf("gRPC server listening on port %s", port)
	log.Println("Server ready to accept requests...")

	// Graceful shutdown
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down gRPC server...")
	s.GracefulStop()
	log.Println("Server stopped")
}
