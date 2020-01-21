package middlewares

import (
	"fantlab/base/protobuf"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func CheckBan(services *app.Services) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := app.GetUserAuth(r.Context())

			banInfo, err := services.DB().FetchUserBlockInfo(r.Context(), auth.User.UserId)

			if err != nil {
				protobuf.Handle(func(r *http.Request) (int, proto.Message) {
					return http.StatusInternalServerError, &pb.Error_Response{
						Status: pb.Error_SOMETHING_WENT_WRONG,
					}
				}).ServeHTTP(w, r)
			} else {
				if banInfo.Blocked > 0 {
					protobuf.Handle(func(r *http.Request) (int, proto.Message) {
						return http.StatusForbidden, &pb.Error_Response{
							Status:  pb.Error_USER_IS_BANNED,
							Context: banInfo.BlockReason,
						}
					}).ServeHTTP(w, r)
				} else {
					next.ServeHTTP(w, r)
				}
			}
		})
	}
}
