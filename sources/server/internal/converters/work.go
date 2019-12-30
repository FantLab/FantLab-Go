package converters

import (
	"fantlab/pb"
	"fantlab/server/internal/db"
	"fantlab/server/internal/helpers"
)

func GetSubWorks(rootWorkId uint64, children []db.WorkChild) *pb.Work_SubWorksResponse {
	root := &pb.Work_SubWork{
		Id: rootWorkId,
	}

	workTable := make(map[uint64]*pb.Work_SubWork)
	closeTable := make(map[uint64]bool)

	for _, child := range children {
		closeTable[child.Id] = child.ShowSubworks > 0

		workTable[child.Id] = &pb.Work_SubWork{
			Id:            child.Id,
			OrigName:      child.OrigName,
			RusName:       child.RusName,
			Year:          child.Year,
			WorkType:      helpers.GetWorkType(child.WorkType),
			Rating:        child.Midmark,
			Marks:         child.Marks,
			Reviews:       child.Reviews,
			Plus:          child.IsBonus > 0,
			PublishStatus: helpers.GetWorkPublishStatus(child.IsPublished, child.NotFinished > 0, child.IsPlanned > 0),
		}
	}

	for _, child := range children {
		work := workTable[child.Id]
		if work == nil {
			continue
		}

		parentWork := workTable[child.ParentId]

		if parentWork == nil {
			if child.ParentId == rootWorkId {
				root.Subworks = append(root.Subworks, work)
			}
		} else if !closeTable[parentWork.Id] {
			parentWork.Subworks = append(parentWork.Subworks, work)
		}
	}

	return &pb.Work_SubWorksResponse{
		WorkId:   rootWorkId,
		Subworks: root.Subworks,
	}
}
