package api

type AdminModelInfo struct {
	Models map[string]AdminModel `json:"models"`
}

type AdminModelConfig struct {
}

type AdminModel struct {
	Name          string             `json:"name"`
	NameSingular  string             `json:"name_singular"`
	Fields        []*AdminModelField `json:"fields"`
	Config        AdminModelConfig   `json:"config"`
	CanSoftDelete bool               `json:"can_soft_delete"`
	URLName       string             `json:"url_name"`
	DataName      string             `json:"data_name"`
}

type AdminModelFieldConfig struct {
	ShowOnGraph bool `json:"show_on_graph"`
	Editable    bool `json:"editable"`
	IsEmail     bool `json:"is_email"`
}

func NewDefaultAdminModelFieldConfig() AdminModelFieldConfig {
	return AdminModelFieldConfig{
		ShowOnGraph: true,
		Editable:    true,
		IsEmail:     false,
	}
}

type AdminModelField struct {
	ID               string                   `json:"id"`
	Name             string                   `json:"name"`
	Type             string                   `json:"type"`
	Nullable         bool                     `json:"nullable"`
	Required         bool                     `json:"required"`
	Editable         bool                     `json:"editable"`
	Config           AdminModelFieldConfig    `json:"config"`
	FilterOperations []string                 `json:"filter_operations"`
	ForeignFields    []AdminModelForeignField `json:"foreign_fields"`
}

type AdminModelForeignField struct {
	Model    string `json:"model"`
	Field    string `json:"field"`
	Nullable bool   `json:"nullable"`
	Unique   bool   `json:"unique"`
}
