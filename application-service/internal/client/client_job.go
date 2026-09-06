package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type JobClient struct {
	baseURL    string
	httpClient *http.Client
}

// Response untuk GET /jobs/:id
type jobResponse struct {
	ResponseCode    int    `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ResponseData    struct {
		ID        uint   `json:"id"`
		CompanyID uint   `json:"company_id"`
		Status    string `json:"status"`
	} `json:"responseData"`
}

// Response untuk GET /internal/jobs/company/:company_id
type jobsResponse struct {
	ResponseCode    int    `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ResponseData    []struct {
		ID        uint `json:"id"`
		CompanyID uint `json:"company_id"`
	} `json:"responseData"`
}

func NewJobClient(baseURL string) *JobClient {
	return &JobClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// GetJobByID mengambil informasi job berdasarkan ID.
//
// Endpoint:
// GET /jobs/:id
//
// Return:
// - companyID
// - jobStatus
// - error
func (c *JobClient) GetJobByID(ctx context.Context, jobID uint) (uint, string, error) {

	url := fmt.Sprintf(
		"%s/jobs/%d",
		c.baseURL,
		jobID,
	)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return 0, "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, "", fmt.Errorf(
			"job service returned status %d",
			resp.StatusCode,
		)
	}

	var result jobResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, "", err
	}

	if result.ResponseCode != http.StatusOK {
		return 0, "", fmt.Errorf(
			"job service error: %s",
			result.ResponseMessage,
		)
	}

	return result.ResponseData.CompanyID,
		result.ResponseData.Status,
		nil
}

// GetJobsByCompanyID mengambil semua job milik company.
//
// Endpoint internal:
// GET /internal/jobs/company/:company_id
//
// Return:
// []uint berisi daftar JobID
func (c *JobClient) GetJobsByCompanyID(ctx context.Context, companyID uint) ([]uint, error) {

	url := fmt.Sprintf(
		"%s/internal/jobs/company/%d",
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

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"job service returned status %d",
			resp.StatusCode,
		)
	}

	var result jobsResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.ResponseCode != http.StatusOK {
		return nil, fmt.Errorf(
			"job service error: %s",
			result.ResponseMessage,
		)
	}

	jobIDs := make(
		[]uint,
		0,
		len(result.ResponseData),
	)

	for _, job := range result.ResponseData {
		jobIDs = append(
			jobIDs,
			job.ID,
		)
	}

	return jobIDs, nil
}
