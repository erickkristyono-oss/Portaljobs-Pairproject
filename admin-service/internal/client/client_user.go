package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserClient interface {
	UpdateStatus(
		ctx context.Context,
		id uint,
		status string,
	) error
}

type userClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewUserClient(baseURL string) UserClient {
	return &userClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *userClient) UpdateStatus(
	ctx context.Context,
	id uint,
	status string,
) error {

	payload := map[string]string{
		"status": status,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(
		"%s/internal/users/%d/status",
		c.baseURL,
		id,
	)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPatch,
		url,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf(
			"user-service returned status code %d",
			resp.StatusCode,
		)
	}

	return nil
}
