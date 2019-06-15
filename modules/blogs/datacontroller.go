package blogsapi

func getCommunities(dbCommunities []dbCommunity) communitiesWrapper {
	var mainCommunities []community
	var additionalCommunities []community

	for _, dbCommunity := range dbCommunities {
		community := community{
			Id:          dbCommunity.BlogId,
			Title:       dbCommunity.Name,
			Description: dbCommunity.Description,
			Stats: blogStats{
				ArticleCount:    dbCommunity.TopicsCount,
				SubscriberCount: dbCommunity.SubscriberCount,
			},
			LastArticle: lastCommunityArticle{
				Id:    dbCommunity.LastTopicId,
				Title: dbCommunity.LastTopicHead,
				User: userLink{
					Id:    dbCommunity.LastUserId,
					Login: dbCommunity.LastUserName,
				},
				Date: dbCommunity.DateOfAdd.Unix(),
			},
		}
		if dbCommunity.IsPublic {
			mainCommunities = append(mainCommunities, community)
		} else {
			additionalCommunities = append(additionalCommunities, community)
		}
	}

	return communitiesWrapper{
		Main:       mainCommunities,
		Additional: additionalCommunities,
	}
}
