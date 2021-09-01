{{- $alias := .Aliases.Table .Table.Name -}}
{{- $orig_tbl_name := .Table.Name -}}
{{- $canSoftDelete := .Table.CanSoftDelete -}}
{{- $soft := and .AddSoftDeletes $canSoftDelete }}

type API{{$alias.UpSingular}} struct {
{{- range $column := .Table.Columns -}}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{- $orig_col_name := $column.Name -}}
    {{- range $column.Comment | splitLines -}} // {{ . }}
    {{end -}}
    {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{- else -}}
        {{- if eq $column.Type "null.String" }}
            {{$colAlias}} *string `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "types.String" }}
            {{$colAlias}} string `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "null.JSON" }}
            {{$colAlias}} map[string]interface{} `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "types.JSON" }}
            {{$colAlias}} map[string]interface{} `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "null.Time" }}
            {{$colAlias}} *time.Time `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "types.StringArray" }}
            {{$colAlias}} []string `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else -}}
            {{$colAlias}} {{ $column.Type }} `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- end }}
    {{end -}}
{{end -}}
}


type Get{{ $alias.UpPlural }}Response struct {
{{ $alias.UpPlural }} models.{{ $alias.UpSingular }}Slice `json:"{{ $orig_tbl_name }}"`
Total int64            `json:"total"`
NextOffset int64 `json:"next_offset"`
}

type APIGet{{ $alias.UpPlural }}Response struct {
{{ $alias.UpPlural }} []API{{ $alias.UpSingular }} `json:"{{ $orig_tbl_name }}"`
Total int64            `json:"total"`
NextOffset int64 `json:"next_offset"`
}

// Get{{ $alias.UpSingular }}ByID godoc
// @Summary Gets a single {{ $alias.UpSingular }} entity by their id
// @Produce json
// @Success 200 {object} APIGet{{ $alias.UpPlural }}Response
// @Param id path string true "{{ $alias.UpSingular }} id"
// @Router /crud/{{ $alias.DownPlural }}/:id [get]
func (*GeneratedCrudController) Get{{ $alias.UpSingular }}ByID(c *gin.Context)  {
id := c.Param("id")

if id == "" {
api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
return
}

{{ $alias.UpSingular }}, err := models.{{ $alias.UpPlural }}(qm.Where("id = ?", id)).OneG(c.Request.Context())
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

c.JSON(http.StatusOK, {{ $alias.UpSingular }})
}

// Get{{ $alias.UpPlural }} godoc
// @Summary Gets a list for all entities of the {{ $alias.UpSingular }} type
// @Produce json
// @Success 200 {object} API{{$alias.UpSingular}}
{{- if $soft }}
    // @Param withDeleted query string false "Include deleted {{ $alias.DownPlural }} in the results"
{{- end }}
// @Router /crud/{{ $alias.DownPlural }} [get]
func (*GeneratedCrudController) Get{{ $alias.UpPlural }}(c *gin.Context) {
limit, offset := api.ExtractLimitOffset(c)

count, err := models.{{ $alias.UpPlural }}().CountG(c.Request.Context())
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

queryMods := []qm.QueryMod{
qm.Limit(limit),
qm.Offset(offset),
}

{{ if $soft }}
    withDeleted := c.Query("withDeleted") == "true"

    if withDeleted {
    queryMods = append(queryMods, qm.WithDeleted())
    }
{{ end }}

{{ $alias.DownPlural }}, err := models.{{ $alias.UpPlural }}(queryMods...).AllG(c.Request.Context())
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

if {{ $alias.DownPlural }} == nil {
{{ $alias.DownPlural }} = models.{{ $alias.UpSingular }}Slice{}
}

c.JSON(http.StatusOK, Get{{ $alias.UpPlural }}Response{
{{ $alias.UpPlural }}: {{ $alias.DownPlural }},
Total: count,
NextOffset: int64(offset + limit),
})
}

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

var updateReq Update{{ $alias.UpSingular }}Request

err := c.BindJSON(&updateReq)
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

existing{{ $alias.UpSingular }}, err := models.{{ $alias.UpPlural }}(qm.Where("id = ?", id), qm.For("UPDATE")).OneG(c.Request.Context())
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
    {{else -}}
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

// Delete{{ $alias.UpSingular }}ByID godoc
// @Summary Soft deletes a single {{ $alias.UpSingular }} entity based on their id
// @Produce json
// @Success 200 {object} API{{$alias.UpSingular}}
// @Param id path string true "{{ $alias.UpSingular }} id"
{{- if $soft }}
    // @Param hardDelete query string false "Hard delete {{ $alias.DownSingular }}"
{{- end }}
// @Router /crud/{{ $alias.DownPlural }}/:id [delete]
func (*GeneratedCrudController) Delete{{ $alias.UpSingular }}ByID(c *gin.Context) {
id := c.Param("id")

if id == "" {
api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
return
}

{{ if $soft }}
    hardDelete := c.Query("hardDelete") == "true"
{{ end }}

existing{{ $alias.UpSingular }}, err := models.{{ $alias.UpPlural }}(qm.Where("id = ?", id)).OneG(c.Request.Context())
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

_, err = existing{{ $alias.UpSingular }}.DeleteG(c.Request.Context(){{if $soft}}, hardDelete{{end}})
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

c.JSON(http.StatusOK, existing{{ $alias.UpSingular }})
}


func (gcc *GeneratedCrudController) Register{{ $alias.UpPlural }}(rg *gin.RouterGroup) {
rg.GET("/{{ $alias.DownPlural }}/:id", gcc.Get{{ $alias.UpSingular }}ByID)
rg.GET("/{{ $alias.DownPlural }}", gcc.Get{{ $alias.UpPlural }})
rg.PUT("/{{ $alias.DownPlural }}/:id", gcc.Update{{ $alias.UpSingular }}ByID)
rg.DELETE("/{{ $alias.DownPlural }}/:id", gcc.Delete{{ $alias.UpSingular }}ByID)
}