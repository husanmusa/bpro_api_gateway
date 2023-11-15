package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husanmusa/bpro_api_gateway/api/http"
	pb "github.com/husanmusa/bpro_api_gateway/genproto/book_pro_service"
)

// CreateBook godoc
// @Security ApiKeyAuth
// @Summary Create a new book
// @Description This API for creating a new book
// ID create_book
// @Tags Book
// @Accept json
// @Produce json
// @Param book body book_pro_service.Book true "BookCreateRequest"
// Success 201 {object} http.Response{data=string} "Book data"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/book/ [POST]
func (h *Handler) CreateBook(c *fiber.Ctx) error {
	ok, err := h.HasAccess(c, "bookUpt", "write")
	if !ok {
		return err
	}

	var book pb.Book

	err = c.BodyParser(&book)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	resp, err := h.services.BookService().CreateBook(
		c.Context(),
		&book,
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.Created, resp)
}

// GetBook godoc
// @Security ApiKeyAuth
// @Summary Get book by book_id
// ID get_book
// @Tags Book
// @Accept json
// @Produce json
// @Param book_id path string true "book_id"
// @Success 200 {object} http.Response{data=book_pro_service.Book} "GetBook ResponseBody"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/book/{book_id} [GET]
func (h *Handler) GetBook(c *fiber.Ctx) error {
	ok, err := h.HasAccess(c, "bookGet", "read")
	if !ok {
		return err
	}

	id := c.Params("book_id")
	resp, err := h.services.BookService().GetBook(
		c.Context(),
		&pb.ById{
			Id: id,
		},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// GetBookList godoc
// @Security ApiKeyAuth
// @Summary Get book list
// ID get_book_list
// @Tags Book
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param category_id query string false "category_id"
// @Success 200 {object} http.Response{data=book_pro_service.GetBookListResponse} "GetBook ResponseBody"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/book [GET]
func (h *Handler) GetBookList(c *fiber.Ctx) error {
	ok, err := h.HasAccess(c, "bookGet", "read")
	if !ok {
		return err
	}

	limit, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.InvalidArgument, err.Error())
	}
	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.InvalidArgument, err.Error())
	}

	resp, err := h.services.BookService().GetBookList(
		c.Context(),
		&pb.GetBookListRequest{
			Limit:          int32(limit),
			Offset:         int32(offset),
			BookCategoryId: c.Query("category_id"),
		},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// UpdateBook godoc
// @Security ApiKeyAuth
// @Summary Update book by_id
// @ID update_book
// @Tags Book
// @Accept json
// @Produce json
// @Param book_id path string false "book_id"
// @Param book body book_pro_service.Book true "BookUpdateRequest"
// @Success 200 {object} http.Response{data=string} "Success Update"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/book/{book_id} [PUT]
func (h *Handler) UpdateBook(c *fiber.Ctx) error {
	ok, err := h.HasAccess(c, "bookUpt", "write")
	if !ok {
		return err
	}

	var book pb.Book
	book.Id = c.Params("book_id")

	err = c.BodyParser(&book)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	_, err = h.services.BookService().UpdateBook(
		c.Context(),
		&book,
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, "SUCCESS")
}

// DeleteBook godoc
// @Security ApiKeyAuth
// @Summary Delete book by_id
// @ID delete_book
// @Tags Book
// @Accept json
// @Produce json
// @Param book_id path string false "book_id"
// @Success 200 {object} http.Response{data=string} "Success DeleteBook"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/book/{book_id} [DELETE]
func (h *Handler) DeleteBook(c *fiber.Ctx) error {
	ok, err := h.HasAccess(c, "bookUpt", "write")
	if !ok {
		return err
	}

	id := c.Params("book_id")

	_, err = h.services.BookService().DeleteBook(
		c.Context(),
		&pb.ById{
			Id: id,
		},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, "SUCCESS")
}
