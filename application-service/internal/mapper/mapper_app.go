package mapper

import (
	"application-service/internal/domain/entity"
	"application-service/internal/dto/response"
	"application-service/internal/repository/postgres/model"
)

func ToApplicationEntity(applicationModel *model.ApplicationModel) *entity.Application {

	if applicationModel == nil {
		return nil
	}

	return &entity.Application{
		ID:        applicationModel.ID,
		UserID:    applicationModel.UserID,
		JobID:     applicationModel.JobID,
		Status:    applicationModel.Status,
		AppliedAt: applicationModel.AppliedAt,
	}
}

func ToApplicationModel(application *entity.Application) *model.ApplicationModel {

	if application == nil {
		return nil
	}

	return &model.ApplicationModel{
		ID:        application.ID,
		UserID:    application.UserID,
		JobID:     application.JobID,
		Status:    application.Status,
		AppliedAt: application.AppliedAt,
	}
}

func ToApplicationResponse(application *entity.Application) response.ApplicationResponse {

	return response.FromEntity(application)
}
