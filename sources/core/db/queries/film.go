package queries

const (
	FilmGetFilm = `
		SELECT
			film_id,
			name
		FROM
			films
		WHERE
			film_id = ?
	`
)
