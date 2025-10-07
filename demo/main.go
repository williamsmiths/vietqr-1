package main

import (
	"image/png"
	"os"

	"github.com/skip2/go-qrcode"
	"github.com/sunary/vietqr"
)

func main() {
	// Tạo chuỗi VietQR
	qrString := vietqr.Encode(vietqr.TransferInfo{
		BankCode: vietqr.TPBANK,
		BankNo:   "03370007501",
		Message:  "gen by sunary/vietqr",
	})

	println("VietQR String:", qrString)

	// Tạo mã QR code hình ảnh
	qr, err := qrcode.New(qrString, qrcode.Medium)
	if err != nil {
		println("Lỗi tạo QR:", err.Error())
		return
	}

	// Lưu file QR code
	file, err := os.Create("vietqr.png")
	if err != nil {
		println("Lỗi tạo file:", err.Error())
		return
	}
	defer file.Close()

	err = png.Encode(file, qr.Image(256))
	if err != nil {
		println("Lỗi lưu QR:", err.Error())
		return
	}

	println("Đã tạo file vietqr.png thành công!")

	// Test decode
	info, err := vietqr.Decode(qrString)
	if err != nil {
		println("err", err.Error())
		return
	}

	println("bank_code", info.BankCode)
	println("bank_no", info.BankNo)
	println("amount", info.Amount)
	println("message", info.Message)
}
