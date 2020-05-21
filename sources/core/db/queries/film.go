package queries

const (
	FilmGetFilms = `
		SELECT
			film_id,
			name
		FROM
			films
		WHERE
			film_id IN (?)
	`
)
