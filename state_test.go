package chief_test

import (
	"bytes"
	"io"
	"testing"

	"ireul.com/chief"
)

func TestDecodeState(t *testing.T) {
	type args struct {
		r io.Reader
		s chief.State
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				r: bytes.NewReader([]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x10,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06,
				}),
				s: chief.State{
					Seed:        1,
					ShardLen:    272,
					ShardStart:  1,
					StripeLen:   4,
					StripeStart: 5,
					Index:       6,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := chief.State{}
			if err := chief.DecodeState(tt.args.r, &s); (err != nil) != tt.wantErr {
				t.Errorf("DecodeState() error = %v, wantErr %v", err, tt.wantErr)
			}
			if s != tt.args.s {
				t.Errorf("DecodeState() failed, expect %v, got %v", tt.args.s, s)
			}
		})
	}
}
