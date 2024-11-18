package system

type Page struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"` // desc降序 asc升序
}
