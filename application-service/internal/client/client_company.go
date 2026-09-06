package client

import (
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
	ResponseCode    int    `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ResponseData    struct {
		ID     uint   `json:"id"`
		UserID uint   `json:"user_id"`
		Name   string `json:"name"`
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

func (c *CompanyClient) GetCompanyByUserID(ctx context.Context, userID uint) (uint, error) {

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

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(
			"company service returned status %d",
			resp.StatusCode,
		)
	}

	var result companyResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	if result.ResponseCode != http.StatusOK {
		return 0, fmt.Errorf(
			"company service error: %s",
			result.ResponseMessage,
		)
	}

	return result.ResponseData.ID, nil
}
