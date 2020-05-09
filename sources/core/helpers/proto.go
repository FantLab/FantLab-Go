package helpers

import "fantlab/pb"

var (
	UserClassMap = map[uint8]pb.Common_UserClass{
		0: pb.Common_USERCLASS_BEGINNER,
		1: pb.Common_USERCLASS_ACTIVIST,
		2: pb.Common_USERCLASS_AUTHORITY,
		3: pb.Common_USERCLASS_PHILOSOPHER,
		4: pb.Common_USERCLASS_MASTER,
		5: pb.Common_USERCLASS_GRANDMASTER,
		6: pb.Common_USERCLASS_PEACEKEEPER,
		7: pb.Common_USERCLASS_PEACEMAKER,
	}

	UserClassDescriptionMap = map[pb.Common_UserClass]string{
		pb.Common_USERCLASS_BEGINNER:    "новичок",
		pb.Common_USERCLASS_ACTIVIST:    "активист",
		pb.Common_USERCLASS_AUTHORITY:   "авторитет",
		pb.Common_USERCLASS_PHILOSOPHER: "философ",
		pb.Common_USERCLASS_MASTER:      "магистр",
		pb.Common_USERCLASS_GRANDMASTER: "гранд-мастер",
		pb.Common_USERCLASS_PEACEKEEPER: "миродержец",
		pb.Common_USERCLASS_PEACEMAKER:  "миротворец",
	}

	EditionTypeMap = map[uint64]pb.EditionType{
		10: pb.EditionType_EDITION_TYPE_AUTHOR_BOOK,
		11: pb.EditionType_EDITION_TYPE_AUTHOR_COMPILATION,
		12: pb.EditionType_EDITION_TYPE_COMPILATION,
		15: pb.EditionType_EDITION_TYPE_ANTHOLOGY,
		16: pb.EditionType_EDITION_TYPE_CHRESTOMATHY,
		20: pb.EditionType_EDITION_TYPE_MAGAZINE,
		21: pb.EditionType_EDITION_TYPE_FANZINE,
		22: pb.EditionType_EDITION_TYPE_ALMANAC,
		25: pb.EditionType_EDITION_TYPE_NEWSPAPER,
		30: pb.EditionType_EDITION_TYPE_AUDIOBOOK,
		32: pb.EditionType_EDITION_TYPE_ILLUSTRATED_ALBUM,
		34: pb.EditionType_EDITION_TYPE_FILM_STRIP,
	}

	EditionCorrectnessLevelMap = map[uint64]pb.EditionCorrectnessLevel{
		0: pb.EditionCorrectnessLevel_EDITION_CORRECTNESS_LEVEL_RED,
		1: pb.EditionCorrectnessLevel_EDITION_CORRECTNESS_LEVEL_ORANGE,
		2: pb.EditionCorrectnessLevel_EDITION_CORRECTNESS_LEVEL_GREEN,
	}

	WorkTypeMap = map[uint64]pb.WorkType{
		1:  pb.WorkType_WORK_TYPE_NOVEL,
		3:  pb.WorkType_WORK_TYPE_COMPILATION,
		4:  pb.WorkType_WORK_TYPE_SERIES,
		5:  pb.WorkType_WORK_TYPE_VERSE,
		8:  pb.WorkType_WORK_TYPE_FAIRY_TALE,
		11: pb.WorkType_WORK_TYPE_ESSAY,
		12: pb.WorkType_WORK_TYPE_ARTICLE,
		13: pb.WorkType_WORK_TYPE_EPIC_NOVEL,
		17: pb.WorkType_WORK_TYPE_ANTHOLOGY,
		18: pb.WorkType_WORK_TYPE_PLAY,
		19: pb.WorkType_WORK_TYPE_SCREENPLAY,
		20: pb.WorkType_WORK_TYPE_DOCUMENTARY,
		21: pb.WorkType_WORK_TYPE_MICROTALE,
		22: pb.WorkType_WORK_TYPE_DISSERTATION,
		23: pb.WorkType_WORK_TYPE_MONOGRAPH,
		24: pb.WorkType_WORK_TYPE_EDUCATIONAL_PUBLICATION,
		25: pb.WorkType_WORK_TYPE_ENCYCLOPEDIA,
		26: pb.WorkType_WORK_TYPE_MAGAZINE,
		27: pb.WorkType_WORK_TYPE_POEM,
		28: pb.WorkType_WORK_TYPE_POETRY,
		29: pb.WorkType_WORK_TYPE_PROSE_VERSE,
		41: pb.WorkType_WORK_TYPE_COMIC_BOOK,
		42: pb.WorkType_WORK_TYPE_MANGA,
		43: pb.WorkType_WORK_TYPE_GRAPHIC_NOVEL,
		44: pb.WorkType_WORK_TYPE_NOVELETTE,
		45: pb.WorkType_WORK_TYPE_STORY,
		46: pb.WorkType_WORK_TYPE_FEATURE_ARTICLE,
		47: pb.WorkType_WORK_TYPE_REPORTAGE,
		48: pb.WorkType_WORK_TYPE_CONDITIONAL_SERIES,
		49: pb.WorkType_WORK_TYPE_EXCERPT,
		51: pb.WorkType_WORK_TYPE_INTERVIEW,
		52: pb.WorkType_WORK_TYPE_REVIEW,
		53: pb.WorkType_WORK_TYPE_POPULAR_SCIENCE_BOOK,
	}

	FilmTypeMap = map[uint64]pb.FilmType{
		// В таблице фильмов есть записи с type = 1, но с точки зрения сервера такого типа просто не существует:
		// https://github.com/parserpro/fantlab/blob/f1e3aa00c05b0fd332259f4c580dcb07523fecd5/pm/Film.pm#L11-L18
		10: pb.FilmType_FILM_TYPE_FILM,
		20: pb.FilmType_FILM_TYPE_SERIES,
		21: pb.FilmType_FILM_TYPE_EPISODE,
		30: pb.FilmType_FILM_TYPE_DOCUMENTARY,
		40: pb.FilmType_FILM_TYPE_ANIMATION,
		50: pb.FilmType_FILM_TYPE_SHORT,
		60: pb.FilmType_FILM_TYPE_SPECTACLE,
	}

	BookcaseTypeMap = map[string]pb.Bookcase_BookcaseType{
		"read": pb.Bookcase_BOOKCASE_TYPE_READ,
		"wait": pb.Bookcase_BOOKCASE_TYPE_WAIT,
		"buy":  pb.Bookcase_BOOKCASE_TYPE_BUY,
		"sale": pb.Bookcase_BOOKCASE_TYPE_SALE,
		"free": pb.Bookcase_BOOKCASE_TYPE_FREE,
	}
)

func GetGender(userId uint64, rawUserSex uint8) pb.Common_Gender {
	if userId > 0 {
		if rawUserSex == 0 {
			return pb.Common_GENDER_FEMALE
		}
		return pb.Common_GENDER_MALE
	}
	return pb.Common_GENDER_UNKNOWN
}

func GetWorkPublishStatus(rawPublishStatus uint8, notFinished bool, planned bool) (result []pb.Work_PublishStatus) {
	if notFinished {
		result = append(result, pb.Work_PUBLISH_STATUS_NOT_FINISHED)
	}
	switch rawPublishStatus {
	case 0:
		result = append(result, pb.Work_PUBLISH_STATUS_NOT_PUBLISHED)
	case 2:
		result = append(result, pb.Work_PUBLISH_STATUS_NETWORK_PUBLICATION)
	case 3:
		result = append(result, pb.Work_PUBLISH_STATUS_AVAILABLE_ONLINE)
	default:
		break
	}
	if planned {
		result = append(result, pb.Work_PUBLISH_STATUS_PLANNED_BY_THE_AUTHOR)
	}
	return
}

func GetWorkCycleType(rawWorkType uint64) pb.WorkType {
	switch rawWorkType {
	case 4, 13, 48:
		return pb.WorkType_WORK_TYPE_SERIES
	default:
		return pb.WorkType_WORK_TYPE_OTHER
	}
}
