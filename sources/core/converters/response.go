package converters

import (
	"fantlab/core/db"
	"fantlab/pb"
)

func GetResponseRating(dbResponse db.Response) *pb.Response_ResponseRatingResponse {
	return &pb.Response_ResponseRatingResponse{
		Rating: int64(dbResponse.VotePlus - dbResponse.VoteMinus),
	}
}
