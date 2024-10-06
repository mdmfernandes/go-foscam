package foscam

import (
	"testing"

	"github.com/mdmfernandes/go-foscam/testutil"
)

func Test_isJpeg(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		wantErr error
	}{
		{
			name:    "JPEG image",
			data:    []byte("\xFF\xD8\xFF"),
			wantErr: nil,
		},
		{
			name:    "Plain text",
			data:    []byte(`This is plain text`),
			wantErr: &InvalidMIMETypeError{Want: jpegMime, Got: "text/plain; charset=utf-8"},
		},
		{
			name:    "Empty data",
			data:    []byte{},
			wantErr: &InvalidMIMETypeError{Want: jpegMime, Got: "text/plain; charset=utf-8"},
		},
		{
			name:    "Unknown format",
			data:    []byte{1, 2},
			wantErr: &InvalidMIMETypeError{Want: jpegMime, Got: "application/octet-stream"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isJpeg(tt.data); !testutil.EqualError(t, err, tt.wantErr) {
				t.Errorf("Error: got = %v; want = %v", err, tt.wantErr)
			}
		})
	}
}

func Test_b2u(t *testing.T) {
	type args struct {
		b bool
	}

	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "true",
			args: args{true},
			want: 1,
		},
		{
			name: "false",
			args: args{false},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b2u(tt.args.b); got != tt.want {
				t.Errorf("b2u() = %v, want %v", got, tt.want)
			}
		})
	}
}
