package helpers

import (
	"fmt"
)

func GetUserAvatarUrl(baseURL string, userId uint64, photoNumber uint64) string {
	if photoNumber > 0 {
		return fmt.Sprintf("%s/users/%d_%d", baseURL, userId, photoNumber)
	}
	return ""
}

func GetCommunityAvatarUrl(baseURL string, communityId uint64) string {
	return fmt.Sprintf("%s/communities/%d.jpg", baseURL, communityId)
}

func GetFilmPosterUrl(baseURL string, filmId uint64) string {
	return fmt.Sprintf("%s/films/poster/%d", baseURL, filmId)
}
