package helpers

import "fantlab/pb"

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
