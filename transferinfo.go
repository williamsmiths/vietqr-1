package vietqr

type TransferInfo struct {
	guid       string
	merchantID string // VNPAY-merchantId, disable for VIETQR generator
	BankCode   string
	bankBin    string
	BankNo     string
	Amount     int64
	Message    string
}
