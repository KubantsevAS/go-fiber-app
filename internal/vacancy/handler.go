package vacancy

import (
	"demo/go-fiber/pkg/tadaptor"
	"demo/go-fiber/pkg/validator"
	"demo/go-fiber/views/components"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
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
	var component templ.Component
	form := VacancyCreateForm{
		Email: c.FormValue("email"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email is empty"},
	)
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), "fail")
		return tadaptor.Render(c, component)
	}

	h.customLogger.Info().Msg(form.Email)

	component = components.Notification("Vacancy successfully created", "success")
	return tadaptor.Render(c, component)
}
