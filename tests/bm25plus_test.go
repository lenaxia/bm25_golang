package bm25

import (
	"testing"
)

func TestNewBM25Plus(t *testing.T) {
	corpus := []string{"hello world", "goodbye world"}
	tokenizer := func(s string) []string {
		return strings.Split(s, " ")
	}
	k1 := 1.2
	b := 0.75
	delta := 0.5
	epsilon := 0.1
	logger := log.New(os.Stderr, "", log.LstdFlags)

	_, err := NewBM25Plus(corpus, tokenizer, k1, b, delta, epsilon, logger)
	if err != nil {
		t.Errorf("NewBM25Plus() failed: %v", err)
	}
}

func TestBM25Plus_GetScores(t *testing.T) {
	corpus := []string{"hello world", "goodbye world"}
	tokenizer := func(s string) []string {
		return strings.Split(s, " ")
	}
	k1 := 1.2
	b := 0.75
	delta := 0.5
	epsilon := 0.1
	logger := log.New(os.Stderr, "", log.LstdFlags)

	bm25, _ := NewBM25Plus(corpus, tokenizer, k1, b, delta, epsilon, logger)
	query := []string{"hello"}
	scores := bm25.GetScores(query)

	if len(scores) != 2 {
		t.Errorf("GetScores() returned incorrect number of scores: %d", len(scores))
	}
}

func TestBM25Plus_GetBatchScores(t *testing.T) {
	corpus := []string{"hello world", "goodbye world", "another doc"}
	tokenizer := func(s string) []string {
		return strings.Split(s, " ")
	}
	k1 := 1.2
	b := 0.75
	delta := 0.5
	epsilon := 0.1
	logger := log.New(os.Stderr, "", log.LstdFlags)

	bm25, _ := NewBM25Plus(corpus, tokenizer, k1, b, delta, epsilon, logger)
	query := []string{"hello"}
	docIDs := []int{0, 2}
	scores := bm25.GetBatchScores(query, docIDs)

	if len(scores) != 2 {
		t.Errorf("GetBatchScores() returned incorrect number of scores: %d", len(scores))
	}
}

func TestBM25Plus_GetTopN(t *testing.T) {
	corpus := []string{"hello world", "goodbye world", "another doc"}
	tokenizer := func(s string) []string {
		return strings.Split(s, " ")
	}
	k1 := 1.2
	b := 0.75
	delta := 0.5
	epsilon := 0.1
	logger := log.New(os.Stderr, "", log.LstdFlags)

	bm25, _ := NewBM25Plus(corpus, tokenizer, k1, b, delta, epsilon, logger)
	query := []string{"hello"}
	topDocs := bm25.GetTopN(query, 2)

	if len(topDocs) != 2 {
		t.Errorf("GetTopN() returned incorrect number of documents: %d", len(topDocs))
	}
}