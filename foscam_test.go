package foscam

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/mdmfernandes/go-foscam/mocks"
	"github.com/mdmfernandes/go-foscam/testutil"
)

func TestModel(t *testing.T) {
	tests := []struct {
		name  string
		model Model
		want  string
	}{
		{
			name:  "model by name",
			model: FI9800P,
			want:  "FI9800P",
		},
		{
			name:  "model by number",
			model: 1,
			want:  "FI8919W",
		},
		{
			name:  "invalid model",
			model: 255, // A camera model that does not exist
			want:  "invalid model",
		},
	}

	// Test the String() method of the Model type
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.model.String()

			if got != tt.want {
				t.Errorf("got = %q; want = %q", got, tt.want)
			}
		})
	}
}

func TestNewCamera(t *testing.T) {
	// Example configuration
	cfg := Config{
		URL:      "http://example.com",
		User:     "user",
		Password: "password",
	}

	// Custom client that always returns 200 OK
	clientOk := &mocks.MockHTTPClient{
		GetFunc: func(url string) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNotAcceptable,
			}, nil
		},
	}

	// NewCamera() arguments
	type args struct {
		model   Model
		cfg     Config
		clients []HTTPClient
	}

	tests := []struct {
		name    string
		args    args
		wantCam Camera
		wantErr error
	}{
		{
			name: "ok",
			args: args{
				model: FI8919W,
				cfg:   cfg,
			},
			wantCam: &fi8919w{
				Client:   &http.Client{},
				URL:      cfg.URL,
				User:     cfg.User,
				Password: cfg.Password,
			},
			wantErr: nil,
		},
		{
			name: "ok with custom client",
			args: args{
				model:   FI9800P,
				cfg:     cfg,
				clients: []HTTPClient{clientOk},
			},
			wantCam: &fi9800p{
				Client:   clientOk,
				URL:      cfg.URL,
				User:     cfg.User,
				Password: cfg.Password,
			},
			wantErr: nil,
		},
		{
			name: "invalid model",
			args: args{
				model: 255, // A camera model that does not exist
				cfg:   cfg,
			},
			wantErr: ErrCameraInvalidModel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new camera
			cam, err := NewCamera(tt.args.model, tt.args.cfg, tt.args.clients...)

			if !testutil.EqualError(t, err, tt.wantErr) {
				t.Errorf("Error: got = %v; want = %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(cam, tt.wantCam) {
				t.Errorf("Camera: got = %q; want = %q", cam, tt.wantCam)
			}
		})
	}
}
