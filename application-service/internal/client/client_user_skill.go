package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"application-service/internal/dto/response"
)

func (c *UserClient) GetUserSkills(ctx context.Context, userID uint) ([]response.SkillResponse, error) {

	url := fmt.Sprintf(
		"%s/internal/users/%d/skills",
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

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"user service returned status %d",
			resp.StatusCode,
		)
	}

	var result struct {
		ResponseCode    int                      `json:"responseCode"`
		ResponseMessage string                   `json:"responseMessage"`
		ResponseData    []response.SkillResponse `json:"responseData"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.ResponseData, nil
}
