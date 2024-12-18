package vietqr

import (
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		ti TransferInfo
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "only beneficiary info",
			args: args{
				TransferInfo{
					BankCode: CAKE,
					BankNo:   "0905555999",
				},
			},
			want: "00020101021138540010A00000072701240006546034011009055559990208QRIBFTTA53037045802VN6304B2A0",
		},
		{
			name: "VNPAY",
			args: args{
				TransferInfo{
					merchantID: "VNP-123456",
				},
			},
			want: "00020101021138400010A0000007750110VNP-1234560208QRIBFTTA53037045802VN6304E56C",
		},
		{
			name: "with amount",
			args: args{
				TransferInfo{
					BankCode: ACB,
					BankNo:   "13579",
					Amount:   120000,
				},
			},
			want: "00020101021238490010A000000727011900069704160105135790208QRIBFTTA530370454061200005802VN63049A71",
		},
		{
			name: "with message",
			args: args{
				TransferInfo{
					BankCode: TECHCOMBANK,
					BankNo:   "9796868",
					Message:  "gen by sunary/vietqr",
				},
			},
			want: "00020101021138510010A00000072701210006970407010797968680208QRIBFTTA53037045802VN62240820gen by sunary/vietqr6304BE74",
		},
		{
			name: "full info",
			args: args{
				TransferInfo{
					BankCode: VPBANK,
					BankNo:   "19372",
					Amount:   152000,
					Message:  "gen by go-vietqr",
				},
			},
			want: "00020101021238490010A000000727011900069704320105193720208QRIBFTTA530370454061520005802VN62200816gen by go-vietqr63040ED4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.ti); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	equalTransferInfoFn := func(t1, t2 *TransferInfo) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil {
			return false
		}
		return t1.merchantID == t2.merchantID &&
			t1.BankCode == t2.BankCode &&
			t1.BankNo == t2.BankNo &&
			t1.Amount == t2.Amount &&
			t1.Message == t2.Message
	}

	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *TransferInfo
		wantErr bool
	}{
		{
			name: "wrong crc",
			args: args{
				"00020101021138540010A00000072701240006546034011009055559990208QRIBFTTA53037045802VN6304xxxx",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "guid VNPAY",
			args: args{
				"00020101021138400010A0000007750110VNP-1234560208QRIBFTTA53037045802VN6304E56C",
			},
			want: &TransferInfo{
				merchantID: "VNP-123456",
			},
			wantErr: false,
		},
		{
			name: "missing crc",
			args: args{
				"00020101021138540010A00000072701240006546034011009055559990208QRIBFTTA53037045802VN6304",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "correct data: only beneficiary info",
			args: args{
				"00020101021138540010A00000072701240006546034011009055559990208QRIBFTTA53037045802VN6304B2A0",
			},
			want: &TransferInfo{
				BankCode: CAKE,
				BankNo:   "0905555999",
			},
			wantErr: false,
		},
		{
			name: "with amount",
			args: args{
				"00020101021238490010A000000727011900069704160105135790208QRIBFTTA530370454061200005802VN63049A71",
			},
			want: &TransferInfo{
				BankCode: ACB,
				BankNo:   "13579",
				Amount:   120000,
			},
			wantErr: false,
		},
		{
			name: "correct data: with message",
			args: args{
				"00020101021138510010A00000072701210006970407010797968680208QRIBFTTA53037045802VN62240820gen by sunary/vietqr6304BE74",
			},
			want: &TransferInfo{
				BankCode: TECHCOMBANK,
				BankNo:   "9796868",
				Message:  "gen by sunary/vietqr",
			},
			wantErr: false,
		},
		{
			name: "correct data: full info",
			args: args{
				"00020101021238490010A000000727011900069704320105193720208QRIBFTTA530370454061520005802VN62200816gen by go-vietqr63040ED4",
			},
			want: &TransferInfo{
				BankCode: VPBANK,
				BankNo:   "19372",
				Amount:   152000,
				Message:  "gen by go-vietqr",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !equalTransferInfoFn(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
