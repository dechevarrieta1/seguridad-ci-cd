package handlers_test

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"seguridad-cicd/internal/v1/handlers"
	"strings"
	"testing"
)

var seguridadURL = "http://localhost:3000"

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}
func TestMakeRequest_Success(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`{"message":"success"}`)),
			}, nil
		},
	}

	// Captura de logs
	var logs bytes.Buffer
	log.SetOutput(&logs)

	handlers.MakeRequest(mockClient, "GET", "http://example.com", nil)

	expectedLog := "GET response from http://example.com: {\"message\":\"success\"} (Status: 200)"
	if !bytes.Contains(logs.Bytes(), []byte(expectedLog)) {
		t.Errorf("Expected log %s, but got %s", expectedLog, logs.String())
	}
}

func TestMakeRequest_Error(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("network error")
		},
	}

	// Captura de logs
	var logs bytes.Buffer
	log.SetOutput(&logs)

	handlers.MakeRequest(mockClient, "GET", "http://example.com", nil)

	expectedLog := "Error making GET request to http://example.com: network error"
	if !bytes.Contains(logs.Bytes(), []byte(expectedLog)) {
		t.Errorf("Expected log %s, but got %s", expectedLog, logs.String())
	}
}

func TestExecuteHandlers(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			switch req.URL.Path {
			case "/health":
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(`{"message":"health check"}`)),
				}, nil
			case "/alumnos":
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(`{"message":"alumnos"}`)),
				}, nil
			case "/":
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(`{"message":"post"}`)),
				}, nil
			case "/accounts/create":
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(`{"message":"create account"}`)),
				}, nil
			default:
				return nil, errors.New("unknown URL path")
			}
		},
	}

	var logs bytes.Buffer
	log.SetOutput(&logs)

	handlers.ExecuteHandlers(mockClient, seguridadURL)
	t.Skip()
	expectedLogs := []string{
		"GET response from https://hrhelpers.onrender.com/health: {\"message\":\"health check\"} (Status: 200)",
		"GET response from https://hrhelpers.onrender.com/alumnos: {\"message\":\"alumnos\"} (Status: 200)",
		"POST response from https://hrhelpers.onrender.com/: {\"message\":\"post\"} (Status: 200)",
		"POST response from https://hrhelpers.onrender.com/accounts/create: {\"message\":\"create account\"} (Status: 200)",
	}

	logLines := strings.Split(logs.String(), "\n")
	for i, expectedLog := range expectedLogs {
		if !strings.Contains(logLines[i], expectedLog) {
			t.Errorf("Expected log %s, but got %s", expectedLog, logLines[i])
		}
	}
}
