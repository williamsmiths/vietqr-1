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
	KIENLONG_BANK           = "KIENLONG_BANK"
	KOOKMIN_BANK_HCM        = "KOOKMIN_BANK_HCM"
	KOOKMIN_BANK_HN         = "KOOKMIN_BANK_HN"
	LIENVIETPOST_BANK       = "LIENVIETPOST_BANK"
	MBBANK                  = "MBBANK"
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
	STANDARD_CHARTERED_BANK = "STANDARD_CHARTERED_BANK"
	TECHCOMBANK             = "TECHCOMBANK"
	TIMO                    = "TIMO"
	TPBANK                  = "TPBANK"
	UBANK                   = "UBANK"
	UNITED_OVERSEAS_BANK    = "UNITED_OVERSEAS_BANK"
	VIB                     = "VIB"
	VIET_A_BANK             = "VIET_A_BANK"
	VIET_BANK               = "VIET_BANK"
	VIETCOMBANK             = "VIETCOMBANK"
	VIETINBANK              = "VIETINBANK"
	VPBANK                  = "VPBANK"
	VRB                     = "VRB"
	WOORI_BANK              = "WOORI_BANK"
)

var bankBin = map[string]string{
	ABBANK:                  "970425",
	ACB:                     "970416",
	AGRIBANK:                "970405",
	BAC_A_BANK:              "970409",
	BAOVIET_BANK:            "970438",
	BANVIET:                 "970454",
	BIDV:                    "970418",
	CAKE:                    "546034",
	CBBANK:                  "970444",
	CIMB:                    "422589",
	COOP_BANK:               "970446",
	DBS_BANK:                "796500",
	DONG_A_BANK:             "970406",
	EXIMBANK:                "970431",
	GPBANK:                  "970408",
	HDBANK:                  "970437",
	HONGLEONG_BANK:          "970442",
	HSBC:                    "458761",
	IBK_HCM:                 "970456",
	IBK_HN:                  "970455",
	INDOVINA_BANK:           "970434",
	KASIKORN_BANK:           "668888",
	KIENLONG_BANK:           "970452",
	KOOKMIN_BANK_HCM:        "970463",
	KOOKMIN_BANK_HN:         "970462",
	LIENVIETPOST_BANK:       "970449",
	MBBANK:                  "970422",
	MSB:                     "970426",
	NAM_A_BANK:              "970428",
	NCB:                     "970419",
	NONGHYUP_BANK_HN:        "801011",
	OCB:                     "970448",
	OCEANBANK:               "970414",
	PGBANK:                  "970430",
	PUBLIC_BANK:             "970439",
	PVCOM_BANK:              "970412",
	SACOMBANK:               "970403",
	SAIGONBANK:              "970400",
	SCB:                     "970429",
	SEA_BANK:                "970440",
	SHB:                     "970443",
	SHINHAN_BANK:            "970424",
	STANDARD_CHARTERED_BANK: "970410",
	TECHCOMBANK:             "970407",
	TIMO:                    "963388",
	TPBANK:                  "970423",
	UBANK:                   "546035",
	UNITED_OVERSEAS_BANK:    "970458",
	VIB:                     "970441",
	VIET_A_BANK:             "970427",
	VIET_BANK:               "970433",
	VIETCOMBANK:             "970436",
	VIETINBANK:              "970415",
	VPBANK:                  "970432",
	VRB:                     "970421",
	WOORI_BANK:              "970457",
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
	bankBin[KIENLONG_BANK]:           KIENLONG_BANK,
	bankBin[KOOKMIN_BANK_HCM]:        KOOKMIN_BANK_HCM,
	bankBin[KOOKMIN_BANK_HN]:         KOOKMIN_BANK_HN,
	bankBin[LIENVIETPOST_BANK]:       LIENVIETPOST_BANK,
	bankBin[MBBANK]:                  MBBANK,
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
	bankBin[STANDARD_CHARTERED_BANK]: STANDARD_CHARTERED_BANK,
	bankBin[TECHCOMBANK]:             TECHCOMBANK,
	bankBin[TIMO]:                    TIMO,
	bankBin[TPBANK]:                  TPBANK,
	bankBin[UBANK]:                   UBANK,
	bankBin[UNITED_OVERSEAS_BANK]:    UNITED_OVERSEAS_BANK,
	bankBin[VIB]:                     VIB,
	bankBin[VIET_A_BANK]:             VIET_A_BANK,
	bankBin[VIET_BANK]:               VIET_BANK,
	bankBin[VIETCOMBANK]:             VIETCOMBANK,
	bankBin[VIETINBANK]:              VIETINBANK,
	bankBin[VPBANK]:                  VPBANK,
	bankBin[VRB]:                     VRB,
	bankBin[WOORI_BANK]:              WOORI_BANK,
}
