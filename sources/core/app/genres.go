package app

import (
	"context"
	"fantlab/core/converters"
	"time"
)

const genreTreeCacheKey = "go:genres"

func (s *Services) GetGenreTree(ctx context.Context) *converters.GenreTree {
	value, _ := s.localStorage.Get(genreTreeCacheKey)

	if tree, ok := value.(*converters.GenreTree); ok {
		return tree
	}

	response, _ := s.db.FetchGenres(ctx)

	if response == nil {
		return nil
	}

	tree := converters.MakeGenreTree(response)

	s.localStorage.Set(genreTreeCacheKey, tree, time.Now().AddDate(0, 1, 0), true)

	return tree
}
