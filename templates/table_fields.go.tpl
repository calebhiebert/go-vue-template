{{- $alias := .Aliases.Table .Table.Name -}}

type {{ $alias.UpSingular }}ColumnData struct {
{{ range $column := .Table.Columns }}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{ $colAlias }} TableField `json:"{{ $column.Name }}"`
{{ end }}
}

var {{ $alias.UpSingular }}ColumnsData {{ $alias.UpSingular }}ColumnData = {{ $alias.UpSingular }}ColumnData{
{{ range $column := .Table.Columns }}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{ $colAlias }}: TableField{
    Name: "{{ $column.Name }}",
    Type: "{{ $column.Type }}",
    {{- if ne $column.Default "" }}
        HasDefault: true,
    {{- else }}
        HasDefault: false,
    {{- end }}
    Nullable: {{ $column.Nullable }},
    },
{{ end }}
}