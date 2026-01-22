package vacancy

import (
	"demo/go-fiber/pkg/tadaptor"
	"demo/go-fiber/pkg/validator"
	"demo/go-fiber/views/components"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type vacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *VacancyRepository
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *VacancyRepository) {
	h := &vacancyHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
	vacancyGroup.Get("/getall", h.getAll)
}

func (h *vacancyHandler) getAll(c *fiber.Ctx) error {
	vacancies, err := h.repository.getAll()
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
	}
	return c.JSON(vacancies)
}

func (h *vacancyHandler) createVacancy(c *fiber.Ctx) error {
	var component templ.Component
	form := VacancyCreateForm{
		Email:        c.FormValue("email"),
		JobTitle:     c.FormValue("job_title"),
		Company:      c.FormValue("company"),
		CompanyScope: c.FormValue("company_scope"),
		Salary:       c.FormValue("salary"),
		Location:     c.FormValue("location"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email is empty"},
		&validators.StringIsPresent{Name: "Job Title", Field: form.JobTitle, Message: "Job Title is empty"},
		&validators.StringIsPresent{Name: "Company", Field: form.Company, Message: "Company is empty"},
		&validators.StringIsPresent{Name: "Company scope", Field: form.CompanyScope, Message: "Company scope is empty"},
		&validators.StringIsPresent{Name: "Salary", Field: form.Salary, Message: "Salary is empty"},
		&validators.StringIsPresent{Name: "Location", Field: form.Location, Message: "Location is empty"},
	)

	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), "fail")
		return tadaptor.Render(c, component, http.StatusBadRequest)
	}

	h.customLogger.Info().Msg(form.Email)

	if err := h.repository.addVacancy(form); err != nil {
		h.customLogger.Error().Msg(err.Error())
		component = components.Notification("Server failed to respond, try again", "fail")
		return tadaptor.Render(c, component, http.StatusBadRequest)
	}

	component = components.Notification("Vacancy successfully created", "success")
	return tadaptor.Render(c, component, http.StatusOK)
}
