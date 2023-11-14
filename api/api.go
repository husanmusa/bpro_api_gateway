package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/husanmusa/bpro_api_gateway/api/handlers"
	"github.com/husanmusa/bpro_api_gateway/config"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
	logger2 "github.com/saidamir98/udevs_pkg/logger"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/husanmusa/bpro_api_gateway/api/docs"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// SetUpRouter godoc
// @description This is a api gateway
// @termsOfService https://udevs.io
func SetUpRouter(h handlers.Handler, cfg config.Config, log logger2.LoggerI) *fiber.App {
	router := fiber.New(fiber.Config{JSONEncoder: json.Marshal, BodyLimit: 100 * 1024 * 1024})
	router.Use(logger.New(), cors.New())
	// router.Use(h.MaxAllowed(12000))
	router.Use(func(c *fiber.Ctx) error {
		log.Info("request", logger2.String("method", c.Method()), logger2.String("path", c.Path()),
			logger2.String("ip", c.Get("X-Forwarded-For")), logger2.String("ip", c.Get("X-Real-IP")))
		return c.Next()
	})
	router.Server().MaxConnsPerIP = 20
	router.Server().DisableKeepalive = true
	router.Server().IdleTimeout = time.Hour
	router.Server().LogAllErrors = true

	router.Get("/api/swagger/*", swagger.HandlerDefault)
	r := router.Group("/api")

	// BOOK
	r.Post("/book", h.CreateBook)
	r.Get("/book/:book_id", h.GetBook)
	r.Get("/book", h.GetBookList)
	r.Put("/book/:book_id", h.UpdateBook)
	r.Delete("/book/:book_id", h.DeleteBook)

	// BOOK CATEGORY
	r.Post("/book_category", h.CreateBookCategory)
	r.Get("/book_category/:book_category_id", h.GetBookCategory)
	r.Get("/book_category", h.GetBookCategoryList)
	r.Put("/book_category/:book_category_id", h.UpdateBookCategory)
	r.Delete("/book_category/:book_category_id", h.DeleteBookCategory)
	return router
}
