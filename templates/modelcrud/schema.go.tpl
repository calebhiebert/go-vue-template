{{- $alias := .Aliases.Table .Table.Name -}}
{{- $orig_tbl_name := .Table.Name -}}
{{- $canSoftDelete := .Table.CanSoftDelete -}}
{{- $soft := and .AddSoftDeletes $canSoftDelete }}

{{- define "filter_operations" -}}
    []string{
    "eq", "null",
    {{- if eq .DBType "uuid" -}}
    {{- else if eq .DBType "jsonb" -}}
    {{- else if eq .Type "string" -}}
        "cont",
    {{- else if eq .Type "null.String" -}}
        "cont",
    {{- else if eq .Type "int" -}}
        "gt", "lt", "gte", "lte",
    {{- else if eq .Type "null.Int" -}}
        "gt", "lt", "gte", "lte",
    {{- else if eq .Type "null.Bool" -}}
    {{- else if eq .Type "time.Time" -}}
        "gt", "lt", "gte", "lte",
    {{- else if eq .Type "null.Time" -}}
        "gt", "lt", "gte", "lte",
    {{- else if eq .Type "null.Float64" -}}
        "gt", "lt", "gte", "lte",
    {{- end -}}
    }
{{- end }}

var {{ $alias.UpPlural }}Admin = api.AdminModel{
Name: "{{ $alias.UpPlural }}",
NameSingular: "{{ $alias.UpSingular }}",
CanSoftDelete: {{ if $soft }}true{{ else }}false{{end}},
URLName: "{{ $alias.DownPlural }}",
DataName: "{{ $orig_tbl_name }}",
Fields: []*api.AdminModelField{
{{- range $field := .Table.Columns }}
    {{- $colAlias := $alias.Column $field.Name -}}
    {{- $orig_col_name := $field.Name -}}
    {{- $defaultLen := len $field.Default -}}
    {{- $hasDefault := gt $defaultLen 0 -}}
    {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{- else }}
        &api.AdminModelField{
        ID: "{{ $orig_col_name }}",
        Name: "{{ $colAlias }}",
        Nullable: {{ $field.Nullable }},
        Required: {{ if or $field.Nullable $hasDefault }}false{{else}}true{{ end }},
        FilterOperations: {{ template "filter_operations" $field }},
        ForeignFields: []api.AdminModelForeignField{
        {{- range $fkey := $.Table.FKeys -}}
            {{ if eq $fkey.Column $orig_col_name}}
                {{- $ltable := $.Aliases.Table $fkey.Table -}}
                {{- $ftable := $.Aliases.Table $fkey.ForeignTable -}}
                {{- $rel := $ltable.Relationship $fkey.Name -}}
                {{- $canSoftDelete := (getTable $.Tables $fkey.ForeignTable).CanSoftDelete }}
                {
                Model: "{{ $fkey.ForeignTable }}",
                Field: "{{ $fkey.ForeignColumn }}",
                Nullable: {{ $fkey.Nullable }},
                Unique: {{ $fkey.Unique }},
                },
            {{ end }}
        {{- end -}}
        },
        Editable:
        {{- if and (eq $orig_col_name "id") (eq $field.DBType "uuid") -}}
            false
        {{- else if eq $orig_col_name "id" -}}
            true
        {{- else if eq $orig_col_name "created_at" -}}
            false
        {{- else if eq $orig_col_name "updated_at" -}}
            false
        {{- else if eq $orig_col_name "deleted_at" -}}
            false
        {{- else -}}
            true
        {{- end -}},
        Config: api.AdminModelFieldConfig{
        ShowOnGraph: {{ if or (eq $orig_col_name "deleted_at") (eq $field.DBType "jsonb") }}false{{ else }}true{{ end }},
        Editable: true,
        IsEmail: {{- if eq $orig_col_name "email" -}}true{{- else -}}false{{- end -}},
        MiniSearchable: {{ if eq $orig_col_name "name" }}true{{ else }}false{{ end }},
        },
        Type:
        {{- if eq $field.DBType "uuid" -}}
            "uuid"
        {{- else if eq $field.Type "null.String" -}}
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
        {{- else if eq $field.Type "null.Float64" -}}
            "float"
        {{- else if eq $field.Type "types.StringArray" -}}
            "array"
        {{- else if eq $field.Type "null.Int" -}}
            "int"
        {{- else if eq $field.Type "null.Bool" -}}
            "bool"
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
        ShowOnGraph: {{ if or (eq $orig_col_name "deleted_at") (eq $field.DBType "jsonb") }}false{{ else }}true{{ end }},
        Editable: true,
        IsEmail: {{- if eq $orig_tbl_name "email" -}}true{{- else -}}false{{- end -}},
        },
    {{- end -}}
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