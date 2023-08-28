package request

type Page struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}
