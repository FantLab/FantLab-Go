package authapi

import "fantlab/pb"

func createSession(userData dbUserPasswordHash, token string) *pb.UserSessionResponse {
	return &pb.UserSessionResponse{
		UserId:       userData.UserID,
		SessionToken: token,
	}
}
