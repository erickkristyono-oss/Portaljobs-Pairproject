package mapper

import (
	"company-service/internal/domain/entity"
	"company-service/internal/dto/response"
	"company-service/internal/repository/postgres/model"
)

func CompanyModelToEntity(companyModel *model.CompanyModel) *entity.Company {
	return &entity.Company{
		ID:          companyModel.ID,
		UserID:      companyModel.UserID,
		Name:        companyModel.Name,
		FieldOf:     companyModel.FieldOf,
		Address:     companyModel.Address,
		Description: companyModel.Description,
		CreatedAt:   companyModel.CreatedAt,
		UpdatedAt:   companyModel.UpdatedAt,
	}
}

func CompanyEntityToModel(company *entity.Company) *model.CompanyModel {
	return &model.CompanyModel{
		ID:          company.ID,
		UserID:      company.UserID,
		Name:        company.Name,
		FieldOf:     company.FieldOf,
		Address:     company.Address,
		Description: company.Description,
		CreatedAt:   company.CreatedAt,
		UpdatedAt:   company.UpdatedAt,
	}
}

func CompanyEntityToResponse(company *entity.Company) *response.CompanyResponse {
	return &response.CompanyResponse{
		ID:          company.ID,
		UserID:      company.UserID,
		Name:        company.Name,
		FieldOf:     company.FieldOf,
		Address:     company.Address,
		Description: company.Description,
		CreatedAt:   company.CreatedAt,
		UpdatedAt:   company.UpdatedAt,
	}
}
