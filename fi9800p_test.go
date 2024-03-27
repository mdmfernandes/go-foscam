package foscam

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/mdmfernandes/go-foscam/mocks"
	"github.com/mdmfernandes/go-foscam/testutil"
)

func TestFI9800p_GetMotionDetect(t *testing.T) {
	tests := []struct {
		name          string
		Client        *mocks.MockHTTPClient
		want          fi9800pMotion
		wantErr       error
		wantGetCalled int
	}{
		{
			name: "Motion config ok",
			Client: &mocks.MockHTTPClient{
				GetFunc: func(_ string) (*http.Response, error) {
					xml := `<CGI_Result>
									<result>0</result>
									<isEnable>1</isEnable>
									<area4>69</area4>
								</CGI_Result>`
					b := io.NopCloser(strings.NewReader(xml))
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       b,
					}, nil
				},
			},
			want: fi9800pMotion{
				XMLName:  xml.Name{Local: "CGI_Result"},
				IsEnable: 1,
				Area4:    69,
			},
			wantErr:       nil,
			wantGetCalled: 1,
		},
		{
			name: "Camera error",
			Client: &mocks.MockHTTPClient{
				GetFunc: func(url string) (*http.Response, error) {
					return nil, errors.New("some error")
				},
			},
			want:          fi9800pMotion{},
			wantErr:       &CameraError{"some error"},
			wantGetCalled: 1,
		},
		{
			name: "Unexpected status code (404)",
			Client: &mocks.MockHTTPClient{
				GetFunc: func(url string) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       nil,
					}, nil
				},
			},
			wantErr:       &BadStatusError{Status: http.StatusNotFound, Expected: http.StatusOK},
			wantGetCalled: 1,
		},
		{
			name: "Invalid XML response", // Didn't close CGI_Result
			Client: &mocks.MockHTTPClient{
				GetFunc: func(_ string) (*http.Response, error) {
					xml := `<CGI_Result>
									<result>0</result>
									<isEnable>1</isEnable>
									<area4>69</area4>`
					b := io.NopCloser(strings.NewReader(xml))
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       b,
					}, nil
				},
			},
			want: fi9800pMotion{
				XMLName:  xml.Name{Local: "CGI_Result"},
				IsEnable: 1,
				Area4:    69,
			},
			wantErr:       &xml.SyntaxError{Line: 4, Msg: "unexpected EOF"},
			wantGetCalled: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cam := &fi9800p{
				Client: tt.Client,
			}
			got, err := cam.GetMotionDetect()

			if !testutil.EqualError(t, err, tt.wantErr) {
				t.Errorf("Error: got = %v; want = %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got = %q; want = %q", got, tt.want)
			}

			if tt.wantGetCalled != tt.Client.GetCount {
				t.Errorf("Number of requests: got = %d; want = %d", tt.Client.GetCount, tt.wantGetCalled)
			}
		})
	}
}

func TestFI9800p_updateMotionDetect(t *testing.T) {
	// updateMotionDetect() arguments
	type args struct {
		mc fi9800pMotion
	}

	tests := []struct {
		name          string
		args          args
		Client        *mocks.MockHTTPClient
		wantErr       error
		wantGetCalled int
	}{
		{
			name: "update ok",
			Client: &mocks.MockHTTPClient{
				GetFunc: func(_ string) (*http.Response, error) {
					xml := `<CGI_Result>
									<result>0</result>
								</CGI_Result>`
					b := io.NopCloser(strings.NewReader(xml))
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       b,
					}, nil
				},
			},
			wantErr:       nil,
			wantGetCalled: 1,
		},
		{
			name: "update fails (result not 0)", // result != 0
			Client: &mocks.MockHTTPClient{
				GetFunc: func(_ string) (*http.Response, error) {
					xml := `<CGI_Result>
									<result>1</result>
								</CGI_Result>`
					b := io.NopCloser(strings.NewReader(xml))
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       b,
					}, nil
				},
			},
			wantErr:       &BadResponseError{0, 1},
			wantGetCalled: 1,
		},
		{
			name: "Camera error",
			Client: &mocks.MockHTTPClient{
				GetFunc: func(url string) (*http.Response, error) {
					return nil, errors.New("some error")
				},
			},
			wantErr:       &CameraError{"some error"},
			wantGetCalled: 1,
		},
		{
			name: "Unexpected status code (404)",
			Client: &mocks.MockHTTPClient{
				GetFunc: func(url string) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       nil,
					}, nil
				},
			},
			wantErr:       &BadStatusError{"", http.StatusNotFound, http.StatusOK},
			wantGetCalled: 1,
		},
		{
			name: "Invalid XML response", // Didn't close CGI_Result
			Client: &mocks.MockHTTPClient{
				GetFunc: func(_ string) (*http.Response, error) {
					xml := `<CGI_Result>
											<result>0</result>`
					b := io.NopCloser(strings.NewReader(xml))
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       b,
					}, nil
				},
			},
			wantErr:       &xml.SyntaxError{Line: 2, Msg: "unexpected EOF"},
			wantGetCalled: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cam := &fi9800p{
				Client: tt.Client,
			}

			if err := cam.updateMotionDetect(tt.args.mc); !testutil.EqualError(t, err, tt.wantErr) {
				t.Errorf("Error: got = %v; want = %v", err, tt.wantErr)
			}

			if tt.wantGetCalled != tt.Client.GetCount {
				t.Errorf("Number of requests: got = %d; want = %d", tt.Client.GetCount, tt.wantGetCalled)
			}
		})
	}
}

func TestFI9800p_ChangeMotionStatus(t *testing.T) {
	t.SkipNow()
}
