package category

type Category struct {
	ID          int    `json:"id" yaml:"id"`
	UUID        string `json:"uuid" yaml:"uuid"`
	Title       string `json:"name" yaml:"name"`
	UID         string `json:"shorthand" yaml:"shorthand"`
	Description string `json:"description" yaml:"description"`
}
