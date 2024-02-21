package storage

import (
	"context"
	"exam3/api/models"
)

type IStorage interface {
	Close()
	Book() IBookStorage
}

type IBookStorage interface {
	Create(context.Context, models.CreateBook) (string, error)
	GetByID(context.Context, string) (models.Book, error)
	GetList(context.Context, models.GetListRequest) (models.BooksResponse, error)
	Update(context.Context, models.UpdateBook) (string, error)
	UpdatePageNumber(context.Context, models.UpdateBookPageNumber) (string, error)
	Delete(context.Context, string) error
}
