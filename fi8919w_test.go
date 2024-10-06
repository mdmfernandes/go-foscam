package foscam

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/mdmfernandes/go-foscam/mocks"
	"github.com/mdmfernandes/go-foscam/testutil"
)

func TestFI8919w_ChangeMotionStatus(t *testing.T) {
	// ChangeMotionStatus() arguments
	type args struct {
		enable bool
	}
	tests := []struct {
		name    string
		args    args
		client  *mocks.MockHTTPClient
		wantErr error
	}{
		{
			name: "ok",
			client: &mocks.MockHTTPClient{
				GetFunc: func(url string) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader("ok.\n")),
					}, nil
				},
			},
			wantErr: nil,
		},
		{
			name: "Unexpected response",
			client: &mocks.MockHTTPClient{
				GetFunc: func(url string) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader("notok.\n")),
					}, nil
				},
			},
			wantErr: &BadResponseError{"ok.\n", "notok.\n"},
		},
		{
			name: "Bad request",
			client: &mocks.MockHTTPClient{
				GetFunc: func(url string) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Body:       io.NopCloser(nil),
					}, nil
				},
			},
			wantErr: &BadStatusError{URL: "/set_alarm.cgi?pwd=&user=&motion_armed=0", Status: http.StatusForbidden, Expected: http.StatusOK},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cam := &fi8919w{
				Client: tt.client,
			}

			if err := cam.ChangeMotionStatus(tt.args.enable); !testutil.EqualError(t, err, tt.wantErr) {
				t.Errorf("Error: got = %v; want = %v", err, tt.wantErr)
			}

			if tt.client.GetCount != 1 {
				t.Errorf("Number of requests: got = %d; want = 1", tt.client.GetCount)
			}
		})
	}
}
