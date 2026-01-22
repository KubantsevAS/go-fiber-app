package vacancy

import "time"

type VacancyCreateForm struct {
	Email        string
	JobTitle     string
	Company      string
	CompanyScope string
	Salary       string
	Location     string
}

type Vacancy struct {
	Id        int       `db:"id"`
	Email     string    `db:"email"`
	Role      string    `db:"role"`
	Company   string    `db:"company"`
	Salary    string    `db:"salary"`
	Type      string    `db:"type"`
	Location  string    `db:"location"`
	CreatedAt time.Time `db:"createdat"`
}
