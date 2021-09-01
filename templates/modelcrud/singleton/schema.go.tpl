
var AdminInfo api.AdminModelInfo = api.AdminModelInfo{
    Models: map[string]api.AdminModel{
        {{- range $table := .Tables }}
            "{{ $table.Name }}": {{ titleCase $table.Name }}Admin,
        {{- end }}
    },
}