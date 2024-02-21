package postgres

import (
	"context"
	"database/sql"
	"exam3/api/models"
	"exam3/pkg/logger"
	"exam3/storage"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type bookRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewBookRepo(db *pgxpool.Pool, log logger.ILogger) storage.IBookStorage {
	return bookRepo{
		db:  db,
		log: log,
	}
}

func (c bookRepo) Create(ctx context.Context, book models.CreateBook) (string, error) {
	id := uuid.New()
	query := `INSERT INTO books (id, name, author_name, page_number) VALUES($1, $2, $3, $4)`
	if _, err := c.db.Exec(ctx, query, id, book.Name, book.AuthorName, book.PageNumber); err != nil {
		c.log.Error("error is while inserting data", logger.Error(err))
		return "", err
	}
	return id.String(), nil
}

func (c bookRepo) GetByID(ctx context.Context, id string) (models.Book, error) {
	var updatedAt, createdAt sql.NullString
	book := models.Book{}
	query := `select id, name, author_name, page_number, created_at, updated_at FROM books WHERE id = $1 and deleted_at = 0`
	if err := c.db.QueryRow(ctx, query, id).Scan(
		&book.ID,
		&book.Name,
		&book.AuthorName,
		&book.PageNumber,
		&createdAt,
		&updatedAt,
	); err != nil {
		c.log.Error("error is while selecting by id", logger.Error(err))
		return models.Book{}, err
	}

	if createdAt.Valid {
		book.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		book.UpdatedAt = updatedAt.String
	}

	return book, nil
}

func (c bookRepo) GetList(ctx context.Context, request models.GetListRequest) (models.BooksResponse, error) {
	var (
		page                 = request.Page
		offset               = (page - 1) * request.Limit
		query, countQuery    string
		count                = 0
		books                = []models.Book{}
		search               = request.Search
		updatedAt, createdAt sql.NullString
	)
	countQuery = `SELECT count(1) FROM books WHERE deleted_at = 0 `
	if search != "" {
		countQuery += fmt.Sprintf(` and name ilike '%%%s%%'`, search)
	}
	if err := c.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		c.log.Error("error is while scanning count", logger.Error(err))
		return models.BooksResponse{}, err
	}

	query = `SELECT id, name, author_name, page_number, created_at, updated_at FROM books WHERE deleted_at = 0 `
	if search != "" {
		query += fmt.Sprintf(` and name ilike '%%%s%%'`, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := c.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		c.log.Error("error is while selecting all", logger.Error(err))
		return models.BooksResponse{}, err
	}

	for rows.Next() {
		book := models.Book{}
		if err = rows.Scan(
			&book.ID,
			&book.Name,
			&book.AuthorName,
			&book.PageNumber,
			&createdAt,
			&updatedAt,
		); err != nil {
			c.log.Error("error is while scanning book", logger.Error(err))
			return models.BooksResponse{}, err
		}

		if createdAt.Valid {
			book.CreatedAt = createdAt.String
		}

		if updatedAt.Valid {
			book.UpdatedAt = updatedAt.String
		}

		books = append(books, book)
	}
	return models.BooksResponse{
		Books: books,
		Count: count,
	}, nil
}

func (c bookRepo) Update(ctx context.Context, book models.UpdateBook) (string, error) {
	query := `UPDATE books SET name = $1, author_name = $2, updated_at = now() WHERE id = $3 AND deleted_at = 0`
	if _, err := c.db.Exec(ctx, query, &book.Name, book.AuthorName, &book.ID); err != nil {
		c.log.Error("error is while updating", logger.Error(err))
		return "", err
	}
	return book.ID, nil
}

func (c bookRepo) Delete(ctx context.Context, id string) error {
	query := `update books set deleted_at = extract(epoch FROM current_timestamp) WHERE id = $1`
	if _, err := c.db.Exec(ctx, query, id); err != nil {
		c.log.Error("error is while deleting", logger.Error(err))
		return err
	}
	return nil
}

func (c bookRepo) UpdatePageNumber(ctx context.Context, request models.UpdateBookPageNumber) (string, error) {
	query := `
		UPDATE books 
				SET page_number = $1, updated_at = now()
					WHERE id = $2 AND deleted_at = 0`

	if _, err := c.db.Exec(ctx, query, request.PageNumber, request.ID); err != nil {
		fmt.Println("error while updating page number", err.Error())
		return "", err
	}

	return request.ID, nil
}
