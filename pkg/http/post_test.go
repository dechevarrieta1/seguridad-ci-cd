package httphelpers_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	httphelpers "seguridad-cicd/pkg/http"
	"testing"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("error reading response body")
}

func (e *errorReader) Close() error {
	return nil
}
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestRequest_Success(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{"success":true}`)),
			}, nil
		},
	}

	reqData := []byte(`{"test":"data"}`)
	body, status, err := httphelpers.Request(mockClient, reqData, "http://example.com", "POST")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}

	expectedBody := `{"success":true}`
	if string(body) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, body)
	}
}

func TestRequest_Error(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("network error")
		},
	}

	reqData := []byte(`{"test":"data"}`)
	_, status, err := httphelpers.Request(mockClient, reqData, "http://example.com", "POST")
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if status != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", status)
	}
}

func TestRequest_StatusCodeError(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusBadRequest, // Simulamos un 400
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{"error":"bad request"}`)),
			}, nil
		},
	}

	reqData := []byte(`{"test":"data"}`)
	body, status, err := httphelpers.Request(mockClient, reqData, "http://example.com", "POST")

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if status != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", status)
	}

	expectedBody := `{"error":"bad request"}`
	if string(body) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, body)
	}

	expectedErrMsg := "error in response"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message %s, got %s", expectedErrMsg, err.Error())
	}
}

func TestRequest_ResponseNil(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, nil
		},
	}

	reqData := []byte(`{"test":"data"}`)
	_, status, err := httphelpers.Request(mockClient, reqData, "http://example.com", "POST")

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if status != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", status)
	}

	expectedErrMsg := "response is nil"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message %s, got %s", expectedErrMsg, err.Error())
	}
}

func TestRequest_ReadBodyError(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// Simulamos una respuesta válida pero con un body ilegible
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(&errorReader{}), // Devuelve un error al intentar leer
			}, nil
		},
	}

	reqData := []byte(`{"test":"data"}`)
	_, status, err := httphelpers.Request(mockClient, reqData, "http://example.com", "POST")

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if status != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", status)
	}

	expectedErrMsg := "error reading response body"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message %s, got %s", expectedErrMsg, err.Error())
	}
}
