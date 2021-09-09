{{- $alias := .Aliases.Table .Table.Name -}}
{{- $orig_tbl_name := .Table.Name -}}
{{- $canSoftDelete := .Table.CanSoftDelete -}}
{{- $soft := and .AddSoftDeletes $canSoftDelete }}

type Generate{{ $alias.UpSingular }}StringID func(*Create{{ $alias.UpSingular }}Request) string

var Generate{{ $alias.UpSingular }}ID Generate{{ $alias.UpSingular }}StringID

type APICreate{{ $alias.UpSingular }}Request struct {
{{- range $column := .Table.Columns -}}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{- $orig_col_name := $column.Name -}}
    {{- range $column.Comment | splitLines -}} // {{ . }}
    {{end -}}
    {{ if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{- else if eq $orig_col_name "created_at" }}
    {{- else if eq $orig_col_name "updated_at" }}
    {{- else if eq $orig_col_name "deleted_at" }}
    {{- else if eq $orig_col_name "id" }}
    {{- else -}}
        {{- if eq $column.Type "null.String" }}
            {{$colAlias}} *string `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "types.String" }}
            {{$colAlias}} *string `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "null.Int" }}
            {{$colAlias}} *int `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "null.Bool" }}
            {{$colAlias}} *bool `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "null.JSON" }}
            {{$colAlias}} map[string]interface{} `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "types.JSON" }}
            {{$colAlias}} map[string]interface{} `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "null.Time" }}
            {{$colAlias}} *time.Time `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "types.StringArray" }}
            {{$colAlias}} []string `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else -}}
            {{$colAlias}} *{{ $column.Type }} `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- end }}
    {{end -}}
{{end -}}
}

type Create{{ $alias.UpSingular }}Request struct {
{{- range $column := .Table.Columns -}}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{- $orig_col_name := $column.Name -}}
    {{- $defaultLen := len $column.Default -}}
    {{- $hasDefault := gt $defaultLen 0 -}}
    {{- range $column.Comment | splitLines -}} // {{ . }}
    {{end -}}
    {{if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{ else if eq $orig_col_name "created_at" }}
    {{ else if eq $orig_col_name "updated_at" }}
    {{ else if eq $orig_col_name "deleted_at" }}
    {{ else if eq $orig_col_name "id" }}
        {{$colAlias}} *{{$column.Type}} `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}},omitempty" toml:"{{$column.Name}}" yaml:"{{$column.Name}},omitempty"`
    {{else -}}
        {{$colAlias}} {{ if or $column.Nullable $hasDefault }}*{{$column.Type}}{{ else }}{{ $column.Type }}{{ end }} `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}},omitempty" toml:"{{$column.Name}}" {{ if or $column.Nullable $hasDefault }}{{else}}binding:"required"{{ end }} yaml:"{{$column.Name}}{{if or $column.Nullable $hasDefault }},omitempty{{end}}"`
    {{end -}}
{{end -}}
}

// Create{{ $alias.UpSingular }} godoc
// @Summary Creates a new {{ $alias.UpSingular }}
// @Produce json
// @Accept json
// @Param req body APICreate{{ $alias.UpSingular }}Request true "Creation parameters"
// @Success 200 {object} API{{ $alias.UpSingular }}
// @Router /crud/{{ $alias.DownPlural }}  [post]
func (*GeneratedCrudController) Create{{ $alias.UpSingular }} (c *gin.Context) {
var createReq Create{{ $alias.UpSingular }}Request

err := c.BindJSON(&createReq)
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

new{{ $alias.UpSingular }} := models.{{ $alias.UpSingular }} {}

{{- range $column := .Table.Columns -}}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{- $orig_col_name := $column.Name -}}
    {{- $defaultLen := len $column.Default -}}
    {{- $hasDefault := gt $defaultLen 0 -}}
    {{if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{ else if eq $orig_col_name "created_at" }}
    {{ else if eq $orig_col_name "updated_at" }}
    {{ else if eq $orig_col_name "deleted_at" }}
    {{ else if eq $orig_col_name "id" }}
        {{ if and (eq $orig_col_name "id") (eq $column.DBType "uuid") }}
            new{{ $alias.UpSingular }}.{{ $colAlias }} = uuid.Must(uuid.NewV4()).String()
        {{ else }}
            if createReq.{{ $colAlias }} != nil {
                new{{ $alias.UpSingular }}.{{ $colAlias }} = *createReq.{{ $colAlias }}
            } else {
                new{{ $alias.UpSingular }}.{{ $colAlias }} = Generate{{ $alias.UpSingular }}ID(&createReq)
            }
        {{ end }}
    {{else -}}
        {{ if or $column.Nullable $hasDefault }}
            if createReq.{{ $colAlias }} != nil {
            new{{ $alias.UpSingular }}.{{ $colAlias }} = *createReq.{{ $colAlias }}
            }
        {{ else }}
            new{{ $alias.UpSingular }}.{{ $colAlias }} = createReq.{{ $colAlias }}
        {{ end }}
    {{end -}}
{{end -}}

err = new{{ $alias.UpSingular }}.InsertG(c.Request.Context(), boil.Infer())
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

c.JSON(http.StatusOK, new{{ $alias.UpSingular }} )
}