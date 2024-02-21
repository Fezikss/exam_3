package handler

import (
	"context"
	"exam3/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateBook godoc
// @Router       /book [POST]
// @Summary      Create a new book
// @Description  create a new book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 book body models.CreateBook false "book"
// @Success      200  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBook(c *gin.Context) {
	book := models.CreateBook{}

	if err := c.ShouldBindJSON(&book); err != nil {
		handleResponse(c, h.log, "error is while reading body from client", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.Book().Create(context.Background(), book)
	if err != nil {
		handleResponse(c, h.log, "error is while creating category", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusCreated, resp)
}

// GetBook godoc
// @Router       /book/{id} [GET]
// @Summary      Get book by id
// @Description  get book by id
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 id path string true "book_id"
// @Success      200  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBook(c *gin.Context) {
	uid := c.Param("id")

	book, err := h.services.Book().Get(context.Background(), uid)
	if err != nil {
		handleResponse(c, h.log, "error is while getting by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, book)
}

// GetBookList godoc
// @Router       /books [GET]
// @Summary      Get book list
// @Description  get book list
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.BooksResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBookList(c *gin.Context) {
	var (
		page, limit int
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, h.log, "error while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, h.log, "error while converting limit", http.StatusBadRequest, err.Error())
		return
	}

	search := c.Query("search")

	response, err := h.services.Book().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, h.log, "error while getting book list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, response)
}

// UpdateBook godoc
// @Router       /book/{id} [PUT]
// @Summary      Update book
// @Description  get book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 id path string true "book_id"
// @Param 		 book body models.UpdateBook false "book"
// @Success      200  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBook(c *gin.Context) {
	updateBook := models.UpdateBook{}

	uid := c.Param("id")
	if err := c.ShouldBindJSON(&updateBook); err != nil {
		handleResponse(c, h.log, "error is while decoding ", http.StatusBadRequest, err)
		return
	}

	updateBook.ID = uid

	book, err := h.services.Book().Update(context.Background(), updateBook)
	if err != nil {
		handleResponse(c, h.log, "error is while updating book", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, book)
}

// UpdatePageNumber godoc
// @Router       /book/{id} [PATCH]
// @Summary      Update book page number
// @Description  update book page number
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 id path string true "book_id"
// @Param        book body models.UpdateBookPageNumber true "book"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdatePageNumber(c *gin.Context) {
	updateBookPageNumber := models.UpdateBookPageNumber{}

	if err := c.ShouldBindJSON(&updateBookPageNumber); err != nil {
		handleResponse(c, h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		handleResponse(c, h.log, "error while parsing uuid", http.StatusBadRequest, err.Error())
		return
	}

	updateBookPageNumber.ID = uid.String()

	book, err := h.services.Book().UpdatePageNumber(context.Background(), updateBookPageNumber)
	if err != nil {
		handleResponse(c, h.log, "error while updating book page number", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, book)
}

// DeleteBook godoc
// @Router       /book/{id} [DELETE]
// @Summary      Delete book
// @Description  delete book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 id path string true "book_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBook(c *gin.Context) {
	uid := c.Param("id")

	if err := h.services.Book().Delete(context.Background(), uid); err != nil {
		handleResponse(c, h.log, "error is while deleting basket", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, nil)
}
