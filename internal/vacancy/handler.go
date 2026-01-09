package vacancy

import (
	"demo/go-fiber/pkg/tadaptor"
	"demo/go-fiber/views/components"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type vacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &vacancyHandler{
		router:       router,
		customLogger: customLogger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *vacancyHandler) createVacancy(c *fiber.Ctx) error {
	email := c.FormValue("email")
	h.customLogger.Info().Msg(email)
	component := components.Notification("Vacancy successfully created")
	return tadaptor.Render(c, component)
}
