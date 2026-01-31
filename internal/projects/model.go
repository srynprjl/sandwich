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

type ProjectFields struct {
	Field []any
}

func (p *ProjectFields) Init(data *Project) {
	p.Field = []any{&data.Id, &data.Uuid, &data.ProjectId, &data.Title, &data.Description, &data.Path, &data.Favourite, &data.Completed, &data.InProgress, &data.Github, &data.URL, &data.Category}
}
