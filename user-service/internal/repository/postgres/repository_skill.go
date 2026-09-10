package postgres

import (
	"context"

	"user-service/internal/domain/entity"
	domainrepo "user-service/internal/domain/repository"
	"user-service/internal/mapper"
	"user-service/internal/repository/postgres/model"

	"gorm.io/gorm"
)

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) domainrepo.SkillRepository {
	return &skillRepository{
		db: db,
	}
}

func (r *skillRepository) Create(
	ctx context.Context,
	skill *entity.Skill,
) error {
	skillModel := mapper.SkillEntityToModel(skill)

	if err := r.db.WithContext(ctx).
		Create(skillModel).Error; err != nil {
		return err
	}

	skill.ID = skillModel.ID
	skill.CreatedAt = skillModel.CreatedAt

	return nil
}

func (r *skillRepository) FindByUserID(ctx context.Context, userID uint) ([]entity.Skill, error) {
	var skillModels []model.SkillModel

	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("id ASC").
		Find(&skillModels).Error; err != nil {
		return nil, err
	}

	skills := make([]entity.Skill, 0, len(skillModels))

	for i := range skillModels {
		skill := mapper.SkillModelToEntity(&skillModels[i])

		if skill != nil {
			skills = append(skills, *skill)
		}
	}

	return skills, nil
}

func (r *skillRepository) FindByID(ctx context.Context, id uint) (*entity.Skill, error) {
	var skillModel model.SkillModel

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&skillModel).Error

	if err != nil {
		return nil, err
	}

	return mapper.SkillModelToEntity(&skillModel), nil
}

func (r *skillRepository) Update(
	ctx context.Context,
	skill *entity.Skill,
) error {

	result := r.db.WithContext(ctx).
		Model(&model.SkillModel{}).
		Where("id = ?", skill.ID).
		Updates(map[string]interface{}{
			"name_license": skill.NameLicense,
			"skill_tag":    skill.SkillTag,
			"level":        skill.Level,
			"organization": skill.Organization,
			"grade":        skill.Grade,
			"expired_date": skill.ExpiredDate,
			"description":  skill.Description,
			"updated_at":   skill.UpdatedAt,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *skillRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).
		Delete(&model.SkillModel{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
