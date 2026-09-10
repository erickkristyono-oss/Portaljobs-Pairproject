package client

import (
	"application-service/internal/dto/response"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CompanyClient struct {
	baseURL    string
	httpClient *http.Client
}

type companyResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ResponseData    struct {
		ID     uint   `json:"id"`
		UserID uint   `json:"user_id"`
		Name   string `json:"name"`
		Phone  string `json:"phone"`
	} `json:"responseData"`
}

func NewCompanyClient(baseURL string) *CompanyClient {
	return &CompanyClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *CompanyClient) GetCompanyByUserID(
	ctx context.Context,
	userID uint,
) (uint, error) {

	url := fmt.Sprintf(
		"%s/internal/companies/user/%d",
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
		return 0, err
	}

	authHeader, ok := GetAuthorizationToken(ctx)
	if ok && authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result companyResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf(
			"failed to decode company service response: %w",
			err,
		)
	}

	// Gunakan HTTP status sebagai penentu berhasil/gagal
	if resp.StatusCode < http.StatusOK ||
		resp.StatusCode >= http.StatusMultipleChoices {

		return 0, fmt.Errorf(
			"company service returned status %d: %s",
			resp.StatusCode,
			result.ResponseMessage,
		)
	}

	return result.ResponseData.ID, nil
}

func (c *CompanyClient) GetCompanyByID(
	ctx context.Context,
	companyID uint,
) (*response.CompanyDetailResponse, error) {

	url := fmt.Sprintf(
		"%s/companies/%d",
		c.baseURL,
		companyID,
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
		ResponseCode    string                         `json:"responseCode"`
		ResponseMessage string                         `json:"responseMessage"`
		ResponseData    response.CompanyDetailResponse `json:"responseData"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf(
			"failed to decode company service response: %w",
			err,
		)
	}

	// Gunakan HTTP status sebagai penentu berhasil/gagal
	if resp.StatusCode < http.StatusOK ||
		resp.StatusCode >= http.StatusMultipleChoices {

		return nil, fmt.Errorf(
			"company service returned status %d: %s",
			resp.StatusCode,
			result.ResponseMessage,
		)
	}

	return &result.ResponseData, nil
}
