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

func GetEditionCoverUrl(baseURL string, editionId uint64) string {
	return fmt.Sprintf("%s/editions/big/%d", baseURL, editionId)
}

func GetOzonOfferUrl(ozonOfferId uint64) string {
	return fmt.Sprintf("https://www.ozon.ru/context/detail/id/%d/?partner=iwant", ozonOfferId)
}

func GetLabirintOfferUrl(labirintOfferId uint64) string {
	return fmt.Sprintf("https://www.labirint.ru/books/%d/&p=1758", labirintOfferId)
}

func GetFilmPosterUrl(baseURL string, filmId uint64) string {
	return fmt.Sprintf("%s/films/poster/%d", baseURL, filmId)
}
