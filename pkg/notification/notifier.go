package notification

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"portaljob/internal/domain"
)

// StubNotifier just logs. Default in local/dev so the app runs without any real
// WA credentials.
type StubNotifier struct{}

func NewStub() domain.Notifier { return &StubNotifier{} }

func (s *StubNotifier) Send(ctx context.Context, to, message string) error {
	log.Printf("[WA stub] to=%s message=%q", to, message)
	return nil
}

// FonnteNotifier sends a WhatsApp message via the Fonnte gateway.
// Fonnte expects application/x-www-form-urlencoded with fields `target` and
// `message`, and the raw token in the Authorization header (no "Bearer").
type FonnteNotifier struct {
	token  string
	apiURL string
	client *http.Client
}

func NewFonnte(token, apiURL string) domain.Notifier {
	return &FonnteNotifier{
		token:  token,
		apiURL: apiURL,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (f *FonnteNotifier) Send(ctx context.Context, to, message string) error {
	form := url.Values{}
	form.Set("target", to)
	form.Set("message", message)
	form.Set("countryCode", "62") // replaces a leading 0 with 62; no-op if already 62

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, f.apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", f.token) // Fonnte: token directly, no "Bearer"

	resp, err := f.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fonnte: http %d: %s", resp.StatusCode, string(body))
	}
	// Fonnte can reply 200 with {"status":false,...} when the device is
	// disconnected or the number is invalid; surface that for easier debugging.
	if strings.Contains(string(body), "\"status\":false") {
		return fmt.Errorf("fonnte: send rejected: %s", string(body))
	}
	return nil
}
