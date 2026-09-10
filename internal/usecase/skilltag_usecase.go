package usecase

import (
	"context"
	"strings"

	"portaljob/internal/domain"
)

type SkillTagUsecase struct{ tags domain.SkillTagRepository }

func NewSkillTagUsecase(tags domain.SkillTagRepository) *SkillTagUsecase {
	return &SkillTagUsecase{tags: tags}
}

func (u *SkillTagUsecase) ListVocab(ctx context.Context) ([]domain.SkillTag, error) {
	return u.tags.List(ctx)
}

// SetUserTags replaces a jobseeker's skill tags from a list of names.
func (u *SkillTagUsecase) SetUserTags(ctx context.Context, userID uint, names []string) ([]domain.SkillTag, error) {
	ids, err := resolveTagIDs(ctx, u.tags, names)
	if err != nil {
		return nil, err
	}
	if err := u.tags.SetUserTags(ctx, userID, ids); err != nil {
		return nil, err
	}
	return u.tags.TagsByIDs(ctx, ids)
}

func (u *SkillTagUsecase) GetUserTags(ctx context.Context, userID uint) ([]domain.SkillTag, error) {
	ids, err := u.tags.UserTagIDs(ctx, userID)
	if err != nil {
		return nil, err
	}
	return u.tags.TagsByIDs(ctx, ids)
}

// resolveTagIDs turns tag names into IDs (creating tags on first use), deduped.
func resolveTagIDs(ctx context.Context, repo domain.SkillTagRepository, names []string) ([]uint, error) {
	seen := map[uint]bool{}
	ids := make([]uint, 0, len(names))
	for _, name := range names {
		if strings.TrimSpace(name) == "" {
			continue
		}
		tag, err := repo.FindOrCreateByName(ctx, name)
		if err != nil {
			return nil, err
		}
		if !seen[tag.ID] {
			seen[tag.ID] = true
			ids = append(ids, tag.ID)
		}
	}
	return ids, nil
}
