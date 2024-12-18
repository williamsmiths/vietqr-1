package vietqr

const (
	ABBANK                  = "ABBANK"
	ACB                     = "ACB"
	AGRIBANK                = "AGRIBANK"
	BAC_A_BANK              = "BAC_A_BANK"
	BAOVIET_BANK            = "BAOVIET_BANK"
	BANVIET                 = "BANVIET"
	BIDV                    = "BIDV"
	CAKE                    = "CAKE"
	CBBANK                  = "CBBANK"
	CIMB                    = "CIMB"
	COOP_BANK               = "COOP_BANK"
	DBS_BANK                = "DBS_BANK"
	DONG_A_BANK             = "DONG_A_BANK"
	EXIMBANK                = "EXIMBANK"
	GPBANK                  = "GPBANK"
	HDBANK                  = "HDBANK"
	HONGLEONG_BANK          = "HONGLEONG_BANK"
	HSBC                    = "HSBC"
	IBK_HCM                 = "IBK_HCM"
	IBK_HN                  = "IBK_HN"
	INDOVINA_BANK           = "INDOVINA_BANK"
	KASIKORN_BANK           = "KASIKORN_BANK"
	KEB_HANA_BANK_HCM       = "KEB_HANA_BANK_HCM"
	KEB_HANA_BANK_HN        = "KEB_HANA_BANK_HN"
	KIENLONG_BANK           = "KIENLONG_BANK"
	KOOKMIN_BANK_HCM        = "KOOKMIN_BANK_HCM"
	KOOKMIN_BANK_HN         = "KOOKMIN_BANK_HN"
	LIENVIETPOST_BANK       = "LIENVIETPOST_BANK"
	MBBANK                  = "MBBANK"
	MB_SHINSEI              = "MB_SHINSEI"
	MIRAE_ASSET             = "MIRAE_ASSET"
	MSB                     = "MSB"
	NAM_A_BANK              = "NAM_A_BANK"
	NCB                     = "NCB"
	NONGHYUP_BANK_HN        = "NONGHYUP_BANK_HN"
	OCB                     = "OCB"
	OCEANBANK               = "OCEANBANK"
	PGBANK                  = "PGBANK"
	PUBLIC_BANK             = "PUBLIC_BANK"
	PVCOM_BANK              = "PVCOM_BANK"
	SACOMBANK               = "SACOMBANK"
	SAIGONBANK              = "SAIGONBANK"
	SCB                     = "SCB"
	SEA_BANK                = "SEA_BANK"
	SHB                     = "SHB"
	SHINHAN_BANK            = "SHINHAN_BANK"
	SINOPAC_BANK_HCM        = "SINOPAC_BANK_HCM"
	STANDARD_CHARTERED_BANK = "STANDARD_CHARTERED_BANK"
	TECHCOMBANK             = "TECHCOMBANK"
	TIMO                    = "TIMO"
	TNEX                    = "TNEX"
	TPBANK                  = "TPBANK"
	UBANK                   = "UBANK"
	UNITED_OVERSEAS_BANK    = "UNITED_OVERSEAS_BANK"
	VIB                     = "VIB"
	VIET_A_BANK             = "VIET_A_BANK"
	VIET_BANK               = "VIET_BANK"
	VIET_CREDIT             = "VIET_CREDIT" // CFC
	VIETCOMBANK             = "VIETCOMBANK"
	VIETINBANK              = "VIETINBANK"
	VPBANK                  = "VPBANK"
	VRB                     = "VRB"
	WOORI_BANK              = "WOORI_BANK"
)

// follow latest data at: https://www.sbv.gov.vn/webcenter/portal/vi/menu/trangchu/ttvnq/htmtcqht
var bankBin = map[string]string{
	HSBC:                    "458761",
	CAKE:                    "546034",
	UBANK:                   "546035",
	KASIKORN_BANK:           "668888",
	DBS_BANK:                "796500",
	NONGHYUP_BANK_HN:        "801011",
	TIMO:                    "963388",
	SAIGONBANK:              "970400",
	SACOMBANK:               "970403",
	AGRIBANK:                "970405",
	DONG_A_BANK:             "970406",
	TECHCOMBANK:             "970407",
	GPBANK:                  "970408",
	BAC_A_BANK:              "970409",
	STANDARD_CHARTERED_BANK: "970410",
	PVCOM_BANK:              "970412",
	OCEANBANK:               "970414",
	VIETINBANK:              "970415",
	ACB:                     "970416",
	BIDV:                    "970418",
	NCB:                     "970419",
	VRB:                     "970421",
	MBBANK:                  "970422",
	TPBANK:                  "970423",
	SHINHAN_BANK:            "970424",
	ABBANK:                  "970425",
	MSB:                     "970426",
	VIET_A_BANK:             "970427",
	NAM_A_BANK:              "970428",
	SCB:                     "970429",
	PGBANK:                  "970430",
	EXIMBANK:                "970431",
	VPBANK:                  "970432",
	VIET_BANK:               "970433",
	INDOVINA_BANK:           "970434",
	VIETCOMBANK:             "970436",
	HDBANK:                  "970437",
	BAOVIET_BANK:            "970438",
	PUBLIC_BANK:             "970439",
	SEA_BANK:                "970440",
	VIB:                     "970441",
	HONGLEONG_BANK:          "970442",
	SHB:                     "970443",
	CBBANK:                  "970444",
	COOP_BANK:               "970446",
	OCB:                     "970448",
	LIENVIETPOST_BANK:       "970449",
	KIENLONG_BANK:           "970452",
	BANVIET:                 "970454",
	IBK_HN:                  "970455",
	IBK_HCM:                 "970456",
	WOORI_BANK:              "970457",
	UNITED_OVERSEAS_BANK:    "970458",
	CIMB:                    "970459", // 422589
	VIET_CREDIT:             "970460",
	KOOKMIN_BANK_HN:         "970462",
	KOOKMIN_BANK_HCM:        "970463",
	TNEX:                    "970464",
	SINOPAC_BANK_HCM:        "970465",
	KEB_HANA_BANK_HCM:       "970466",
	KEB_HANA_BANK_HN:        "970467",
	MIRAE_ASSET:             "970468",
	MB_SHINSEI:              "970470",
}

var revBankBin = map[string]string{
	bankBin[ABBANK]:                  ABBANK,
	bankBin[ACB]:                     ACB,
	bankBin[AGRIBANK]:                AGRIBANK,
	bankBin[BAC_A_BANK]:              BAC_A_BANK,
	bankBin[BAOVIET_BANK]:            BAOVIET_BANK,
	bankBin[BANVIET]:                 BANVIET,
	bankBin[BIDV]:                    BIDV,
	bankBin[CAKE]:                    CAKE,
	bankBin[CBBANK]:                  CBBANK,
	bankBin[CIMB]:                    CIMB,
	bankBin[COOP_BANK]:               COOP_BANK,
	bankBin[DBS_BANK]:                DBS_BANK,
	bankBin[DONG_A_BANK]:             DONG_A_BANK,
	bankBin[EXIMBANK]:                EXIMBANK,
	bankBin[GPBANK]:                  GPBANK,
	bankBin[HDBANK]:                  HDBANK,
	bankBin[HONGLEONG_BANK]:          HONGLEONG_BANK,
	bankBin[HSBC]:                    HSBC,
	bankBin[IBK_HCM]:                 IBK_HCM,
	bankBin[IBK_HN]:                  IBK_HN,
	bankBin[INDOVINA_BANK]:           INDOVINA_BANK,
	bankBin[KASIKORN_BANK]:           KASIKORN_BANK,
	bankBin[KEB_HANA_BANK_HCM]:       KEB_HANA_BANK_HCM,
	bankBin[KEB_HANA_BANK_HN]:        KEB_HANA_BANK_HN,
	bankBin[KIENLONG_BANK]:           KIENLONG_BANK,
	bankBin[KOOKMIN_BANK_HCM]:        KOOKMIN_BANK_HCM,
	bankBin[KOOKMIN_BANK_HN]:         KOOKMIN_BANK_HN,
	bankBin[LIENVIETPOST_BANK]:       LIENVIETPOST_BANK,
	bankBin[MBBANK]:                  MBBANK,
	bankBin[MB_SHINSEI]:              MB_SHINSEI,
	bankBin[MIRAE_ASSET]:             MIRAE_ASSET,
	bankBin[MSB]:                     MSB,
	bankBin[NAM_A_BANK]:              NAM_A_BANK,
	bankBin[NCB]:                     NCB,
	bankBin[NONGHYUP_BANK_HN]:        NONGHYUP_BANK_HN,
	bankBin[OCB]:                     OCB,
	bankBin[OCEANBANK]:               OCEANBANK,
	bankBin[PGBANK]:                  PGBANK,
	bankBin[PUBLIC_BANK]:             PUBLIC_BANK,
	bankBin[PVCOM_BANK]:              PVCOM_BANK,
	bankBin[SACOMBANK]:               SACOMBANK,
	bankBin[SAIGONBANK]:              SAIGONBANK,
	bankBin[SCB]:                     SCB,
	bankBin[SEA_BANK]:                SEA_BANK,
	bankBin[SHB]:                     SHB,
	bankBin[SHINHAN_BANK]:            SHINHAN_BANK,
	bankBin[SINOPAC_BANK_HCM]:        SINOPAC_BANK_HCM,
	bankBin[STANDARD_CHARTERED_BANK]: STANDARD_CHARTERED_BANK,
	bankBin[TECHCOMBANK]:             TECHCOMBANK,
	bankBin[TIMO]:                    TIMO,
	bankBin[TNEX]:                    TNEX,
	bankBin[TPBANK]:                  TPBANK,
	bankBin[UBANK]:                   UBANK,
	bankBin[UNITED_OVERSEAS_BANK]:    UNITED_OVERSEAS_BANK,
	bankBin[VIB]:                     VIB,
	bankBin[VIET_A_BANK]:             VIET_A_BANK,
	bankBin[VIET_BANK]:               VIET_BANK,
	bankBin[VIET_CREDIT]:             VIET_CREDIT,
	bankBin[VIETCOMBANK]:             VIETCOMBANK,
	bankBin[VIETINBANK]:              VIETINBANK,
	bankBin[VPBANK]:                  VPBANK,
	bankBin[VRB]:                     VRB,
	bankBin[WOORI_BANK]:              WOORI_BANK,
}
