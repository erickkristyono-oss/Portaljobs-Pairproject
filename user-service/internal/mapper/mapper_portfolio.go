package mapper

import (
	"user-service/internal/domain/entity"
	"user-service/internal/dto/response"
	"user-service/internal/repository/postgres/model"
)

func PortfolioModelToEntity(portfolio *model.PortfolioModel) *entity.Portfolio {
	if portfolio == nil {
		return nil
	}

	return &entity.Portfolio{
		ID:           portfolio.ID,
		UserID:       portfolio.UserID,
		NameProject:  portfolio.NameProject,
		Organization: portfolio.Organization,
		Description:  portfolio.Description,
		CreatedAt:    portfolio.CreatedAt,
		UpdatedAt:    portfolio.UpdatedAt,
	}
}

func PortfolioEntityToModel(portfolio *entity.Portfolio) *model.PortfolioModel {
	if portfolio == nil {
		return nil
	}

	return &model.PortfolioModel{
		ID:           portfolio.ID,
		UserID:       portfolio.UserID,
		NameProject:  portfolio.NameProject,
		Organization: portfolio.Organization,
		Description:  portfolio.Description,
		CreatedAt:    portfolio.CreatedAt,
		UpdatedAt:    portfolio.UpdatedAt,
	}
}

func PortfolioEntityToResponse(portfolio *entity.Portfolio) *response.PortfolioResponse {
	if portfolio == nil {
		return nil
	}

	return &response.PortfolioResponse{
		ID:           portfolio.ID,
		UserID:       portfolio.UserID,
		NameProject:  portfolio.NameProject,
		Organization: portfolio.Organization,
		Description:  portfolio.Description,
		CreatedAt:    portfolio.CreatedAt,
		UpdatedAt:    portfolio.UpdatedAt,
	}
}
