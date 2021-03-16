package queries

const (
	ForumGetMessageUserVoteExists = `
		SELECT
			EXISTS (
				SELECT
					*
				FROM
					f_messages_votes
				WHERE
					user_id = ? AND message_id = ?
			)
	`

	ForumMessageVoteInsert = `
		INSERT 
			f_messages_votes 
		SET 
			user_id = ?,
			message_id = ?,
			voteone = ?,
			date_of_vote = NOW()
	`

	ForumMessageVotePlusUpdate = `
		UPDATE 
			f_messages 
		SET
			vote_plus = vote_plus + 1
		WHERE 
			message_id = ?
	`

	ForumMessageVoteMinusUpdate = `
		UPDATE 
			f_messages 
		SET 
			vote_minus = vote_minus - 1
		WHERE 
			message_id = ?
	`

	ForumMessageVotesDelete = `
		DELETE FROM 
			f_messages_votes 
		WHERE 
			message_id = ?
	`

	// NOTE Неясно, зачем выставляется флаг is_red. Предоложительно, чтобы запретить повторную оценку сообщения после
	// удаления его оценки модератором. Но подобное применение флага нивелирует его смысл как стопроцентного индикатора
	// модераторского сообщения. Похоже на баг. https://github.com/parserpro/fantlab/issues/945
	ForumMessageVoteCountUpdateByModerator = `
		UPDATE 
			f_messages 
		SET 
			is_red = 1,
			vote_plus = 0,
			vote_minus = 0 
		WHERE 
			message_id = ?
	`
)
