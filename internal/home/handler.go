package home

import (
	"demo/go-fiber/internal/vacancy"
	"demo/go-fiber/pkg/tadaptor"
	"demo/go-fiber/views"
	"demo/go-fiber/views/components"
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *vacancy.VacancyRepository
	store        *session.Store
}

type HomeHandlerDeps struct {
	Router       fiber.Router
	CustomLogger *zerolog.Logger
	Repository   *vacancy.VacancyRepository
	Store        *session.Store
}

type User struct {
	Id   string
	Name string
}

func NewHandler(deps HomeHandlerDeps) {
	h := &HomeHandler{
		router:       deps.Router,
		customLogger: deps.CustomLogger,
		repository:   deps.Repository,
		store:        deps.Store,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
	api.Post("/login", h.apiLogin)

	h.router.Get("/login", h.login)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)

	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	if name, ok := sess.Get("name").(string); ok {
		h.customLogger.Info().Msg(name)
	}

	count := h.repository.CountAll()
	vacancies, err := h.repository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}

	component := views.Main(vacancies, int(math.Ceil(float64(count)/float64(PAGE_ITEMS))), page)
	return tadaptor.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.ru").
		Int("id", 10).
		Msg("INFO")
	return fiber.NewError(fiber.StatusBadRequest, "Limit params is undefined")
}

func (h *HomeHandler) apiLogin(c *fiber.Ctx) error {
	form := LoginForm{
		Login:    c.FormValue("login"),
		Password: c.FormValue("password"),
	}

	if form.Login == "a@a.ru" && form.Password == "1" {
		sess, err := h.store.Get(c)

		if err != nil {
			panic(err)
		}
		sess.Set("login", form.Login)
		if err := sess.Save(); err != nil {
			panic(err)
		}

		c.Response().Header.Add("Hx-Redirect", "/api")

		return c.Redirect("/", http.StatusOK)
	}

	component := components.Notification("Wrong credentials", "fail")
	return tadaptor.Render(c, component, http.StatusBadRequest)
}

func (h *HomeHandler) login(c *fiber.Ctx) error {
	component := views.Login()
	return tadaptor.Render(c, component, http.StatusOK)
}
