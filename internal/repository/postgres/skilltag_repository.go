package postgres

import (
	"context"
	"errors"
	"strings"

	"portaljob/internal/domain"

	"gorm.io/gorm"
)

type skillTagRepository struct{ db *gorm.DB }

func NewSkillTagRepository(db *gorm.DB) domain.SkillTagRepository {
	return &skillTagRepository{db: db}
}

func (r *skillTagRepository) FindOrCreateByName(ctx context.Context, name string) (*domain.SkillTag, error) {
	n := strings.ToLower(strings.TrimSpace(name))
	var t domain.SkillTag
	err := r.db.WithContext(ctx).Where("name = ?", n).First(&t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t = domain.SkillTag{Name: n}
		if err := r.db.WithContext(ctx).Create(&t).Error; err != nil {
			return nil, err
		}
		return &t, nil
	}
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *skillTagRepository) List(ctx context.Context) ([]domain.SkillTag, error) {
	var tags []domain.SkillTag
	err := r.db.WithContext(ctx).Order("name ASC").Find(&tags).Error
	return tags, err
}

func (r *skillTagRepository) TagsByIDs(ctx context.Context, ids []uint) ([]domain.SkillTag, error) {
	if len(ids) == 0 {
		return []domain.SkillTag{}, nil
	}
	var tags []domain.SkillTag
	err := r.db.WithContext(ctx).Where("id IN ?", ids).Order("name ASC").Find(&tags).Error
	return tags, err
}

func (r *skillTagRepository) SetUserTags(ctx context.Context, userID uint, tagIDs []uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userID).Delete(&domain.UserSkillTag{}).Error; err != nil {
			return err
		}
		if len(tagIDs) == 0 {
			return nil
		}
		rows := make([]domain.UserSkillTag, 0, len(tagIDs))
		for _, id := range tagIDs {
			rows = append(rows, domain.UserSkillTag{UserID: userID, SkillTagID: id})
		}
		return tx.Create(&rows).Error
	})
}

func (r *skillTagRepository) UserTagIDs(ctx context.Context, userID uint) ([]uint, error) {
	var ids []uint
	err := r.db.WithContext(ctx).Model(&domain.UserSkillTag{}).
		Where("user_id = ?", userID).Pluck("skill_tag_id", &ids).Error
	return ids, err
}

func (r *skillTagRepository) SetJobTags(ctx context.Context, jobID uint, tagIDs []uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("job_id = ?", jobID).Delete(&domain.JobSkillTag{}).Error; err != nil {
			return err
		}
		if len(tagIDs) == 0 {
			return nil
		}
		rows := make([]domain.JobSkillTag, 0, len(tagIDs))
		for _, id := range tagIDs {
			rows = append(rows, domain.JobSkillTag{JobID: jobID, SkillTagID: id})
		}
		return tx.Create(&rows).Error
	})
}

func (r *skillTagRepository) JobTagIDs(ctx context.Context, jobID uint) ([]uint, error) {
	var ids []uint
	err := r.db.WithContext(ctx).Model(&domain.JobSkillTag{}).
		Where("job_id = ?", jobID).Pluck("skill_tag_id", &ids).Error
	return ids, err
}
