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
		name          string
		args          args
		Client        *mocks.MockHTTPClient
		wantErr       error
		wantGetCalled int
	}{
		{
			name: "ok",
			Client: &mocks.MockHTTPClient{
				GetFunc: func(url string) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader("ok.\n")),
					}, nil
				},
			},
			wantErr:       nil,
			wantGetCalled: 1,
		},
		{
			name: "Camera error",
			Client: &mocks.MockHTTPClient{
				GetFunc: func(url string) (*http.Response, error) {
					return nil, http.ErrAbortHandler
				},
			},
			wantErr:       &CameraError{http.ErrAbortHandler.Error()},
			wantGetCalled: 1,
		},
		{
			name: "Get status code unexpected",
			Client: &mocks.MockHTTPClient{
				GetFunc: func(url string) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader("ok.\n")),
					}, nil
				},
			},
			wantErr:       &BadStatusError{Status: http.StatusNotFound, Expected: http.StatusOK},
			wantGetCalled: 1,
		},
		{
			name: "Get unexpected response",
			Client: &mocks.MockHTTPClient{
				GetFunc: func(url string) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader("notok.\n")),
					}, nil
				},
			},
			wantErr:       &BadResponseError{"ok.\n", "notok.\n"},
			wantGetCalled: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cam := &fi8919w{
				Client: tt.Client,
			}

			if err := cam.ChangeMotionStatus(tt.args.enable); !testutil.EqualError(t, err, tt.wantErr) {
				t.Errorf("Error: got = %v; want = %v", err, tt.wantErr)
			}

			if tt.wantGetCalled != tt.Client.GetCount {
				t.Errorf("Number of requests: got = %d; want = %d", tt.Client.GetCount, tt.wantGetCalled)
			}
		})
	}
}
