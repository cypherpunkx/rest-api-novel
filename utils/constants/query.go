package constants

const (
	INSERT_NOVEL = "INSERT INTO novels (id, code, title, publisher, year, author) VALUES ($1, $2, $3, $4, $5, $6)"
	NOVEL_LIST   = `
	SELECT
		id,
		code,
		title,
		publisher,
		year,
		author,
		created_at,
		updated_at
	FROM novels
	`

	NOVEL_GET = `
	SELECT
		id,
		code,
		title,
		publisher,
		year,
		author,
		created_at,
		updated_at
	FROM novels
	WHERE id = $1
	`

	NOVEL_COUNT = `
	SELECT COUNT(id) FROM novels
	`

	NOVEL_UPDATE = `UPDATE novels 
	SET 
	code = $2,
	title = $3,
	publisher = $4,
	year = $5,
	author = $6,
	updated_at = NOW()
	WHERE id = $1 RETURNING *`

	NOVEL_DELETE = "DELETE FROM novels WHERE id = $1"

	NOVEL_LIST_ALL = `
	SELECT
		id,
		code,
		title,
		publisher,
		year,
		author,
		created_at,
		updated_at
	FROM novels
	`
)
