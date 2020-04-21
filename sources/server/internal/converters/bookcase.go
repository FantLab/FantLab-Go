package converters

import (
	"fantlab/pb"
	"fantlab/server/internal/config"
	"fantlab/server/internal/db"
	"fantlab/server/internal/helpers"
)

var (
	BookcaseGroupTitleMap = map[string]string{
		db.BookcaseWorkType:    "Подборки произведений",
		db.BookcaseEditionType: "Книжные полки изданий",
		db.BookcaseFilmType:    "Кинополки",
	}
)

func GetBookcases(dbBookcases []db.Bookcase) *pb.Bookcase_BookcaseBlocksResponse {
	var bookcaseBlocks []*pb.Bookcase_BookcaseBlock
	var bookcases []*pb.Bookcase_Bookcase

	group := ""
	if len(dbBookcases) > 0 {
		group = dbBookcases[0].BookcaseType
	}

	for index, dbBookcase := range dbBookcases {
		bookcase := &pb.Bookcase_Bookcase{
			Id:        dbBookcase.BookcaseId,
			IsPrivate: dbBookcase.BookcaseShared == 0,
			Type:      helpers.GetBookcaseType(dbBookcase.BookcaseGroup),
			Title:     dbBookcase.BookcaseName,
			Comment:   dbBookcase.BookcaseComment,
			Index:     dbBookcase.Sort,
			ItemCount: dbBookcase.ItemCount,
		}

		bookcases = append(bookcases, bookcase)

		if index < len(dbBookcases)-1 && dbBookcases[index+1].BookcaseType != group {
			bookcaseBlocks = append(bookcaseBlocks, &pb.Bookcase_BookcaseBlock{
				Title:     BookcaseGroupTitleMap[group],
				Bookcases: bookcases,
			})
			bookcases = []*pb.Bookcase_Bookcase{}
			group = dbBookcases[index+1].BookcaseType
		} else if index == len(dbBookcases)-1 {
			bookcaseBlocks = append(bookcaseBlocks, &pb.Bookcase_BookcaseBlock{
				Title:     BookcaseGroupTitleMap[group],
				Bookcases: bookcases,
			})
		}
	}

	return &pb.Bookcase_BookcaseBlocksResponse{
		BookcaseBlocks: bookcaseBlocks,
	}
}

func GetFilmBookcase(dbResponse db.FilmBookcaseDbResponse, page, limit uint64, cfg *config.AppConfig) *pb.Bookcase_FilmBookcaseResponse {
	//noinspection GoPreferNilSlice
	var films = []*pb.Bookcase_Film{}

	for _, dbFilm := range dbResponse.Films {
		poster := helpers.GetFilmPosterUrl(cfg.ImagesBaseURL, dbFilm.FilmId)

		var year uint64
		var startYear uint64
		var endYear uint64

		if helpers.FilmTypeMap[dbFilm.Type] == pb.Bookcase_FILM_TYPE_SERIES {
			startYear = dbFilm.Year
			endYear = dbFilm.Year2
		} else {
			year = dbFilm.Year
		}

		film := &pb.Bookcase_Film{
			Id:            dbFilm.FilmId,
			Type:          helpers.FilmTypeMap[dbFilm.Type],
			Poster:        poster,
			Title:         dbFilm.RusName,
			OriginalTitle: dbFilm.Name,
			Year:          year,
			StartYear:     startYear,
			EndYear:       endYear,
			Countries:     dbFilm.Country,
			Genres:        dbFilm.Genre,
			Directors:     dbFilm.Director,
			ScreenWriters: dbFilm.ScreenWriter,
			Actors:        dbFilm.Actors,
			Description:   dbFilm.Description,
			Comment:       dbFilm.Comment,
		}

		films = append(films, film)
	}

	pageCount := helpers.CalculatePageCount(dbResponse.TotalCount, limit)

	return &pb.Bookcase_FilmBookcaseResponse{
		Films: films,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}
