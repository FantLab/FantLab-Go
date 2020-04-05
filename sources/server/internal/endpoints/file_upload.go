package endpoints

import (
	"fantlab/pb"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"
)

// Возвращает урл (удалить после привязки к форумам/блогам)
func (api *API) FileUpload(r *http.Request) (int, proto.Message) {
	var params struct {
		// Путь к файлу на сервере (напр. forum/14/image.jpg)
		PathToFile string `http:"path_to_file,form"`
	}

	api.bindParams(&params, r)

	url := api.services.GetFileUploadURL(r.Context(), params.PathToFile, 5*time.Minute)
	if url == nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}
	return http.StatusOK, &pb.Common_FileUploadResponse{
		Url: url.String(),
	}
}
