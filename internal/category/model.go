package category

type Category struct {
	Id          int    `json:"id" yaml:"id"`
	Uuid        string `json:"uuid" yaml:"uuid"`
	Title       string `json:"name" yaml:"name"`
	Shorthand   string `json:"shorthand" yaml:"shorthand"`
	Description string `json:"description" yaml:"description"`
}
