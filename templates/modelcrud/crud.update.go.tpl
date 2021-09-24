{{- $alias := .Aliases.Table .Table.Name -}}
{{- $orig_tbl_name := .Table.Name -}}
{{- $canSoftDelete := .Table.CanSoftDelete -}}
{{- $soft := and .AddSoftDeletes $canSoftDelete }}

type APIUpdate{{ $alias.UpSingular }}Request struct {
{{- range $column := .Table.Columns -}}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{- $orig_col_name := $column.Name -}}
    {{- range $column.Comment | splitLines -}} // {{ . }}
    {{end -}}
    {{if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{ else if eq $orig_col_name "created_at" }}
    {{ else if eq $orig_col_name "updated_at" }}
    {{ else if eq $orig_col_name "deleted_at" }}
    {{ else if eq $orig_col_name "id" }}
    {{else -}}
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
        {{- else if eq $column.Type "null.Float64" }}
            {{$colAlias}} *float64 `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "types.StringArray" }}
            {{$colAlias}} []string `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else -}}
            {{$colAlias}} *{{ $column.Type }} `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- end }}
    {{end -}}
{{end -}}
}

type Update{{ $alias.UpSingular }}Request struct {
{{- range $column := .Table.Columns -}}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{- $orig_col_name := $column.Name -}}
    {{- range $column.Comment | splitLines -}} // {{ . }}
    {{end -}}
    {{if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{ else if eq $orig_col_name "created_at" }}
    {{ else if eq $orig_col_name "updated_at" }}
    {{ else if eq $orig_col_name "deleted_at" }}
    {{ else if eq $orig_col_name "id" }}
    {{else -}}
        {{$colAlias}} *{{$column.Type}} `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}},omitempty" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
    {{end -}}
{{end -}}
}

// Update{{ $alias.UpSingular }}ByID godoc
// @Summary Updates a single {{ $alias.UpSingular }} entity based on their id
// @Produce json
// @Accept json
// @Param req body APIUpdate{{ $alias.UpSingular }}Request true "Update parameters"
// @Param id path string true "{{ $alias.UpSingular }} id"
// @Success 200 {object} API{{$alias.UpSingular}}
// @Router /crud/{{ $alias.DownPlural }}/:id [put]
func (*GeneratedCrudController) Update{{ $alias.UpSingular }}ByID(c *gin.Context) {
id := c.Param("id")

if id == "" {
api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
return
}

pks := strings.Split(id, ";")

if len(pks) != {{ len .Table.PKey.Columns }} {
api.NewAPIError("invalid-id", http.StatusBadRequest, fmt.Sprintf("Expected {{ len .Table.PKey.Columns }} ids, got %d", len(pks))).Respond(c)
return
}

qms := []qm.QueryMod{}

{{ range $i, $pk := .Table.PKey.Columns }}
    qms = append(qms, qm.Where("{{ $pk }} = ?", pks[{{ $i }}]))
{{ end }}

var updateReq Update{{ $alias.UpSingular }}Request

err := c.BindJSON(&updateReq)
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

qms = append(qms, qm.For("UPDATE"))

existing{{ $alias.UpSingular }}, err := models.{{ $alias.UpPlural }}(qms...).OneG(c.Request.Context())
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

{{- range $column := .Table.Columns -}}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{- $orig_col_name := $column.Name -}}
    {{- range $column.Comment | splitLines -}} // {{ . }}
    {{end -}}
    {{if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{ else if eq $orig_col_name "created_at" }}
    {{ else if eq $orig_col_name "updated_at" }}
    {{ else if eq $orig_col_name "deleted_at" }}
    {{ else if eq $orig_col_name "id" }}
    {{ else }}
        if updateReq.{{$colAlias}} != nil {
        existing{{ $alias.UpSingular}}.{{$colAlias}} = *updateReq.{{$colAlias}}
        }
    {{end}}
{{end -}}

_, err = existing{{ $alias.UpSingular }}.UpdateG(c.Request.Context(), boil.Infer())
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

c.JSON(http.StatusOK, existing{{ $alias.UpSingular }})
}