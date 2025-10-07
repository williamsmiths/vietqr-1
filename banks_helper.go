package vietqr

// BankInfo chứa thông tin một ngân hàng
type BankInfo struct {
	Code string
	Name string
	Bin  string
}

// GetBankList trả về danh sách tất cả các ngân hàng được hỗ trợ
func GetBankList() []BankInfo {
	banks := []BankInfo{}
	bankNames := map[string]string{
		ABBANK:                  "Ngân hàng TMCP An Bình",
		ACB:                     "Ngân hàng TMCP Á Châu",
		AGRIBANK:                "Ngân hàng Nông nghiệp và Phát triển Nông thôn Việt Nam",
		BAC_A_BANK:              "Ngân hàng TMCP Bắc Á",
		BAOVIET_BANK:            "Ngân hàng TMCP Bảo Việt",
		BANVIET:                 "Ngân hàng TMCP Bản Việt",
		BIDV:                    "Ngân hàng TMCP Đầu tư và Phát triển Việt Nam",
		CAKE:                    "CAKE by VPBank",
		CBBANK:                  "Ngân hàng Thương mại TNHH MTV Xây dựng Việt Nam",
		CIMB:                    "Ngân hàng TNHH MTV CIMB Việt Nam",
		COOP_BANK:               "Ngân hàng Hợp tác xã Việt Nam",
		DBS_BANK:                "DBS Bank Ltd - Chi nhánh TP. Hồ Chí Minh",
		DONG_A_BANK:             "Ngân hàng TMCP Đông Á",
		EXIMBANK:                "Ngân hàng TMCP Xuất Nhập khẩu Việt Nam",
		GPBANK:                  "Ngân hàng Thương mại TNHH MTV Dầu Khí Toàn Cầu",
		HDBANK:                  "Ngân hàng TMCP Phát triển TP. Hồ Chí Minh",
		HONGLEONG_BANK:          "Ngân hàng TNHH MTV Hong Leong Việt Nam",
		HSBC:                    "Ngân hàng TNHH MTV HSBC Việt Nam",
		IBK_HCM:                 "Ngân hàng Công nghiệp Hàn Quốc - Chi nhánh TP. Hồ Chí Minh",
		IBK_HN:                  "Ngân hàng Công nghiệp Hàn Quốc - Chi nhánh Hà Nội",
		INDOVINA_BANK:           "Ngân hàng TNHH Indovina",
		KASIKORN_BANK:           "Ngân hàng Đại chúng TNHH Kasikornbank",
		KEB_HANA_BANK_HCM:       "Ngân hàng KEB Hana - Chi nhánh TP. Hồ Chí Minh",
		KEB_HANA_BANK_HN:        "Ngân hàng KEB Hana - Chi nhánh Hà Nội",
		KIENLONG_BANK:           "Ngân hàng TMCP Kiên Long",
		KOOKMIN_BANK_HCM:        "Ngân hàng Kookmin - Chi nhánh TP. Hồ Chí Minh",
		KOOKMIN_BANK_HN:         "Ngân hàng Kookmin - Chi nhánh Hà Nội",
		LIENVIETPOST_BANK:       "Ngân hàng TMCP Bưu Điện Liên Việt",
		MBBANK:                  "Ngân hàng TMCP Quân đội",
		MB_SHINSEI:              "Ngân hàng TNHH MTV MB Shinsei",
		MIRAE_ASSET:             "Ngân hàng Mirae Asset (Việt Nam)",
		MSB:                     "Ngân hàng TMCP Hàng Hải",
		NAM_A_BANK:              "Ngân hàng TMCP Nam Á",
		NCB:                     "Ngân hàng TMCP Quốc Dân",
		NONGHYUP_BANK_HN:        "Ngân hàng Nonghyup - Chi nhánh Hà Nội",
		OCB:                     "Ngân hàng TMCP Phương Đông",
		OCEANBANK:               "Ngân hàng Thương mại TNHH MTV Đại Dương",
		PGBANK:                  "Ngân hàng TMCP Xăng dầu Petrolimex",
		PUBLIC_BANK:             "Ngân hàng TNHH MTV Public Việt Nam",
		PVCOM_BANK:              "Ngân hàng TMCP Đại Chúng Việt Nam",
		SACOMBANK:               "Ngân hàng TMCP Sài Gòn Thương Tín",
		SAIGONBANK:              "Ngân hàng TMCP Sài Gòn Công Thương",
		SCB:                     "Ngân hàng TMCP Sài Gòn",
		SEA_BANK:                "Ngân hàng TMCP Đông Nam Á",
		SHB:                     "Ngân hàng TMCP Sài Gòn - Hà Nội",
		SHINHAN_BANK:            "Ngân hàng TNHH MTV Shinhan Việt Nam",
		SINOPAC_BANK_HCM:        "Ngân hàng TNHH MTV SINOPAC - Chi nhánh TP. Hồ Chí Minh",
		STANDARD_CHARTERED_BANK: "Ngân hàng TNHH MTV Standard Chartered Bank Việt Nam",
		TECHCOMBANK:             "Ngân hàng TMCP Kỹ Thương Việt Nam",
		TIMO:                    "Timo by Ban Viet Bank",
		TNEX:                    "Ngân hàng TMCP Tiên Phong",
		TPBANK:                  "Ngân hàng TMCP Tiên Phong",
		UBANK:                   "Ubank by VPBank",
		UNITED_OVERSEAS_BANK:    "Ngân hàng United Overseas Bank - Chi nhánh TP. Hồ Chí Minh",
		VIB:                     "Ngân hàng TMCP Quốc tế Việt Nam",
		VIET_A_BANK:             "Ngân hàng TMCP Việt Á",
		VIET_BANK:               "Ngân hàng TMCP Việt Nam Thương Tín",
		VIET_CREDIT:             "Ngân hàng TMCP Bản Việt",
		VIETCOMBANK:             "Ngân hàng TMCP Ngoại thương Việt Nam",
		VIETINBANK:              "Ngân hàng TMCP Công thương Việt Nam",
		VPBANK:                  "Ngân hàng TMCP Việt Nam Thịnh Vượng",
		VRB:                     "Ngân hàng Liên doanh Việt - Nga",
		WOORI_BANK:              "Ngân hàng TNHH MTV Woori Việt Nam",
	}

	for code, bin := range bankBin {
		name := bankNames[code]
		if name == "" {
			name = code
		}
		banks = append(banks, BankInfo{
			Code: code,
			Name: name,
			Bin:  bin,
		})
	}

	return banks
}
