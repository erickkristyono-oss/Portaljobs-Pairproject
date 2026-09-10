package postgres

import (
	"context"

	"job-service/internal/domain/entity"
	domainrepo "job-service/internal/domain/repository"
	"job-service/internal/mapper"
	"job-service/internal/repository/postgres/model"

	"gorm.io/gorm"
)

type jobRequiredSkillRepository struct {
	db *gorm.DB
}

func NewJobRequiredSkillRepository(db *gorm.DB) domainrepo.JobRequiredSkillRepository {
	return &jobRequiredSkillRepository{
		db: db,
	}
}

func (r *jobRequiredSkillRepository) Create(ctx context.Context, skill *entity.JobRequiredSkill) error {
	skillModel := mapper.ToJobRequiredSkillModel(skill)

	if err := r.db.WithContext(ctx).
		Create(skillModel).Error; err != nil {
		return err
	}

	skill.ID = skillModel.ID
	skill.CreatedAt = skillModel.CreatedAt

	return nil
}

func (r *jobRequiredSkillRepository) CreateMany(ctx context.Context, skills []entity.JobRequiredSkill) error {
	if len(skills) == 0 {
		return nil
	}

	skillModels := make([]model.JobRequiredSkillModel, 0, len(skills))

	for i := range skills {
		skillModel := mapper.ToJobRequiredSkillModel(&skills[i])

		skillModels = append(skillModels, *skillModel)
	}

	if err := r.db.WithContext(ctx).
		Create(&skillModels).Error; err != nil {
		return err
	}

	for i := range skillModels {
		skills[i].ID = skillModels[i].ID
		skills[i].CreatedAt = skillModels[i].CreatedAt
	}

	return nil
}

func (r *jobRequiredSkillRepository) FindByJobID(ctx context.Context, jobID uint) ([]entity.JobRequiredSkill, error) {
	var skillModels []model.JobRequiredSkillModel

	err := r.db.WithContext(ctx).
		Where("job_id = ?", jobID).
		Order("id ASC").
		Find(&skillModels).Error

	if err != nil {
		return nil, err
	}

	skills := make([]entity.JobRequiredSkill, 0, len(skillModels))

	for i := range skillModels {
		skill := mapper.ToJobRequiredSkillEntity(&skillModels[i])

		if skill != nil {
			skills = append(skills, *skill)
		}
	}

	return skills, nil
}

func (r *jobRequiredSkillRepository) DeleteByJobID(ctx context.Context, jobID uint) error {
	result := r.db.WithContext(ctx).
		Where("job_id = ?", jobID).
		Delete(&model.JobRequiredSkillModel{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
