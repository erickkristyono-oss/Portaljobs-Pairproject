package usecase

import (
	"context"

	"portaljob/internal/domain"
)

type MatchUsecase struct{ tags domain.SkillTagRepository }

func NewMatchUsecase(tags domain.SkillTagRepository) *MatchUsecase {
	return &MatchUsecase{tags: tags}
}

// Match compares a jobseeker's tags against a job's required tags.
func (u *MatchUsecase) Match(ctx context.Context, userID, jobID uint) (*domain.MatchResult, error) {
	required, err := u.tags.JobTagIDs(ctx, jobID)
	if err != nil {
		return nil, err
	}
	owned, err := u.tags.UserTagIDs(ctx, userID)
	if err != nil {
		return nil, err
	}

	matchedIDs, missingIDs := computeMatch(required, owned)

	res := &domain.MatchResult{JobID: jobID, HasRequirement: len(required) > 0}
	if res.HasRequirement {
		res.Score = int(float64(len(matchedIDs))/float64(len(required))*100 + 0.5)
	} else {
		res.Score = 100 // no requirement declared -> nothing to fail
	}
	res.Warning = res.HasRequirement && res.Score < domain.MatchWarnThreshold

	matched, err := u.tags.TagsByIDs(ctx, matchedIDs)
	if err != nil {
		return nil, err
	}
	missing, err := u.tags.TagsByIDs(ctx, missingIDs)
	if err != nil {
		return nil, err
	}
	res.Matched = tagNames(matched)
	res.Missing = tagNames(missing)
	return res, nil
}

// computeMatch is a pure function: split required into matched/missing given
// what the user owns. Easy to unit-test with no DB.
func computeMatch(required, owned []uint) (matched, missing []uint) {
	ownedSet := make(map[uint]bool, len(owned))
	for _, id := range owned {
		ownedSet[id] = true
	}
	for _, id := range required {
		if ownedSet[id] {
			matched = append(matched, id)
		} else {
			missing = append(missing, id)
		}
	}
	return matched, missing
}

func tagNames(tags []domain.SkillTag) []string {
	names := make([]string, 0, len(tags))
	for _, t := range tags {
		names = append(names, t.Name)
	}
	return names
}
