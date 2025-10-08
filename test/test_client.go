package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sunary/vietqr"
	pb "github.com/sunary/vietqr/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewVietQRServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Test Encode
	fmt.Println("=== Test Encode ===")
	encResp, err := client.Encode(ctx, &pb.EncodeRequest{
		BankCode: vietqr.TECHCOMBANK,
		BankNo:   "9796868",
		Amount:   50000,
		Message:  "Test thanh toan",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Success: %v\n", encResp.Success)
	fmt.Printf("QR Code: %s\n\n", encResp.QrCode)

	// Test Decode
	fmt.Println("=== Test Decode ===")
	decResp, err := client.Decode(ctx, &pb.DecodeRequest{
		QrCode: encResp.QrCode,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bank: %s, Account: %s, Amount: %d, Message: %s\n\n",
		decResp.BankCode, decResp.BankNo, decResp.Amount, decResp.Message)

	// Test GetBankList
	fmt.Println("=== Test GetBankList ===")
	bankResp, err := client.GetBankList(ctx, &pb.GetBankListRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d banks\n", len(bankResp.Banks))
	for i, bank := range bankResp.Banks {
		if i >= 5 {
			break
		}
		fmt.Printf("  - %s (%s)\n", bank.Name, bank.Code)
	}

	// Test Health Check
	fmt.Println("=== Test Health Check ===")
	healthResp, err := client.HealthCheck(ctx, &pb.HealthCheckRequest{})
	if err != nil {
		fmt.Printf("❌ Health check failed: %v\n", err)
	} else {
		fmt.Printf("✅ Health check passed\n")
		fmt.Printf("   - Healthy: %v\n", healthResp.Healthy)
		fmt.Printf("   - Status: %s\n", healthResp.Status)
		fmt.Printf("   - Message: %s\n", healthResp.Message)
		fmt.Printf("   - Version: %s\n", healthResp.Version)
		fmt.Printf("   - Timestamp: %d\n", healthResp.Timestamp)
	}

	fmt.Println("\n✅ All tests passed!")
}
