package main

import (
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
