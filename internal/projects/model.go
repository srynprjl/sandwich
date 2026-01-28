package projects

type Project struct {
	Id          int    `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Favourite   bool   `json:"favorite"`
	Completed   bool   `json:"completed"`
	Category    int    `json:"category"`
	//Working     bool
	//ProjectId string
	//GitHubURL string
}

type ProjectFields struct {
	Field []any
}

func (p *ProjectFields) Init(data *Project) {
	p.Field = []any{&data.Id, &data.Title, &data.Description, &data.Completed, &data.Favourite, &data.Path, &data.Category}
}
