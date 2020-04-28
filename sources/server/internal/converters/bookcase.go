package converters

import (
	"fantlab/pb"
	"fantlab/server/internal/config"
	"fantlab/server/internal/db"
	"fantlab/server/internal/helpers"
	"math"
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
			Type:      helpers.BookcaseTypeMap[dbBookcase.BookcaseGroup],
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

func GetEditionBookcase(dbResponse db.EditionBookcaseDbResponse, dbBookcase db.Bookcase, page, limit uint64, cfg *config.AppConfig) *pb.Bookcase_EditionBookcaseResponse {
	bookcase := &pb.Bookcase_BookcaseInfo{
		Id:        dbBookcase.BookcaseId,
		IsPrivate: dbBookcase.BookcaseShared == 0,
		Type:      helpers.BookcaseTypeMap[dbBookcase.BookcaseGroup],
		Title:     dbBookcase.BookcaseName,
		Comment:   dbBookcase.BookcaseComment,
	}

	//noinspection GoPreferNilSlice
	var editions = []*pb.Bookcase_Edition{}

	for _, dbEdition := range dbResponse.Editions {
		cover := helpers.GetEditionCoverUrl(cfg.ImagesBaseURL, dbEdition.EditionId)

		var ozonOffer *pb.Bookcase_Edition_Offers_Offer
		if dbEdition.OzonAvailable == 1 && dbEdition.OzonId > 0 && dbEdition.OzonCost > 0 {
			ozonOffer = &pb.Bookcase_Edition_Offers_Offer{
				Url:   helpers.GetOzonOfferUrl(dbEdition.OzonId),
				Price: dbEdition.OzonCost,
			}
		}

		var labirintOffer *pb.Bookcase_Edition_Offers_Offer
		if dbEdition.LabirintAvailable == 1 && dbEdition.LabirintId > 0 && dbEdition.LabirintCost > 0 {
			labirintOffer = &pb.Bookcase_Edition_Offers_Offer{
				Url:   helpers.GetLabirintOfferUrl(dbEdition.LabirintId),
				Price: dbEdition.LabirintCost,
			}
		}

		var offers *pb.Bookcase_Edition_Offers
		if ozonOffer != nil || labirintOffer != nil {
			offers = &pb.Bookcase_Edition_Offers{
				Ozon:     ozonOffer,
				Labirint: labirintOffer,
			}
		}

		edition := &pb.Bookcase_Edition{
			ItemId:                 dbEdition.ItemId,
			Id:                     dbEdition.EditionId,
			Type:                   helpers.EditionTypeMap[dbEdition.Type],
			CorrectnessLevel:       helpers.EditionCorrectnessLevelMap[dbEdition.Correct],
			Cover:                  cover,
			Authors:                dbEdition.Autors,
			Title:                  dbEdition.Name,
			Year:                   dbEdition.Year,
			Publishers:             dbEdition.Publisher,
			Description:            dbEdition.Description,
			PlannedPublicationDate: dbEdition.PlanDate,
			Offers:                 offers,
			Comment:                dbEdition.Comment,
		}

		editions = append(editions, edition)
	}

	pageCount := helpers.CalculatePageCount(dbResponse.TotalCount, limit)

	return &pb.Bookcase_EditionBookcaseResponse{
		Bookcase: bookcase,
		Editions: editions,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}

func GetWorkBookcase(dbResponse db.WorkBookcaseDbResponse, dbBookcase db.Bookcase, page, limit uint64) *pb.Bookcase_WorkBookcaseResponse {
	bookcase := &pb.Bookcase_BookcaseInfo{
		Id:        dbBookcase.BookcaseId,
		IsPrivate: dbBookcase.BookcaseShared == 0,
		Type:      helpers.BookcaseTypeMap[dbBookcase.BookcaseGroup],
		Title:     dbBookcase.BookcaseName,
		Comment:   dbBookcase.BookcaseComment,
	}

	//noinspection GoPreferNilSlice
	var works = []*pb.Bookcase_Work{}

	for _, dbWork := range dbResponse.Works {
		stats := &pb.Bookcase_Work_Stats{
			AverageMark:   math.Round(dbWork.MidMark*100) / 100,
			MarkCount:     dbWork.MarkCount,
			ResponseCount: dbWork.ResponseCount,
		}
		own := &pb.Bookcase_Work_Own{
			Mark:                dbResponse.OwnWorkMarks[dbWork.WorkId],
			IsResponsePublished: dbResponse.OwnWorkResponses[dbWork.WorkId] == 1,
		}
		work := &pb.Bookcase_Work{
			ItemId:            dbWork.ItemId,
			Id:                dbWork.WorkId,
			Type:              helpers.WorkTypeMap[dbWork.WorkTypeId],
			Authors:           getWorkAuthors(dbWork, dbResponse.Autors),
			Title:             dbWork.RusName,
			OriginalTitle:     dbWork.Name,
			AlternativeTitles: dbWork.AltName,
			Note:              dbWork.BonusText,
			Year:              dbWork.Year,
			Description:       dbWork.Description,
			IsPublished:       dbWork.Published == 1,
			Stats:             stats,
			Own:               own,
			Comment:           dbWork.Comment,
		}

		works = append(works, work)
	}

	pageCount := helpers.CalculatePageCount(dbResponse.TotalCount, limit)

	return &pb.Bookcase_WorkBookcaseResponse{
		Bookcase: bookcase,
		Works:    works,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}

func getWorkAuthors(work db.Work, autors map[uint64]db.Autor) []*pb.Bookcase_Work_Author {
	author := autors[work.AutorId]
	authors := []*pb.Bookcase_Work_Author{
		{
			Id:       author.AutorId,
			Name:     author.RusName,
			IsOpened: author.IsOpened == 1,
		},
	}
	if work.Autor2Id != 0 {
		author := autors[work.Autor2Id]
		authors = append(authors, &pb.Bookcase_Work_Author{
			Id:       author.AutorId,
			Name:     author.RusName,
			IsOpened: author.IsOpened == 1,
		})
	}
	if work.Autor3Id != 0 {
		author := autors[work.Autor3Id]
		authors = append(authors, &pb.Bookcase_Work_Author{
			Id:       author.AutorId,
			Name:     author.RusName,
			IsOpened: author.IsOpened == 1,
		})
	}
	if work.Autor4Id != 0 {
		author := autors[work.Autor4Id]
		authors = append(authors, &pb.Bookcase_Work_Author{
			Id:       author.AutorId,
			Name:     author.RusName,
			IsOpened: author.IsOpened == 1,
		})
	}
	if work.Autor5Id != 0 {
		author := autors[work.Autor5Id]
		authors = append(authors, &pb.Bookcase_Work_Author{
			Id:       author.AutorId,
			Name:     author.RusName,
			IsOpened: author.IsOpened == 1,
		})
	}

	return authors
}

func GetFilmBookcase(dbResponse db.FilmBookcaseDbResponse, dbBookcase db.Bookcase, page, limit uint64, cfg *config.AppConfig) *pb.Bookcase_FilmBookcaseResponse {
	bookcase := &pb.Bookcase_BookcaseInfo{
		Id:        dbBookcase.BookcaseId,
		IsPrivate: dbBookcase.BookcaseShared == 0,
		Type:      helpers.BookcaseTypeMap[dbBookcase.BookcaseGroup],
		Title:     dbBookcase.BookcaseName,
		Comment:   dbBookcase.BookcaseComment,
	}

	//noinspection GoPreferNilSlice
	var films = []*pb.Bookcase_Film{}

	for _, dbFilm := range dbResponse.Films {
		poster := helpers.GetFilmPosterUrl(cfg.ImagesBaseURL, dbFilm.FilmId)

		var year uint64
		var startYear uint64
		var endYear uint64

		if helpers.FilmTypeMap[dbFilm.Type] == pb.FilmType_FILM_TYPE_SERIES {
			startYear = dbFilm.Year
			endYear = dbFilm.Year2
		} else {
			year = dbFilm.Year
		}

		film := &pb.Bookcase_Film{
			ItemId:        dbFilm.ItemId,
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
		Bookcase: bookcase,
		Films:    films,
		Pages: &pb.Common_Pages{
			Current: page,
			Count:   pageCount,
		},
	}
}
