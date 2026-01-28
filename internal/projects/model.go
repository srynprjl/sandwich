package projects

type Project struct {
	Id          int
	Title       string
	Description string
	Path        string
	Favourite   bool
	Completed   bool
	Category    int
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
