package category

type Category struct {
	Id        int    `json:"id"`
	Title     string `json:"name"`
	Shorthand string `json:"shorthand"`
	//Description string
}
