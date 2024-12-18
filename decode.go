package vietqr

import (
	"errors"
	"strconv"
)

func Decode(s string) (*TransferInfo, error) {
	if !validCrcContent(s) {
		return nil, errors.New("invalid CRC")
	}

	ti := TransferInfo{}
	ti.parseRootContent(s[:len(s)-paddingCrc])
	if ti.merchantID != "" {
		return &ti, nil
	}

	// handle VIETQR guid, require BankCode and BankNo
	ti.BankCode = revBankBin[ti.bankBin]
	if ti.BankCode == "" || ti.BankNo == "" {
		return nil, errors.New("invalid VIETQR code")
	}

	return &ti, nil
}

func (t *TransferInfo) parseRootContent(s string) {
	id, value, nextValue := slideContent(s)
	switch id {
	case _VERSION:
	case _INIT_METHOD:
	case _VNPAYQR, _VIETQR:
		t.parseProviderInfo(value)
	case _CATEGORY:
	case _CURRENCY:
	case _AMOUNT:
		t.Amount, _ = strconv.ParseInt(value, 10, 64)
	case _TIP_AND_FEE_TYPE:
	case _TIP_AND_FEE_AMOUNT:
	case _TIP_AND_FEE_PERCENT:
	case _NATION:
	case _MERCHANT_NAME:
	case _CITY:
	case _ZIP_CODE:
	case _ADDITIONAL_DATA:
		t.parseAdditionalData(value)
	}

	if len(nextValue) > 4 {
		t.parseRootContent(nextValue)
	}
}

func (t *TransferInfo) parseProviderInfo(s string) {
	id, value, nextValue := slideContent(s)
	switch id {
	case _PROVIDER_GUID:
		t.guid = value
	case _PROVIDER_DATA:
		if t.guid == _PROVIDER_VNPAY_GUID {
			t.merchantID = value
		} else if t.guid == _PROVIDER_VIETQR_GUID {
			t.parseVietQRConsumer(value)
		}
	case _PROVIDER_SERVICE:
	}

	if len(nextValue) > 4 {
		t.parseProviderInfo(nextValue)
	}
}

func (t *TransferInfo) parseVietQRConsumer(s string) {
	id, value, nextValue := slideContent(s)
	switch id {
	case _BANK_BIN:
		t.bankBin = value
	case _BANK_NUMBER:
		t.BankNo = value
	}

	if len(nextValue) > 4 {
		t.parseVietQRConsumer(nextValue)
	}
}

func (t *TransferInfo) parseAdditionalData(s string) {
	id, value, nextValue := slideContent(s)
	switch id {
	case _PURPOSE_OF_TRANSACTION:
		t.Message = value
	case _BILL_NUMBER:
	case _MOBILE_NUMBER:
	case _REFERENCE_LABEL:
	case _STORE_LABEL:
	case _TERMINAL_LABEL:
	}

	if len(nextValue) > 4 {
		t.parseAdditionalData(nextValue)
	}
}
