package vacancy

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type vacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	v := &vacancyHandler{
		router:       router,
		customLogger: customLogger,
	}
	vacancyGroup := v.router.Group("/vacancy")
	vacancyGroup.Post("/", v.createVacancy)
}

func (h *vacancyHandler) createVacancy(c *fiber.Ctx) error {
	return c.SendString("vacancy created")
}
