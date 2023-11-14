package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husanmusa/bpro_api_gateway/api/http"
	pb "github.com/husanmusa/bpro_api_gateway/genproto/book_pro_service"
)

// CreateBookCategory godoc
// @Security ApiKeyAuth
// @Summary Create a new bookCategory
// @Description This API for creating a new bookCategory
// ID create_bookCategory
// @Tags BookCategory
// @Accept json
// @Produce json
// @Param bookCategory body book_pro_service.BookCategory true "BookCategoryCreateRequest"
// Success 201 {object} http.Response{data=string} "BookCategory data"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/book_Category/ [POST]
func (h *Handler) CreateBookCategory(c *fiber.Ctx) error {
	ok, err := h.HasAccess(c, "bookUpt", "write")
	if !ok {
		return err
	}

	var bookCategory pb.BookCategory

	err = c.BodyParser(&bookCategory)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	resp, err := h.services.BookCategoryService().CreateBookCategory(
		c.Context(),
		&bookCategory,
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.Created, resp)
}

// GetBookCategory godoc
// @Security ApiKeyAuth
// @Summary Get bookCategory by bookCategory_id
// ID get_bookCategory
// @Tags BookCategory
// @Accept json
// @Produce json
// @Param bookCategory_id path string true "bookCategory_id"
// @Success 200 {object} http.Response{data=book_pro_service.BookCategory} "GetBookCategory ResponseBody"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/book_category/{book_category_id} [GET]
func (h *Handler) GetBookCategory(c *fiber.Ctx) error {
	ok, err := h.HasAccess(c, "bookGet", "read")
	if !ok {
		return err
	}

	id := c.Params("book_category_id")
	resp, err := h.services.BookCategoryService().GetBookCategory(
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

// GetBookCategoryList godoc
// @Security ApiKeyAuth
// @Summary Get bookCategory list
// ID get_bookCategory_list
// @Tags BookCategory
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} http.Response{data=book_pro_service.GetBookCategoryListRes} "GetBookCategory ResponseBody"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/book_category [GET]
func (h *Handler) GetBookCategoryList(c *fiber.Ctx) error {
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

	resp, err := h.services.BookCategoryService().GetBookCategoryList(
		c.Context(),
		&pb.GetBookCategoryListReq{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// UpdateBookCategory godoc
// @Security ApiKeyAuth
// @Summary Update bookCategory by_id
// @ID update_bookCategory
// @Tags BookCategory
// @Accept json
// @Produce json
// @Param bookCategory body book_pro_service.BookCategory true "BookCategoryUpdateRequest"
// @Success 200 {object} http.Response{data=string} "Success Update"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/book_category/{book_category_id} [PUT]
func (h *Handler) UpdateBookCategory(c *fiber.Ctx) error {
	ok, err := h.HasAccess(c, "bookUpt", "write")
	if !ok {
		return err
	}

	var bookCategory pb.BookCategory
	bookCategory.Id = c.Params("book_category_id")

	err = c.BodyParser(&bookCategory)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	_, err = h.services.BookCategoryService().UpdateBookCategory(
		c.Context(),
		&bookCategory,
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, "SUCCESS")
}

// DeleteBookCategory godoc
// @Security ApiKeyAuth
// @Summary Delete bookCategory by_id
// @ID delete_bookCategory
// @Tags BookCategory
// @Accept json
// @Produce json
// @Param bookCategory_id path string false "bookCategory_id"
// @Success 200 {object} http.Response{data=string} "Success DeleteBookCategory"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/book_category/{book_category_id} [DELETE]
func (h *Handler) DeleteBookCategory(c *fiber.Ctx) error {
	ok, err := h.HasAccess(c, "bookUpt", "write")
	if !ok {
		return err
	}

	id := c.Params("book_category_id")

	_, err = h.services.BookCategoryService().DeleteBookCategory(
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
