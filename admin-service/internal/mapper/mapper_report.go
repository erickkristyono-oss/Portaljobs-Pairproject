package mapper

import (
	"admin-service/internal/domain/constant"
	"admin-service/internal/domain/entity"
	"admin-service/internal/dto/response"
	"admin-service/internal/repository/postgres/model"
)

func ToReportEntity(reportModel *model.ReportModel) *entity.Report {

	return &entity.Report{
		ID:         reportModel.ID,
		ReporterID: reportModel.ReporterID,
		TargetType: constant.ReportTargetType(reportModel.TargetType),
		TargetID:   reportModel.TargetID,
		Reason:     reportModel.Reason,
		Status:     constant.ReportStatus(reportModel.Status),
		CreatedAt:  reportModel.CreatedAt,
		UpdatedAt:  reportModel.UpdatedAt,
	}
}

func ToReportModel(report *entity.Report) *model.ReportModel {

	return &model.ReportModel{
		ID:         report.ID,
		ReporterID: report.ReporterID,
		TargetType: string(report.TargetType),
		TargetID:   report.TargetID,
		Reason:     report.Reason,
		Status:     string(report.Status),
		CreatedAt:  report.CreatedAt,
		UpdatedAt:  report.UpdatedAt,
	}
}

func ToReportResponse(report *entity.Report) response.ReportResponse {

	return response.ReportResponse{
		ID:         report.ID,
		ReporterID: report.ReporterID,
		TargetType: string(report.TargetType),
		TargetID:   report.TargetID,
		Reason:     report.Reason,
		Status:     string(report.Status),
		CreatedAt:  report.CreatedAt,
		UpdatedAt:  report.UpdatedAt,
	}
}
