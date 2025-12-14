package runner

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/MichelFortes/httpit/pkg/model"
)

func TestRunner_Run_Success(t *testing.T) {
	// 1. Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assert request if needed
		if r.URL.Path != "/success" {
			t.Errorf("Expected path /success, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Parse server URL to extract host and port
	serverURL, _ := url.Parse(server.URL)
	port, _ := strconv.Atoi(serverURL.Port())

	// 2. Prepare test scheme
	scheme := &model.TestScheme{
		Protocol: "http",
		Host:     serverURL.Hostname(),
		Port:     uint16(port),
		Tests: []model.Test{
			{
				Description: "Success Case",
				Path:        "/success",
				Method:      "GET",
				ExpectedResult: model.ExpectedResult{
					StatusCode: 200,
				},
			},
		},
	}

	// 3. Run Runner
	r := NewRunner()
	err := r.Run(scheme)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestRunner_Run_ValidationFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError) // 500
	}))
	defer server.Close()

	serverURL, _ := url.Parse(server.URL)
	port, _ := strconv.Atoi(serverURL.Port())

	scheme := &model.TestScheme{
		Protocol: "http",
		Host:     serverURL.Hostname(),
		Port:     uint16(port),
		Tests: []model.Test{
			{
				Description: "Failure Case",
				Path:        "/",
				Method:      "GET",
				ExpectedResult: model.ExpectedResult{
					StatusCode: 200, // Expect 200, but get 500
				},
			},
		},
	}

	r := NewRunner()
	// Run currently only logs failures, doesn't return error on validation fail by design (yet).
	// So we assert it runs without crashing.
	err := r.Run(scheme)
	if err != nil {
		t.Fatalf("Expected no error from Run execution, got %v", err)
	}
}
