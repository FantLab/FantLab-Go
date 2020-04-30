package converters

import (
	"fantlab/pb"
	"fantlab/server/internal/db"
)

func GetResponseRating(dbResponse db.Response) *pb.Response_ResponseRatingResponse {
	return &pb.Response_ResponseRatingResponse{
		Rating: int64(dbResponse.VotePlus - dbResponse.VoteMinus),
	}
}
