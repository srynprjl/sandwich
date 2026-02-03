package logic

type Project struct {
	Id          int    `json:"id" yaml:"id"`
	Uuid        string `json:"uuid" yaml:"uuid"`
	ProjectId   string `json:"shorthand" yaml:"shorthand"`
	Title       string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Path        string `json:"path" yaml:"path"`
	Favourite   bool   `json:"favorite" yaml:"favorite"`
	Completed   bool   `json:"released" yaml:"released"`
	InProgress  bool   `json:"progress" yaml:"inprogress"`
	Github      string `json:"github" yaml:"github"`
	URL         string `json:"url" yaml:"url"`
	Category    int    `json:"category" yaml:"category"`
}

type Category struct {
	Id          int    `json:"id" yaml:"id"`
	Uuid        string `json:"uuid" yaml:"uuid"`
	Title       string `json:"name" yaml:"name"`
	Shorthand   string `json:"shorthand" yaml:"shorthand"`
	Description string `json:"description" yaml:"description"`
}
