package queries

const (
	ResponseGetResponse = `
		SELECT
			response_id,
			user_id,
			work_id,
			vote_plus,
			ABS(vote_minus) AS vote_minus
		FROM
			responses
		WHERE
			response_id = ?
	`

	ResponseGetResponseUserVoteCount = `
		SELECT 
			COUNT(*) 
		FROM 
			responses_votes 
		WHERE 
			user_id = ? AND response_id = ?
	`

	ResponseInsertResponseVote = `
		INSERT 
			responses_votes
		SET 
			user_id = ?, 
			response_id = ?, 
			voteone = ?, 
			date_of_vote = NOW()
	`

	ResponseUpdateResponseVotePlus = `
		UPDATE 
			responses 
		SET
			vote_plus = vote_plus + 1
		WHERE 
			response_id = ?
	`

	ResponseUpdateResponseVoteMinus = `
		UPDATE 
			responses 
		SET 
			vote_minus = vote_minus - 1
		WHERE 
			response_id = ?
	`

	ResponseUpdateResponse = `
		UPDATE
			responses
		SET
			response = ?
		WHERE
			response_id = ?
	`

	ResponseDeleteResponse = `
		DELETE
		FROM
			responses
		WHERE
			response_id = ?
	`
)
