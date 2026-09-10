package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type FonnteClient struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

type FonnteResponse struct {
	Status bool   `json:"status"`
	Reason string `json:"reason"`
}

func NewFonnteClient(
	baseURL string,
	token string,
) *FonnteClient {
	return &FonnteClient{
		baseURL: baseURL,
		token:   token,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *FonnteClient) SendMessage(
	ctx context.Context,
	target string,
	message string,
) error {

	form := url.Values{}

	form.Set("target", target)
	form.Set("message", message)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return err
	}

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	req.Header.Set(
		"Authorization",
		c.token,
	)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf(
			"fonnte returned status %d",
			resp.StatusCode,
		)
	}

	var result FonnteResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if !result.Status {
		if result.Reason != "" {
			return fmt.Errorf(
				"fonnte failed: %s",
				result.Reason,
			)
		}

		return fmt.Errorf(
			"fonnte failed to send message",
		)
	}

	return nil
}
