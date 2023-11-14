package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husanmusa/bpro_api_gateway/api/http"
	"github.com/husanmusa/bpro_api_gateway/config"
	"github.com/husanmusa/bpro_api_gateway/grpc/client"
	"github.com/saidamir98/udevs_pkg/logger"
	"strconv"
)

type Handler struct {
	cfg      config.Config
	log      logger.LoggerI
	services client.ServiceManagerI
}

type Map map[string]interface{}

func NewHandler(cfg config.Config, log logger.LoggerI, svcs client.ServiceManagerI) Handler {
	return Handler{
		cfg:      cfg,
		log:      log,
		services: svcs,
	}
}

func (h *Handler) handleResponse(c *fiber.Ctx, status http.Status, data interface{}, err ...string) error {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	return c.Status(status.Code).JSON(http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
		Error:       err,
	})
}

func (h *Handler) handleIntegrationResponse(c *fiber.Ctx, status http.Status, data interface{}) error {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	return c.Status(status.Code).JSON(data)
}

func (h *Handler) getOffsetParam(c *fiber.Ctx) (offset int, err error) {
	offsetStr := c.Query("offset", h.cfg.DefaultOffset)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) getLimitParam(c *fiber.Ctx) (offset int, err error) {
	offsetStr := c.Query("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) MaxAllowed(n int) fiber.Handler {
	var countReq int64
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
		countReq++
	}

	release := func() {
		select {
		case <-sem:
		default:
		}
		countReq--
	}

	return func(c *fiber.Ctx) error {
		lang := c.Get("lang", "uz")
		if lang != "ru" && lang != "uz" {
			c.Context().Request.Header.Set("lang", "uz")
		}

		acquire()       // before request
		defer release() // after request

		return c.Next()
	}
}
