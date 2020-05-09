package converters

import (
	"fantlab/base/protobuf/pbutils"
	"fantlab/core/config"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
)

func GetPrivateMessage(dbMessage db.PrivateMessage, cfg *config.AppConfig) *pb.Private_MessageResponse {
	text := dbMessage.MessageText

	gender := helpers.GetGender(dbMessage.FromUserId, dbMessage.Sex)
	avatar := helpers.GetUserAvatarUrl(cfg.ImagesBaseURL, dbMessage.FromUserId, dbMessage.PhotoNumber)

	message := &pb.Private_Message{
		Id: dbMessage.PrivateMessageId,
		Creation: &pb.Common_Creation{
			User: &pb.Common_UserLink{
				Id:     dbMessage.FromUserId,
				Login:  dbMessage.Login,
				Gender: gender,
				Avatar: avatar,
				Class:  helpers.UserClassMap[dbMessage.UserClass],
				Sign:   dbMessage.Sign,
			},
			Date: pbutils.TimestampProto(dbMessage.DateOfAdd),
		},
		Text:   text,
		Number: dbMessage.Number,
		IsRead: dbMessage.IsRead == 1,
	}

	return &pb.Private_MessageResponse{
		Message: message,
	}
}
