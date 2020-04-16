package helpers

import "fantlab/pb"

func GetUserClass(rawUserClass uint8) pb.Common_UserClass {
	switch rawUserClass {
	case 0:
		return pb.Common_USERCLASS_BEGINNER
	case 1:
		return pb.Common_USERCLASS_ACTIVIST
	case 2:
		return pb.Common_USERCLASS_AUTHORITY
	case 3:
		return pb.Common_USERCLASS_PHILOSOPHER
	case 4:
		return pb.Common_USERCLASS_MASTER
	case 5:
		return pb.Common_USERCLASS_GRANDMASTER
	case 6:
		return pb.Common_USERCLASS_PEACEKEEPER
	case 7:
		return pb.Common_USERCLASS_PEACEMAKER
	default:
		return pb.Common_USERCLASS_UNKNOWN
	}
}

func GetUserClassDescription(userClass pb.Common_UserClass) string {
	switch userClass {
	case pb.Common_USERCLASS_BEGINNER:
		return "новичок"
	case pb.Common_USERCLASS_ACTIVIST:
		return "активист"
	case pb.Common_USERCLASS_AUTHORITY:
		return "авторитет"
	case pb.Common_USERCLASS_PHILOSOPHER:
		return "философ"
	case pb.Common_USERCLASS_MASTER:
		return "магистр"
	case pb.Common_USERCLASS_GRANDMASTER:
		return "гранд-мастер"
	case pb.Common_USERCLASS_PEACEKEEPER:
		return "миродержец"
	case pb.Common_USERCLASS_PEACEMAKER:
		return "миротворец"
	default:
		return ""
	}
}

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

func GetWorkCycleType(rawWorkType uint64) pb.Work_WorkType {
	switch rawWorkType {
	case 4, 13, 48:
		return pb.Work_WORK_TYPE_CYCLE
	default:
		return pb.Work_WORK_TYPE_OTHER
	}
}

func GetWorkType(rawWorkType uint64) pb.Work_WorkType {
	switch rawWorkType {
	case 1:
		return pb.Work_WORK_TYPE_NOVEL
	case 3:
		return pb.Work_WORK_TYPE_COLLECTION
	case 4:
		return pb.Work_WORK_TYPE_CYCLE
	case 5, 27, 28, 29:
		return pb.Work_WORK_TYPE_POEM
	case 7:
		return pb.Work_WORK_TYPE_OTHER
	case 8:
		return pb.Work_WORK_TYPE_TALE
	case 11:
		return pb.Work_WORK_TYPE_ESSAY
	case 12:
		return pb.Work_WORK_TYPE_ARTICLE
	case 13:
		return pb.Work_WORK_TYPE_EPIC
	case 17:
		return pb.Work_WORK_TYPE_ANTOLOGY
	case 18:
		return pb.Work_WORK_TYPE_PIECE
	case 19:
		return pb.Work_WORK_TYPE_SCENARIO
	case 20:
		return pb.Work_WORK_TYPE_DOCUMENTAL
	case 21:
		return pb.Work_WORK_TYPE_MICROSTORY
	case 22:
		return pb.Work_WORK_TYPE_DISSER
	case 23:
		return pb.Work_WORK_TYPE_MONOGRAPHY
	case 24:
		return pb.Work_WORK_TYPE_STUDY
	case 25:
		return pb.Work_WORK_TYPE_ENCYCLOPEDY
	case 26:
		return pb.Work_WORK_TYPE_MAGAZINE
	case 41:
		return pb.Work_WORK_TYPE_COMIX
	case 42:
		return pb.Work_WORK_TYPE_MANGA
	case 43:
		return pb.Work_WORK_TYPE_GRAPHICNOVEL
	case 44:
		return pb.Work_WORK_TYPE_STORY
	case 45:
		return pb.Work_WORK_TYPE_SHORTSTORY
	case 46:
		return pb.Work_WORK_TYPE_SKETCH
	case 47:
		return pb.Work_WORK_TYPE_REPORTAGE
	case 48:
		return pb.Work_WORK_TYPE_CONDITIONALCYCLE
	case 49:
		return pb.Work_WORK_TYPE_EXCERPT
	case 51:
		return pb.Work_WORK_TYPE_INTERVIEW
	case 52:
		return pb.Work_WORK_TYPE_REVIEW
	default:
		return pb.Work_WORK_TYPE_UNKNOWN
	}
}

func GetBookcaseType(rawBookcaseType string) pb.Bookcase_BookcaseType {
	switch rawBookcaseType {
	case "read":
		return pb.Bookcase_BOOKCASE_TYPE_READ
	case "wait":
		return pb.Bookcase_BOOKCASE_TYPE_WAIT
	case "buy":
		return pb.Bookcase_BOOKCASE_TYPE_BUY
	case "sale":
		return pb.Bookcase_BOOKCASE_TYPE_SALE
	case "free":
		return pb.Bookcase_BOOKCASE_TYPE_FREE
	default:
		return pb.Bookcase_BOOKCASE_TYPE_UNKNOWN
	}
}
