package projects

type Project struct {
	Id          int    `json:"id"`
	Uuid        string `json:"uuid"`
	ProjectId   string `json:"shorthand"`
	Title       string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Favourite   bool   `json:"favorite"`
	Completed   bool   `json:"released"`
	InProgress  bool   `json:"progress"`
	Github      string `json:"github"`
	URL         string `json:"url"`
	Category    int    `json:"category"`
}
