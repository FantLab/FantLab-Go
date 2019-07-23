package genresapi

import (
	"fantlab/db"
	"fantlab/pb"
)

func getGenres(dbResponse *db.WorkGenresDBResponse) *pb.Genre_Response {
	genresTable := make(map[uint16]*pb.Genre_Genre, len(dbResponse.Genres))

	for _, dbGenre := range dbResponse.Genres {
		genresTable[dbGenre.Id] = &pb.Genre_Genre{
			Id:        uint32(dbGenre.Id),
			Name:      dbGenre.Name,
			Info:      dbGenre.Info,
			WorkCount: dbGenre.WorkCount,
		}
	}

	groupsMap := make(map[uint16]*pb.Genre_Genre)

	for _, dbGenre := range dbResponse.Genres {
		root := groupsMap[dbGenre.GroupId]

		if root == nil {
			root = &pb.Genre_Genre{}
		}

		genre := genresTable[dbGenre.Id]
		parentGenre := genresTable[dbGenre.ParentId]

		if parentGenre != nil {
			parentGenre.Subgenres = append(parentGenre.Subgenres, genre)
		} else {
			root.Subgenres = append(root.Subgenres, genre)
		}

		groupsMap[dbGenre.GroupId] = root
	}

	genreGroups := make([]*pb.Genre_Group, len(dbResponse.GenreGroups))

	for index, dbGroup := range dbResponse.GenreGroups {
		rootGenre := groupsMap[dbGroup.Id]

		if rootGenre == nil {
			continue
		}

		group := &pb.Genre_Group{
			Id:     uint32(dbGroup.Id),
			Name:   dbGroup.Name,
			Genres: rootGenre.Subgenres,
		}

		genreGroups[index] = group
	}

	return &pb.Genre_Response{
		Groups: genreGroups,
	}
}
