package usecase

import "testing"

func TestComputeMatch_Partial(t *testing.T) {
	matched, missing := computeMatch([]uint{1, 2, 3, 4}, []uint{2, 4})
	if len(matched) != 2 || len(missing) != 2 {
		t.Fatalf("expected 2 matched / 2 missing, got %d / %d", len(matched), len(missing))
	}
}

func TestComputeMatch_Full(t *testing.T) {
	matched, missing := computeMatch([]uint{1, 2}, []uint{1, 2, 3})
	if len(matched) != 2 || len(missing) != 0 {
		t.Fatalf("expected full match, got %d matched / %d missing", len(matched), len(missing))
	}
}

func TestComputeMatch_None(t *testing.T) {
	matched, missing := computeMatch([]uint{1, 2}, []uint{9})
	if len(matched) != 0 || len(missing) != 2 {
		t.Fatalf("expected no match, got %d matched / %d missing", len(matched), len(missing))
	}
}

func TestComputeMatch_NoRequirement(t *testing.T) {
	matched, missing := computeMatch(nil, []uint{1, 2})
	if len(matched) != 0 || len(missing) != 0 {
		t.Fatalf("expected empty when no requirement, got %d / %d", len(matched), len(missing))
	}
}
