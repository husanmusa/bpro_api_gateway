package handlers

import (
	"github.com/gofiber/fiber/v2"
	pba "github.com/husanmusa/bpro_api_gateway/genproto/auth_service"
	"net/http"
	"strings"
)

func (h *Handler) HasAccess(c *fiber.Ctx, obj, act string) (bool, error) {
	const BearerSchema string = "Bearer "
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return false, c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "No Authorization header found",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, BearerSchema)
	ok, err := h.services.UserService().HasAccess(c.Context(), &pba.HasAccessReq{Token: tokenString, Obj: obj, Act: act})
	if err != nil || !ok.HasAccess {
		return false, c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return true, nil
}
