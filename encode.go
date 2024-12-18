package vietqr

import (
	"log"
	"strconv"
)

func Encode(ti TransferInfo) string {
	content := qrContent(ti)
	return content + hashCrc(content)
}

func qrContent(ti TransferInfo) string {
	version := genFieldData(_VERSION, "01")
	initMethod := genFieldData(_INIT_METHOD, "11")
	amount := ""

	if ti.Amount > 0 {
		initMethod = genFieldData(_INIT_METHOD, "12")
		amount = strconv.Itoa(int(ti.Amount))
	}

	var guid, providerDataContent string
	if ti.merchantID == "" { // default
		bin := bankBin[ti.BankCode]
		if bin == "" {
			log.Fatalf("unknown bank code: %s", ti.BankCode)
		}

		bankBin := genFieldData(_BANK_BIN, bin)
		bankNumber := genFieldData(_BANK_NUMBER, ti.BankNo)
		providerDataContent = bankBin + bankNumber
		guid = genFieldData(_PROVIDER_GUID, _PROVIDER_VIETQR_GUID)
	} else {
		providerDataContent = ti.merchantID
		guid = genFieldData(_PROVIDER_GUID, _PROVIDER_VNPAY_GUID)
	}

	provider := genFieldData(_PROVIDER_DATA, providerDataContent)
	service := genFieldData(_PROVIDER_SERVICE, _BY_ACCOUNT_NUMBER)
	providerData := genFieldData(_VIETQR, guid+provider+service)

	category := genFieldData(_CATEGORY, "")
	currency := genFieldData(_CURRENCY, "704")
	amountStr := genFieldData(_AMOUNT, amount)

	tipAndFeeType := genFieldData(_TIP_AND_FEE_TYPE, "")
	tipAndFeeAmount := genFieldData(_TIP_AND_FEE_AMOUNT, "")
	tipAndFeePercent := genFieldData(_TIP_AND_FEE_PERCENT, "")
	nation := genFieldData(_NATION, "VN")
	merchantName := genFieldData(_MERCHANT_NAME, "")
	city := genFieldData(_CITY, "")
	zipCode := genFieldData(_ZIP_CODE, "")

	purpose := genFieldData(_PURPOSE_OF_TRANSACTION, ti.Message)
	//joinString(buildNumber, mobileNumber, storeLabel, loyaltyNumber, reference, customerLabel, terminal, purpose, dataRequest)
	additionalData := genFieldData(_ADDITIONAL_DATA, purpose)

	EVMCoContent := ""
	unreservedContent := ""

	return joinString(version, initMethod, providerData, category, currency, amountStr, tipAndFeeType, tipAndFeeAmount, tipAndFeePercent,
		nation, merchantName, city, zipCode, additionalData, EVMCoContent, unreservedContent, _CRC, "04")
}
