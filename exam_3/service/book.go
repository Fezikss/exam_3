package service

import (
	"context"

	"exam3/api/models"
	"exam3/pkg/logger"
	"exam3/storage"
)

type bookService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewBookService(storage storage.IStorage, log logger.ILogger) bookService {
	return bookService{
		storage: storage,
		log:     log,
	}
}

func (b bookService) Create(ctx context.Context, createBook models.CreateBook) (models.Book, error) {
	b.log.Info("book create service layer", logger.Any("book", createBook))

	pKey, err := b.storage.Book().Create(ctx, createBook)
	if err != nil {
		b.log.Error("ERROR in service layer while creating book", logger.Error(err))
		return models.Book{}, err
	}

	book, err := b.storage.Book().GetByID(ctx, pKey)
	if err != nil {
		b.log.Error("ERROR in service layer while getting book", logger.Error(err))
		return models.Book{}, err
	}

	return book, nil
}

func (b bookService) Get(ctx context.Context, id string) (models.Book, error) {
	book, err := b.storage.Book().GetByID(ctx, id)
	if err != nil {
		b.log.Error("error in service layer while getting book by id", logger.Error(err))
		return models.Book{}, err
	}

	return book, nil
}

func (b bookService) GetList(ctx context.Context, request models.GetListRequest) (models.BooksResponse, error) {
	b.log.Info("book get list service layer", logger.Any("book", request))

	books, err := b.storage.Book().GetList(ctx, request)
	if err != nil {
		b.log.Error("error in service layer  while getting list", logger.Error(err))
		return models.BooksResponse{}, err
	}

	return books, nil
}

func (b bookService) Update(ctx context.Context, book models.UpdateBook) (models.Book, error) {
	id, err := b.storage.Book().Update(ctx, book)
	if err != nil {
		b.log.Error("error in service layer while updating", logger.Error(err))
		return models.Book{}, err
	}

	updatedBook, err := b.storage.Book().GetByID(context.Background(), id)
	if err != nil {
		b.log.Error("error in service layer while getting Book by id", logger.Error(err))
		return models.Book{}, err
	}

	return updatedBook, nil
}

func (b bookService) UpdatePageNumber(ctx context.Context, pNumber models.UpdateBookPageNumber) (models.Book, error) {
	id, err := b.storage.Book().UpdatePageNumber(ctx, pNumber)
	if err != nil {
		b.log.Error("error in service layer while updating page number", logger.Error(err))
		return models.Book{}, err
	}

	updatedPageNumber, err := b.storage.Book().GetByID(context.Background(), id)
	if err != nil {
		b.log.Error("error in service layer while getting Book by id", logger.Error(err))
		return models.Book{}, err
	}

	return updatedPageNumber, nil
}

func (b bookService) Delete(ctx context.Context, key string) error {
	err := b.storage.Book().Delete(ctx, key)

	return err
}
