package category

type Category struct {
	Id          int    `json:"id"`
	Uuid        string `json:"uuid"`
	Title       string `json:"name"`
	Shorthand   string `json:"shorthand"`
	Description string `json:"description"`
}
