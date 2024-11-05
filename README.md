# vietqr

This Go-based library enables encoding and decoding for VietQR, a standardized QR code solution for payment and transfer services across Vietnam's NAPAS network. Once encoded, the QR code can be further customized using any QR code generator that converts text to QR format.

VietQR serves as a unified brand identity for QR-based payments and transfers, seamlessly processed through the NAPAS network, member banks, payment intermediaries, and domestic and international partners. It complies with the EMV Co. QR payment standards and adheres to foundational QR code standards established by the State Bank of Vietnam.

## Sample

```go
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
