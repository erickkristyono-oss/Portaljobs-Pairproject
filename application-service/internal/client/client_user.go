package client

import (
	"application-service/internal/dto/response"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UserClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewUserClient(baseURL string) *UserClient {
	return &UserClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *UserClient) GetUserByID(
	ctx context.Context,
	userID uint,
) error {

	url := fmt.Sprintf(
		"%s/users/%d",
		c.baseURL,
		userID,
	)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return err
	}

	authHeader, ok := GetAuthorizationToken(ctx)
	if ok && authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"user service returned status %d",
			resp.StatusCode,
		)
	}

	return nil
}

func (c *UserClient) GetUserProfile(
	ctx context.Context,
	userID uint,
) (*response.CandidateProfileResponse, error) {

	url := fmt.Sprintf(
		"%s/internal/users/%d/profile",
		c.baseURL,
		userID,
	)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	authHeader, ok := GetAuthorizationToken(ctx)
	if ok && authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result struct {
		ResponseCode    int                               `json:"responseCode"`
		ResponseMessage string                            `json:"responseMessage"`
		ResponseData    response.CandidateProfileResponse `json:"responseData"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf(
			"failed to decode user service response: %w",
			err,
		)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"user-service returned status %d: %s",
			resp.StatusCode,
			result.ResponseMessage,
		)
	}

	if result.ResponseCode != http.StatusOK {
		return nil, fmt.Errorf(
			"user-service error: %s",
			result.ResponseMessage,
		)
	}

	return &result.ResponseData, nil
}
