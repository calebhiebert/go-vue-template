{{- $alias := .Aliases.Table .Table.Name -}}
{{- $orig_tbl_name := .Table.Name -}}
{{- $canSoftDelete := .Table.CanSoftDelete -}}
{{- $soft := and .AddSoftDeletes $canSoftDelete -}}

{{ define "gql_type" }}
    {{- if eq .Type "null.String" -}}
        CString
    {{- else if eq .Type "types.String" -}}
        String
    {{- else if eq .Type "string" -}}
        String
    {{- else if eq .Type "null.Int" -}}
        Int
    {{- else if eq .Type "int" -}}
        Int
    {{- else if eq .Type "null.Bool" -}}
        CBool
    {{- else if eq .Type "bool" -}}
        Boolean
    {{- else if eq .Type "float64" -}}
        Float
    {{- else if eq .Type "null.JSON" -}}
        Any
    {{- else if eq .Type "types.JSON" -}}
        Any
    {{- else if eq .Type "null.Time" -}}
        Time
    {{- else if eq .Type "time.Time" -}}
        Time
    {{- else if eq .Type "null.Float64" -}}
        Float
    {{- else if eq .Type "types.StringArray" -}}
        [String!]
    {{- else -}}
        # TODO Missing type: {{ .Type }}
    {{- end -}}
{{ end }}

type {{$alias.UpSingular}} {
{{- range $column := .Table.Columns -}}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{- $orig_col_name := $column.Name -}}
    {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{- else -}}
        {{ if eq $orig_col_name "id" -}}
            id: ID!
        {{- else -}}
            {{ $orig_col_name }}: {{ template "gql_type" $column}}{{ if not $column.Nullable }}!{{ end }}
        {{- end }}
    {{ end -}}
{{- end -}}
}

input Update{{$alias.UpSingular}} {
{{- range $column := .Table.Columns -}}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{- $orig_col_name := $column.Name -}}
    {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{- else -}}
        {{- if eq $orig_col_name "id" -}}
        {{- else if eq $orig_col_name "created_at" -}}
        {{- else if eq $orig_col_name "updated_at" -}}
        {{- else if eq $orig_col_name "deleted_at" -}}
        {{- else if eq $orig_col_name "created_by_id" -}}
        {{- else if eq $column.Type "null.String" -}}
            {{ $orig_col_name }}: CString
        {{- else if eq $column.Type "types.String" -}}
            {{ $orig_col_name }}: String
        {{- else if eq $column.Type "string" -}}
            {{ $orig_col_name }}: String
        {{- else if eq $column.Type "null.Int" -}}
            {{ $orig_col_name }}: Int
        {{- else if eq $column.Type "int" -}}
            {{ $orig_col_name }}: Int
        {{- else if eq $column.Type "float64" -}}
            {{ $orig_col_name }}: Float
        {{- else if eq $column.Type "null.Bool" -}}
            {{ $orig_col_name }}: CBool
        {{- else if eq $column.Type "bool" -}}
            {{ $orig_col_name }}: Boolean
        {{- else if eq $column.Type "float64" -}}
            {{ $orig_col_name }}: Float
        {{- else if eq $column.Type "null.Float64" -}}
            {{ $orig_col_name }}: Float
        {{- else if eq $column.Type "null.JSON" -}}
        {{- else if eq $column.Type "types.JSON" -}}
        {{- else if eq $column.Type "null.Time" -}}
            {{ $orig_col_name }}: Time
        {{- else if eq $column.Type "time.Time" -}}
            {{ $orig_col_name }}: Time
        {{- else if eq $column.Type "types.StringArray" -}}
            {{ $orig_col_name }}: [String]
        {{- else -}}
            # TODO Missing type: {{ $orig_col_name }} {{ $column.Type }}
        {{- end }}
    {{ end -}}
{{- end -}}
}

input Create{{$alias.UpSingular}} {
{{- range $column := .Table.Columns -}}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{- $orig_col_name := $column.Name -}}
    {{- $defaultLen := len $column.Default -}}
    {{- $hasDefault := gt $defaultLen 0 -}}
    {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{- else -}}
        {{- if eq $orig_col_name "id" -}}
        {{- else if eq $orig_col_name "created_at" -}}
        {{- else if eq $orig_col_name "updated_at" -}}
        {{- else if eq $orig_col_name "deleted_at" -}}
        {{- else if eq $orig_col_name "created_by_id" -}}
        {{- else -}}
            {{ $orig_col_name }}: {{ template "gql_type" $column}}{{ if or $column.Nullable $hasDefault }}{{else}}!{{ end }}
        {{- end }}
    {{ end -}}
{{- end -}}

}