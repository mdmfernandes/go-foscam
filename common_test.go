package foscam

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/mdmfernandes/go-foscam/mocks"
	"github.com/mdmfernandes/go-foscam/testutil"
)

func Test_getRequest(t *testing.T) {
	// request() arguments
	type args struct {
		client *mocks.MockHTTPClient
		url    string
	}

	tests := []struct {
		name          string
		args          *args
		want          []byte
		wantErr       error
		wantGetCalled int
	}{
		{
			name: "Camera error",
			args: &args{
				client: &mocks.MockHTTPClient{
					GetFunc: func(url string) (*http.Response, error) {
						return nil, errors.New("some error")
					},
				},
				url: "",
			},
			want:          nil,
			wantErr:       &CameraError{"some error"},
			wantGetCalled: 1,
		},
		{
			name: "Unexpected status code (404)",
			args: &args{
				client: &mocks.MockHTTPClient{
					GetFunc: func(url string) (*http.Response, error) {
						return &http.Response{
							StatusCode: http.StatusNotFound,
							Body:       io.NopCloser(bytes.NewBufferString("don't care")),
						}, nil
					},
				},
				url: "",
			},
			want:          nil,
			wantErr:       &BadStatusError{"", http.StatusNotFound, http.StatusOK},
			wantGetCalled: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := getRequest(tt.args.client, tt.args.url)

			// Check how many times the client was called
			if tt.wantGetCalled != tt.args.client.GetCount {
				t.Errorf("Number of requests: got = %d; want = %d", tt.args.client.GetCount, tt.wantGetCalled)
			}

			// Check for errors
			if !testutil.EqualError(t, err, tt.wantErr) {
				t.Errorf("Error: got = %v; want = %v", err, tt.wantErr)
			}

			// Check for the response body
			if !bytes.Equal(b, tt.want) {
				t.Errorf("Response: got = %v; want = %v", b, tt.want)
			}
		})
	}
}
