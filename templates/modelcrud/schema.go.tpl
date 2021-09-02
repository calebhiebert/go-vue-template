{{- $alias := .Aliases.Table .Table.Name -}}
{{- $orig_tbl_name := .Table.Name -}}
{{- $canSoftDelete := .Table.CanSoftDelete -}}
{{- $soft := and .AddSoftDeletes $canSoftDelete }}

var {{ $alias.UpPlural }}Admin = api.AdminModel{
Name: "{{ $alias.UpPlural }}",
CanSoftDelete: {{ if $soft }}true{{ else }}false{{end}},
URLName: "{{ $alias.DownPlural }}",
DataName: "{{ $orig_tbl_name }}",
Fields: []*api.AdminModelField{
{{- range $field := .Table.Columns }}
    {{- $colAlias := $alias.Column $field.Name -}}
    {{- $orig_col_name := $field.Name -}}
    {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{- else }}
        &api.AdminModelField{
        ID: "{{ $orig_col_name }}",
        Name: "{{ $colAlias }}",
        Nullable: {{ $field.Nullable }},
        Config: api.NewDefaultAdminModelFieldConfig(),
        Type:
        {{- if eq $field.Type "null.String" -}}
            "string"
        {{- else if eq $field.Type "types.String" -}}
            "string"
        {{- else if eq $field.Type "null.JSON" -}}
            "json"
        {{- else if eq $field.Type "types.JSON" -}}
            "json"
        {{- else if eq $field.Type "null.Time" -}}
            "time"
        {{- else if eq $field.Type "time.Time" -}}
            "time"
        {{- else if eq $field.Type "types.StringArray" -}}
            "array"
        {{- else -}}
            "{{ $field.Type }}"
        {{- end -}},
        },
    {{- end }}
{{- end }}
},
}

type {{ titleCase $.Table.Name }}ModelConfigType struct {
{{ range $idx, $field := .Table.Columns }}
    {{- $colAlias := $alias.Column $field.Name -}}
    {{- $orig_col_name := $field.Name -}}
    {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{- else }}
        {{ $colAlias }} api.AdminModelFieldConfig
    {{- end }}
{{ end }}
}

var {{ titleCase $.Table.Name }}ModelConfig = {{ titleCase $.Table.Name }}ModelConfigType{
{{ range $idx, $field := .Table.Columns }}
    {{- $colAlias := $alias.Column $field.Name -}}
    {{- $orig_col_name := $field.Name -}}
    {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{- else }}
        {{ $colAlias }}: api.AdminModelFieldConfig{
        ShowOnGraph: true,
        },
    {{- end }}
{{ end }}
}

func (c {{ titleCase $.Table.Name }}ModelConfigType) Apply() {
{{ range $idx, $field := .Table.Columns }}
    {{- $colAlias := $alias.Column $field.Name -}}
    {{- $orig_col_name := $field.Name -}}
    {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{- else }}
        {{ $alias.UpPlural }}Admin.Fields[{{$idx}}].Config = c.{{ $colAlias }}
    {{- end }}
{{ end }}
}