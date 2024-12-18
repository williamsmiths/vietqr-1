package vietqr

import (
	b64 "encoding/base64"
	"testing"
)

func Test_Crc16(t *testing.T) {
	equalBytesSliceFn := func(b1, b2 []byte) bool {
		if len(b1) != len(b2) {
			return false
		}

		for i := 0; i < len(b1); i++ {
			if b1[i] != b2[i] {
				return false
			}
		}
		return true
	}

	sameText := "hello world"
	type args struct {
		params CrcParams
		s      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "params CRC16_ARC",
			args: args{
				CRC16_ARC,
				sameText,
			},
			want: "OcE=",
		},
		{
			name: "params CRC16_AUG_CCITT",
			args: args{
				CRC16_AUG_CCITT,
				sameText,
			},
			want: "E+g=",
		},
		{
			name: "params CRC16_BUYPASS",
			args: args{
				CRC16_BUYPASS,
				sameText,
			},
			want: "V8w=",
		},
		{
			name: "params CRC16_CCITT_FALSE",
			args: args{
				CRC16_CCITT_FALSE,
				sameText,
			},
			want: "7+s=",
		},
		{
			name: "params CRC16_CDMA2000",
			args: args{
				CRC16_CDMA2000,
				sameText,
			},
			want: "+fE=",
		},
		{
			name: "params CRC16_DDS_110",
			args: args{
				CRC16_DDS_110,
				sameText,
			},
			want: "lxs=",
		},
		{
			name: "params CRC16_DECT_R",
			args: args{
				CRC16_DECT_R,
				sameText,
			},
			want: "Chc=",
		},
		{
			name: "params CRC16_DECT_X",
			args: args{
				CRC16_DECT_X,
				sameText,
			},
			want: "ChY=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewCrc16(tt.args.params)
			h.Write([]byte(tt.args.s))
			got := h.Sum(nil)
			want, _ := b64.StdEncoding.DecodeString(tt.want)
			if !equalBytesSliceFn(got, want) {
				t.Errorf("b64-Crc16() = %v, want %v", b64.StdEncoding.EncodeToString(got), tt.want)
			}
		})
	}
}
