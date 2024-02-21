package service

import (
	"exam3/pkg/logger"
	"exam3/storage"
)

type IServiceManager interface {
	Book() bookService
}

type Service struct {
	bookService bookService
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	services := Service{}

	services.bookService = NewBookService(storage, log)

	return services
}

func (s Service) Book() bookService {
	return s.bookService
}
