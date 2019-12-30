package queries

const (
	WorkGenreCacheTable = "work_genre_cache"
	UserWorkGenresTable = "user_work_genres"
)

const (
	Genres = `
		SELECT work_genre_id, parent_work_genre_id, work_genre_group_id, name, description
		FROM work_genres
		ORDER BY work_genre_group_id ASC, level ASC
	`
	GenreGroups             = "SELECT work_genre_group_id, name FROM work_genre_groups ORDER BY level ASC"
	UserWorkGenreIds        = "SELECT work_genre_id FROM user_work_genres WHERE user_id = ? AND work_id = ?"
	GenreWorkCounts         = "SELECT work_genre_id, work_count_voting_finished FROM work_genres WHERE work_count_voting_finished > 0"
	WorkGenreVotes          = "SELECT work_genre_id, vote_count FROM work_genre_cache WHERE work_id = ?"
	WorkClassificationCount = "SELECT count(distinct user_id) FROM user_work_genres WHERE work_id = ?"
	DeleteUserGenres        = "DELETE FROM user_work_genres WHERE user_id = ? AND work_id = ?"
	InsertUserGenres        = "INSERT INTO user_work_genres (user_id, work_id, work_genre_id, date_of_add) VALUES %s"
	UpdateWorkVoterCount    = `
		UPDATE works
		SET voter_count = (
			SELECT COUNT(DISTINCT user_work_genres.user_id)
			FROM user_work_genres
			WHERE user_work_genres.work_id = works.work_id
		)
		WHERE work_id = ?
	`
	UserClassifCount       = "SELECT COUNT(DISTINCT work_id) FROM user_work_genres WHERE user_id = ?"
	UpdateUserClassifCount = "UPDATE users SET classifcount = ?, need_recalc_level = 1 WHERE user_id = ? AND classifcount != ?"
	WorkGenreVoteCounts    = `
		SELECT work_genre_id, COUNT(*)
		FROM user_work_genres
		WHERE work_id = ?
		GROUP BY work_genre_id
	`
	WorkClassifCountAfterGenreAdd = `
		SELECT wg.work_genre_id, COUNT(DISTINCT user_id)
		FROM work_genres wg
		JOIN user_work_genres uwg ON (uwg.work_id = ? AND uwg.date_of_add >= wg.date_of_add)
		GROUP BY wg.work_genre_id
	`
	DeleteWorkGenreCache = "DELETE FROM work_genre_cache WHERE work_id = ?"
)
