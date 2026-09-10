package notification

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// Spin up a fake "Fonnte" server and assert our notifier sends the right
// method, headers, and form fields. This proves the integration is wired
// correctly without needing a real token or a connected WhatsApp device.
func TestFonnteNotifier_BuildsCorrectRequest(t *testing.T) {
	var method, auth, contentType, target, message string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		auth = r.Header.Get("Authorization")
		contentType = r.Header.Get("Content-Type")
		raw, _ := io.ReadAll(r.Body)
		vals, _ := url.ParseQuery(string(raw))
		target = vals.Get("target")
		message = vals.Get("message")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":true}`))
	}))
	defer srv.Close()

	n := NewFonnte("SECRET123", srv.URL)
	if err := n.Send(context.Background(), "628123456789", "halo dunia"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if method != http.MethodPost {
		t.Errorf("method = %s, want POST", method)
	}
	if auth != "SECRET123" {
		t.Errorf("Authorization = %q, want raw token without Bearer", auth)
	}
	if !strings.Contains(contentType, "x-www-form-urlencoded") {
		t.Errorf("Content-Type = %q, want form-urlencoded", contentType)
	}
	if target != "628123456789" {
		t.Errorf("target = %q", target)
	}
	if message != "halo dunia" {
		t.Errorf("message = %q", message)
	}
}

func TestFonnteNotifier_ErrorsOnNon200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer srv.Close()

	n := NewFonnte("bad-token", srv.URL)
	if err := n.Send(context.Background(), "628", "x"); err == nil {
		t.Fatal("expected an error on non-200 response, got nil")
	}
}
