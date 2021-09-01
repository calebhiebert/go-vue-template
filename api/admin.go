package api

type AdminModelInfo struct {
	Models map[string]AdminModel `json:"models"`
}

type AdminModelConfig struct {
}

type AdminModel struct {
	Name          string            `json:"name"`
	Fields        []AdminModelField `json:"fields"`
	Config        AdminModelConfig  `json:"config"`
	CanSoftDelete bool              `json:"can_soft_delete"`
}

type AdminModelFieldConfig struct {
	ShowOnGraph bool `json:"show_on_graph"`
}

func NewDefaultAdminModelFieldConfig() AdminModelFieldConfig {
	return AdminModelFieldConfig{
		ShowOnGraph: true,
	}
}

type AdminModelField struct {
	ID       string                `json:"id"`
	Name     string                `json:"name"`
	Type     string                `json:"type"`
	Nullable bool                  `json:"nullable"`
	Config   AdminModelFieldConfig `json:"config"`
}
