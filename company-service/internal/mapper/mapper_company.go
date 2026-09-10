package mapper

import (
	"company-service/internal/domain/entity"
	"company-service/internal/repository/postgres/model"
)

func EntityToCompanyModel(company *entity.Company) *model.CompanyModel {

	return &model.CompanyModel{
		ID:          company.ID,
		UserID:      company.UserID,
		Name:        company.Name,
		Phone:       company.Phone,
		FieldOf:     company.FieldOf,
		Address:     company.Address,
		Description: company.Description,
		CreatedAt:   company.CreatedAt,
		UpdatedAt:   company.UpdatedAt,
	}
}

func CompanyModelToEntity(company *model.CompanyModel) *entity.Company {

	return &entity.Company{
		ID:          company.ID,
		UserID:      company.UserID,
		Name:        company.Name,
		Phone:       company.Phone,
		FieldOf:     company.FieldOf,
		Address:     company.Address,
		Description: company.Description,
		CreatedAt:   company.CreatedAt,
		UpdatedAt:   company.UpdatedAt,
	}
}

func ToCompanyEntity(companyModel *model.CompanyModel) *entity.Company {

	return &entity.Company{
		ID:          companyModel.ID,
		UserID:      companyModel.UserID,
		Name:        companyModel.Name,
		FieldOf:     companyModel.FieldOf,
		Address:     companyModel.Address,
		Description: companyModel.Description,
	}
}
