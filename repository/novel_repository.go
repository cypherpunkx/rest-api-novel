package repository

import (
	"database/sql"

	"enigmacamp.com/rest-api-novel/domain"
	"enigmacamp.com/rest-api-novel/utils/constants"
	"enigmacamp.com/rest-api-novel/utils/exception"
)

type NovelRepository interface {
	BaseRepository[domain.Novel]
}

type novelRepository struct {
	db *sql.DB
}

func (n *novelRepository) ListPaginate(page, perPage int, code, title, publisher, year, author string) ([]domain.Novel, error) {
	novels := []domain.Novel{}

	offset := (page - 1) * perPage

	query := constants.NOVEL_LIST + " WHERE 1=1"

	if code != "" {
		query += " AND code ILIKE '%" + code + "%'"
	}
	if title != "" {
		query += " AND title ILIKE '%" + title + "%'"
	}
	if publisher != "" {
		query += " AND publisher ILIKE '%" + publisher + "%'"
	}
	if year != "" {
		query += " AND year::text ILIKE '%" + year + "%'"
	}
	if author != "" {
		query += " AND author ILIKE '%" + author + "%'"
	}

	query += " ORDER BY created_at DESC LIMIT $1 OFFSET $2"

	rows, err := n.db.Query(query, perPage, offset)
	if err != nil {
		return novels, err
	}
	defer rows.Close()

	for rows.Next() {
		novel := domain.Novel{}

		err := rows.Scan(
			&novel.Id,
			&novel.Code,
			&novel.Title,
			&novel.Publisher,
			&novel.Year,
			&novel.Author,
			&novel.CreatedAt,
			&novel.UpdatedAt,
		)
		if err != nil {
			return novels, err
		}

		novels = append(novels, novel)
	}

	return novels, nil
}

func (n *novelRepository) List() ([]domain.Novel, error) {
	novels := []domain.Novel{}

	rows, err := n.db.Query(constants.NOVEL_LIST_ALL)
	if err != nil {
		return novels, err
	}
	defer rows.Close()

	for rows.Next() {
		novel := domain.Novel{}

		err := rows.Scan(
			&novel.Id,
			&novel.Code,
			&novel.Title,
			&novel.Publisher,
			&novel.Year,
			&novel.Author,
			&novel.CreatedAt,
			&novel.UpdatedAt,
		)
		if err != nil {
			return novels, err
		}

		novels = append(novels, novel)
	}

	return novels, nil
}

func (n *novelRepository) Get(id string) (domain.Novel, error) {
	novel := domain.Novel{}

	err := n.db.QueryRow(
		constants.NOVEL_GET,
		id,
	).Scan(
		&novel.Id,
		&novel.Code,
		&novel.Title,
		&novel.Publisher,
		&novel.Year,
		&novel.Author,
		&novel.CreatedAt,
		&novel.UpdatedAt,
	)
	if err != nil {
		return novel, exception.ErrNotFound
	}

	return novel, nil
}

func (e *novelRepository) Create(payload domain.Novel) error {
	_, err := e.db.Exec(constants.INSERT_NOVEL, payload.Id, payload.Code, payload.Title,
		payload.Publisher, payload.Year, payload.Author)
	if err != nil {
		return err
	}

	return nil
}

func (r *novelRepository) Update(id string, payload domain.Novel) (*domain.Novel, error) {
	novel := domain.Novel{}

	query := constants.NOVEL_UPDATE

	err := r.db.QueryRow(query, id, payload.Code, payload.Title, payload.Publisher, payload.Year, payload.Author).Scan(&novel.Id, &novel.Code, &novel.Title, &novel.Publisher, &novel.Year, &novel.Author, &novel.CreatedAt, &novel.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &novel, nil
}

func (r *novelRepository) Delete(id string) error {

	query := constants.NOVEL_DELETE

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (n *novelRepository) Count(code, title, publisher, year, author string) (int, error) {
	query := constants.NOVEL_COUNT + " WHERE 1=1"

	if code != "" {
		query += " AND code ILIKE '%" + code + "%'"
	}
	if title != "" {
		query += " AND title ILIKE '%" + title + "%'"
	}
	if publisher != "" {
		query += " AND publisher ILIKE '%" + publisher + "%'"
	}
	if year != "" {
		query += " AND year::text ILIKE '%" + year + "%'"
	}
	if author != "" {
		query += " AND author ILIKE '%" + author + "%'"
	}

	var count int
	err := n.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func NewNovelRepository(db *sql.DB) NovelRepository {
	return &novelRepository{
		db: db,
	}
}
