package notification

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"portaljob/internal/domain"
)

// StubNotifier just logs. Used by default and in local/dev so the app runs
// without any real WA credentials.
type StubNotifier struct{}

func NewStub() domain.Notifier { return &StubNotifier{} }

func (s *StubNotifier) Send(ctx context.Context, to, message string) error {
	log.Printf("[WA stub] to=%s message=%q", to, message)
	return nil
}

// FonnteNotifier sends a WhatsApp message via the Fonnte gateway (an example
// Indonesian provider). Swap this out for Twilio / Meta Cloud API later — the
// rest of the app only depends on domain.Notifier, so nothing else changes.
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
	body, _ := json.Marshal(map[string]string{"target": to, "message": message})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, f.apiURL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", f.token)

	resp, err := f.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
