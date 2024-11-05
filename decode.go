package vietqr

import (
	"errors"
	"strconv"
)

func Decode(s string) (*TransferInfo, error) {
	if !validCrcContent(s) {
		return nil, errors.New("invalid CRC")
	}

	t := TransferInfo{}
	t.parseRootContent(s[:len(s)-paddingCrc])
	t.BankCode = revBankBin[t.bankBin]
	if t.BankCode == "" || t.BankNo == "" {
		return nil, errors.New("invalid qr code")
	}

	return &t, nil
}

func (t *TransferInfo) parseRootContent(s string) {
	id, value, nextValue := slideContent(s)
	switch id {
	case _VERSION:
	case _INIT_METHOD:
	case _VNPAYQR:
	case _VIETQR:
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
		if t.guid == _PROVIDER_VIETQR_GUID {
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

func slideContent(s string) (id string, value string, nextContent string) {
	id = s[:2]
	length, _ := strconv.Atoi(s[2:4])
	value = s[4 : 4+length]
	nextContent = s[4+length:]
	return
}
