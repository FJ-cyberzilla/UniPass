package geo

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMockGeoIPProvider(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ip":"127.0.0.1","country":"Localhost"}`))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	client := &http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to execute mock HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK 200, got %d", resp.StatusCode)
	}
}
