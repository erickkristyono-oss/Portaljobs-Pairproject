package client

import (
	"application-service/internal/dto/response"
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

type jobResponse struct {
	ResponseCode    int    `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ResponseData    struct {
		ID        uint   `json:"id"`
		CompanyID uint   `json:"company_id"`
		Judul     string `json:"judul"`
		Status    string `json:"status"`
	} `json:"responseData"`
}

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

func (c *JobClient) GetJobByID(
	ctx context.Context,
	jobID uint,
) (uint, string, string, error) {

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
		return 0, "", "", err
	}

	authHeader, ok := GetAuthorizationToken(ctx)
	if ok && authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, "", "", err
	}
	defer resp.Body.Close()

	var result jobResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, "", "", fmt.Errorf(
			"failed to decode job service response: %w",
			err,
		)
	}

	if resp.StatusCode != http.StatusOK {
		return 0, "", "", fmt.Errorf(
			"job service returned status %d: %s",
			resp.StatusCode,
			result.ResponseMessage,
		)
	}

	if result.ResponseCode != http.StatusOK {
		return 0, "", "", fmt.Errorf(
			"job service error: %s",
			result.ResponseMessage,
		)
	}

	return result.ResponseData.CompanyID,
		result.ResponseData.Status,
		result.ResponseData.Judul,
		nil
}

func (c *JobClient) GetJobDetail(
	ctx context.Context,
	jobID uint,
) (*response.JobDetailResponse, error) {

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
		ResponseCode    int                        `json:"responseCode"`
		ResponseMessage string                     `json:"responseMessage"`
		ResponseData    response.JobDetailResponse `json:"responseData"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf(
			"failed to decode job service response: %w",
			err,
		)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"job service returned status %d: %s",
			resp.StatusCode,
			result.ResponseMessage,
		)
	}

	if result.ResponseCode != http.StatusOK {
		return nil, fmt.Errorf(
			"job service error: %s",
			result.ResponseMessage,
		)
	}

	return &result.ResponseData, nil
}

func (c *JobClient) GetJobsByCompanyID(
	ctx context.Context,
	companyID uint,
) ([]uint, error) {

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

	authHeader, ok := GetAuthorizationToken(ctx)
	if ok && authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result jobsResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf(
			"failed to decode job service response: %w",
			err,
		)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"job service returned status %d: %s",
			resp.StatusCode,
			result.ResponseMessage,
		)
	}

	if result.ResponseCode != http.StatusOK {
		return nil, fmt.Errorf(
			"job service error: %s",
			result.ResponseMessage,
		)
	}

	jobIDs := make([]uint, 0, len(result.ResponseData))

	for _, job := range result.ResponseData {
		jobIDs = append(jobIDs, job.ID)
	}

	return jobIDs, nil
}

func (c *JobClient) GetRequiredSkills(
	ctx context.Context,
	jobID uint,
) ([]response.JobRequiredSkillResponse, error) {

	url := fmt.Sprintf(
		"%s/internal/jobs/%d/required-skills",
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
		ResponseCode    int                                 `json:"responseCode"`
		ResponseMessage string                              `json:"responseMessage"`
		ResponseData    []response.JobRequiredSkillResponse `json:"responseData"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf(
			"failed to decode job service response: %w",
			err,
		)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"job service returned status %d: %s",
			resp.StatusCode,
			result.ResponseMessage,
		)
	}

	if result.ResponseCode != http.StatusOK {
		return nil, fmt.Errorf(
			"job service error: %s",
			result.ResponseMessage,
		)
	}

	return result.ResponseData, nil
}
