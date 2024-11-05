package vietqr

type TransferInfo struct {
	guid     string
	BankCode string
	bankBin  string
	BankNo   string
	Amount   int64
	Message  string
}
