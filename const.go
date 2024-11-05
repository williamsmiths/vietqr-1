package vietqr

const (
	// header
	_VERSION             = "00"
	_INIT_METHOD         = "01"
	_VNPAYQR             = "26"
	_VIETQR              = "38"
	_CATEGORY            = "52"
	_CURRENCY            = "53"
	_AMOUNT              = "54"
	_TIP_AND_FEE_TYPE    = "55"
	_TIP_AND_FEE_AMOUNT  = "56"
	_TIP_AND_FEE_PERCENT = "57"
	_NATION              = "58"
	_MERCHANT_NAME       = "59"
	_CITY                = "60"
	_ZIP_CODE            = "61"
	_ADDITIONAL_DATA     = "62"
	_CRC                 = "63"

	// provider header
	_PROVIDER_GUID    = "00"
	_PROVIDER_DATA    = "01"
	_PROVIDER_SERVICE = "02"

	// bank header
	_BANK_BIN    = "00"
	_BANK_NUMBER = "01"

	_PROVIDER_VNPAY_GUID  = "A000000775"
	_PROVIDER_VIETQR_GUID = "A000000727"

	_BY_ACCOUNT_NUMBER = "QRIBFTTA" // Dịch vụ chuyển nhanh NAPAS247 đến Tài khoản
	_BY_CARD_NUMBER    = "QRIBFTTC" // Dịch vụ chuyển nhanh NAPAS247 đến Thẻ

	// additional header
	_BILL_NUMBER                      = "01" // Số hóa đơn
	_MOBILE_NUMBER                    = "02" // Số ĐT
	_STORE_LABEL                      = "03" // Mã cửa hàng
	_LOYALTY_NUMBER                   = "04" // Mã khách hàng thân thiết
	_REFERENCE_LABEL                  = "05" // Mã tham chiếu
	_CUSTOMER_LABEL                   = "06" // Mã khách hàng
	_TERMINAL_LABEL                   = "07" // Mã số điểm bán
	_PURPOSE_OF_TRANSACTION           = "08" // Mục đích giao dịch
	_ADDITIONAL_CONSUMER_DATA_REQUEST = "09" // Yêu cầu dữ liệu KH bổ sung
)
