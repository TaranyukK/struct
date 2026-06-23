package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestClient(serverURL string) *Client {
	c := NewClient("test-key")
	return c
}

func TestCreate(t *testing.T) {
	var capturedMethod, capturedPath, capturedKey, capturedContentType string
	var capturedBody []byte

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedMethod = r.Method
		capturedPath = r.URL.Path
		capturedKey = r.Header.Get("X-Master-Key")
		capturedContentType = r.Header.Get("Content-Type")
		capturedBody, _ = io.ReadAll(r.Body)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"record": {"id": "bin123"}}`))
	}))
	defer server.Close()

	c := NewClient("test-api-key")
	data := map[string]string{"name": "test"}

	err := c.sendRequest("POST", server.URL+"/v3/b", data)
	if err != nil {
		t.Fatalf("Create returned error: %v", err)
	}

	if capturedMethod != "POST" {
		t.Errorf("Expected POST, got %s", capturedMethod)
	}
	if capturedPath != "/v3/b" {
		t.Errorf("Expected /v3/b, got %s", capturedPath)
	}
	if capturedKey != "test-api-key" {
		t.Errorf("Expected X-Master-Key = test-api-key, got %s", capturedKey)
	}
	if capturedContentType != "application/json" {
		t.Errorf("Expected Content-Type = application/json, got %s", capturedContentType)
	}

	var got map[string]string
	json.Unmarshal(capturedBody, &got)
	if got["name"] != "test" {
		t.Errorf("Expected body name=test, got %v", got)
	}
}

func TestUpdate(t *testing.T) {
	var capturedMethod, capturedPath string
	var capturedBody []byte

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedMethod = r.Method
		capturedPath = r.URL.Path
		capturedBody, _ = io.ReadAll(r.Body)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"record": {"id": "bin123"}}`))
	}))
	defer server.Close()

	c := NewClient("test-api-key")
	data := map[string]string{"name": "updated"}

	err := c.sendRequest("PUT", server.URL+"/v3/b/bin123", data)
	if err != nil {
		t.Fatalf("Update returned error: %v", err)
	}

	if capturedMethod != "PUT" {
		t.Errorf("Expected PUT, got %s", capturedMethod)
	}
	if capturedPath != "/v3/b/bin123" {
		t.Errorf("Expected /v3/b/bin123, got %s", capturedPath)
	}

	var got map[string]string
	json.Unmarshal(capturedBody, &got)
	if got["name"] != "updated" {
		t.Errorf("Expected body name=updated, got %v", got)
	}
}

func TestGet(t *testing.T) {
	var capturedMethod, capturedPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedMethod = r.Method
		capturedPath = r.URL.Path

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"record": {"id": "bin123", "name": "test"}}`))
	}))
	defer server.Close()

	c := NewClient("test-api-key")

	err := c.sendRequest("GET", server.URL+"/v3/b/bin123/latest", nil)
	if err != nil {
		t.Fatalf("Get returned error: %v", err)
	}

	if capturedMethod != "GET" {
		t.Errorf("Expected GET, got %s", capturedMethod)
	}
	if capturedPath != "/v3/b/bin123/latest" {
		t.Errorf("Expected /v3/b/bin123/latest, got %s", capturedPath)
	}
}

func TestDelete(t *testing.T) {
	var capturedMethod, capturedPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedMethod = r.Method
		capturedPath = r.URL.Path

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "deleted"}`))
	}))
	defer server.Close()

	c := NewClient("test-api-key")

	err := c.sendRequest("DELETE", server.URL+"/v3/b/bin123", nil)
	if err != nil {
		t.Fatalf("Delete returned error: %v", err)
	}

	if capturedMethod != "DELETE" {
		t.Errorf("Expected DELETE, got %s", capturedMethod)
	}
	if capturedPath != "/v3/b/bin123" {
		t.Errorf("Expected /v3/b/bin123, got %s", capturedPath)
	}
}

func TestSendRequest_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "invalid key"}`))
	}))
	defer server.Close()

	c := NewClient("bad-key")
	err := c.sendRequest("GET", server.URL+"/v3/b", nil)

	if err == nil {
		t.Fatal("Expected error for 401 response, got nil")
	}
	if !strings.Contains(err.Error(), "ошибка API (код 401)") {
		t.Errorf("Expected error to contain 'ошибка API (код 401)', got: %v", err)
	}
}