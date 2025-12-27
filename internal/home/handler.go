package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

type User struct {
	Id   string
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	users := []User{
		{Id: "TestId1", Name: "User1"},
		{Id: "TestId1", Name: "User2"},
	}
	names := []string{"martin", "john"}
	data := struct {
		Names []string
		Users []User
	}{Names: names, Users: users}
	return c.Render("page", data)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.ru").
		Int("id", 10).
		Msg("INFO")
	return fiber.NewError(fiber.StatusBadRequest, "Limit params is undefined")
}
