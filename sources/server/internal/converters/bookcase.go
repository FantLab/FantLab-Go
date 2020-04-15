package converters

import (
	"fantlab/pb"
	"fantlab/server/internal/db"
	"fantlab/server/internal/helpers"
)

func getBookcaseGroupTitle(group string) string {
	switch group {
	case "work":
		return "Подборки произведений"
	case "edition":
		return "Книжные полки изданий"
	case "film":
		return "Кинополки"
	default:
		return ""
	}
}

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
				Title:     getBookcaseGroupTitle(group),
				Bookcases: bookcases,
			})
			bookcases = []*pb.Bookcase_Bookcase{}
			group = dbBookcases[index+1].BookcaseType
		} else if index == len(dbBookcases)-1 {
			bookcaseBlocks = append(bookcaseBlocks, &pb.Bookcase_BookcaseBlock{
				Title:     getBookcaseGroupTitle(group),
				Bookcases: bookcases,
			})
		}
	}

	return &pb.Bookcase_BookcaseBlocksResponse{
		BookcaseBlocks: bookcaseBlocks,
	}
}
