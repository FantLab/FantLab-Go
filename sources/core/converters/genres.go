package converters

import (
	"fantlab/core/db"
	"fantlab/pb"
)

func convertGenreDBModelToAPIModel(dbResponse *db.WorkGenresDBResponse, enrichGenre func(genre *pb.Genre_Genre), isValidGenre func(genre *pb.Genre_Genre) bool) []*pb.Genre_Group {
	genresTable := make(map[uint64]*pb.Genre_Genre)

	for _, dbGenre := range dbResponse.Genres {
		genre := &pb.Genre_Genre{
			Id:   dbGenre.Id,
			Name: dbGenre.Name,
			Info: dbGenre.Info,
		}
		if enrichGenre != nil {
			enrichGenre(genre)
		}
		genresTable[dbGenre.Id] = genre
	}

	groupsTable := make(map[uint64]*pb.Genre_Genre)

	for _, dbGenre := range dbResponse.Genres {
		genre := genresTable[dbGenre.Id]

		if isValidGenre != nil && !isValidGenre(genre) {
			continue
		}

		root := groupsTable[dbGenre.GroupId]

		if root == nil {
			root = &pb.Genre_Genre{}
		}

		parentGenre := genresTable[dbGenre.ParentId]

		if parentGenre != nil {
			parentGenre.Subgenres = append(parentGenre.Subgenres, genre)
		} else {
			root.Subgenres = append(root.Subgenres, genre)
		}

		groupsTable[dbGenre.GroupId] = root
	}

	var genreGroups []*pb.Genre_Group

	for _, dbGroup := range dbResponse.GenreGroups {
		rootGenre := groupsTable[dbGroup.Id]

		if rootGenre == nil {
			continue
		}

		group := &pb.Genre_Group{
			Id:     dbGroup.Id,
			Name:   dbGroup.Name,
			Genres: rootGenre.Subgenres,
		}

		genreGroups = append(genreGroups, group)
	}

	return genreGroups
}

func GetGenres(genreTree *GenreTree, selectedGenreIds []uint64, workCounts map[uint64]uint64) *pb.Genre_GenresResponse {
	var validGenreIds map[uint64]bool
	if len(selectedGenreIds) > 0 {
		validGenreIds = genreTree.SelectGenreIdsWithParents(selectedGenreIds)
	}

	genreGroups := convertGenreDBModelToAPIModel(genreTree.Data(), func(genre *pb.Genre_Genre) {
		if workCounts != nil {
			genre.WorkCount = workCounts[genre.Id]
		}
	}, func(genre *pb.Genre_Genre) bool {
		if validGenreIds == nil {
			return true
		}
		return validGenreIds[genre.Id]
	})

	return &pb.Genre_GenresResponse{
		Groups: genreGroups,
	}
}

func GetWorkClassification(genreTree *GenreTree, classificationCount uint64, genreVotes map[uint64]uint64) *pb.Genre_ClassificationResponse {
	var selectedGenreIds []uint64
	for id := range genreVotes {
		selectedGenreIds = append(selectedGenreIds, id)
	}

	validGenreIds := genreTree.SelectGenreIdsWithParents(selectedGenreIds)

	genreGroups := convertGenreDBModelToAPIModel(genreTree.Data(), func(genre *pb.Genre_Genre) {
		genre.VoteCount = genreVotes[genre.Id]
	}, func(genre *pb.Genre_Genre) bool {
		return validGenreIds[genre.Id]
	})

	return &pb.Genre_ClassificationResponse{
		Groups:              genreGroups,
		ClassificationCount: classificationCount,
	}
}
