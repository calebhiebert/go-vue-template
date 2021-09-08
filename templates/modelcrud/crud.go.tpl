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
    {{- else }}
        // {{ $column.DBType }}
        {{ if eq $column.Type "null.String" }}
            {{$colAlias}} *string `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "types.String" }}
            {{$colAlias}} string `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
        {{- else if eq $column.Type "null.Int" }}
            {{$colAlias}} *int `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
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

{{ define "numeric_query_operators" }}
    case "{{ . }}.gt":
    queryMods = append(queryMods, qm.Where("{{ . }} > ?", v[0]))
    case "{{ . }}.lt":
    queryMods = append(queryMods, qm.Where("{{ . }} < ?", v[0]))
    case "{{ . }}.gte":
    queryMods = append(queryMods, qm.Where("{{ . }} >= ?", v[0]))
    case "{{ . }}.lte":
    queryMods = append(queryMods, qm.Where("{{ . }} <= ?", v[0]))
{{ end }}

// Get{{ $alias.UpPlural }} godoc
// @Summary Gets a list for all entities of the {{ $alias.UpSingular }} type
// @Produce json
// @Success 200 {object} API{{$alias.UpSingular}}
{{- if $soft }}
    // @Param withDeleted query string false "Include deleted {{ $alias.DownPlural }} in the results"
{{- end }}
{{- range $column := .Table.Columns -}}
    {{- $orig_col_name := $column.Name -}}
    {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{- else }}
        // @Param sort.{{ $orig_col_name }} query string false "Sort by {{ $orig_col_name }}. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
    {{- end }}
{{- end }}
// @Router /crud/{{ $alias.DownPlural }} [get]
func (*GeneratedCrudController) Get{{ $alias.UpPlural }}(c *gin.Context) {
queryMods := []qm.QueryMod{}

{{ if $soft }}
    withDeleted := c.Query("withDeleted") == "true"

    if withDeleted {
    queryMods = append(queryMods, qm.WithDeleted())
    }
{{ end }}

var orderBy []string

for q, v := range c.Request.URL.Query() {
    sortDirection := "ASC"

    if v[0] == "DESC" || v[0] == "desc" {
        sortDirection = "DESC"
    }

    switch q {
        {{- range $column := .Table.Columns -}}
            {{- $orig_col_name := $column.Name -}}
            {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
            {{- else }}
                case "sort.{{ $orig_col_name }}":
                    orderBy = append(orderBy, "{{$orig_col_name}} " + sortDirection)
                case "{{ $orig_col_name }}.eq":
                    queryMods = append(queryMods, qm.Where("{{ $orig_col_name }} = ?", v[0]))
                {{ if eq $column.Type "int" }}
                    {{ template "numeric_query_operators" $orig_col_name }}
                {{ else if eq $column.Type "time.Time" }}
                    {{ template "numeric_query_operators" $orig_col_name }}
                {{ else if eq $column.Type "null.Int" }}
                    {{ template "numeric_query_operators" $orig_col_name }}
                {{ else if eq $column.Type "null.Time" }}
                    {{ template "numeric_query_operators" $orig_col_name }}
                {{ end }}
                {{ if eq $column.DBType "uuid" }}
                {{ else if eq $column.Type "string" }}
                case "{{ $orig_col_name }}.cont":
                    {{ $orig_col_name }}SearchString := fmt.Sprintf("%%%s%%", strings.ReplaceAll(v[0], "%", "\\%"))
                    queryMods = append(queryMods, qm.Where("{{ $orig_col_name }} ILIKE ?", {{ $orig_col_name }}SearchString))
                {{ else if eq $column.Type "null.String"}}
                case "{{ $orig_col_name }}.cont":
                    {{ $orig_col_name }}SearchString := fmt.Sprintf("%%%s%%", strings.ReplaceAll(v[0], "%", "\\%"))
                    queryMods = append(queryMods, qm.Where("{{ $orig_col_name }} ILIKE ?", {{ $orig_col_name }}SearchString))
                {{ end }}
            {{- end }}
        {{- end -}}
    }
}

count, err := models.{{ $alias.UpPlural }}(queryMods...).CountG(c.Request.Context())
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

limit, offset := api.ExtractLimitOffset(c)

queryMods = append(queryMods, qm.Limit(limit), qm.Offset(offset))

if len(orderBy) > 0 {
    queryMods = append(queryMods, qm.OrderBy(strings.Join(orderBy, ", ")))
} {{- $colNames := .Table.Columns | columnNames -}}
{{- if containsAny $colNames "created_at" -}}
    else {
    queryMods = append(queryMods, qm.OrderBy("created_at DESC"))
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
        {{- else if eq $column.Type "null.Int" }}
            {{$colAlias}} *int `{{generateTags $.Tags $column.Name}}boil:"{{$column.Name}}" json:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}" toml:"{{$column.Name}}" yaml:"{{$column.Name}}{{if $column.Nullable}},omitempty{{end}}"`
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

// BulkDelete{{ $alias.UpPlural }}ByIDs godoc
// @Summary Soft deletes a range of {{ $alias.DownPlural }} by their ids
// @Produce json
// @Success 200 {object} DeletedCount
// @Param req body IDList true "List of ids to delete"
// @Param hardDelete query string false "Hard delete {{ $alias.DownSingular }}"
// @Router /crud/{{ $alias.DownPlural }} [delete]
func (*GeneratedCrudController) BulkDelete{{ $alias.UpPlural }}ByIDs(c *gin.Context) {

var ids IDList

err := c.BindJSON(&ids)
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

{{ if $soft }}
    hardDelete := c.Query("hardDelete") == "true"
{{ end }}

var idInterface []interface{}

for _, id := range ids.IDs {
idInterface = append(idInterface, id)
}

deleted, err := models.{{ $alias.UpPlural }}(qm.WhereIn("id IN ?", idInterface...)).DeleteAllG(c.Request.Context(){{if $soft}}, hardDelete{{end}})
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

c.JSON(http.StatusOK, DeletedCount{DeletedCount: int(deleted)})
}

{{ if $soft }}
// UnDelete{{ $alias.UpSingular }}ByID godoc
// @Summary Undeletes a {{ $alias.DownSingular }} by id
// @Produce json
// @Success 200 {object} API{{ $alias.UpSingular }}
// @Param id path string true "{{ $alias.UpSingular }} id"
// @Router /crud/{{ $alias.DownPlural }}/:id/unDelete [post]
func (*GeneratedCrudController) UnDelete{{ $alias.UpSingular }}ByID(c *gin.Context) {
id := c.Param("id")

if id == "" {
api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
return
}

deleted{{ $alias.UpSingular }}, err := models.{{ $alias.UpPlural }}(qm.Where("id = ?", id), qm.WithDeleted()).OneG(c.Request.Context())
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

deleted{{ $alias.UpSingular }}.DeletedAt = null.Time{
Valid: false,
}

_, err = deleted{{ $alias.UpSingular }}.UpdateG(c.Request.Context(), boil.Whitelist("deleted_at"))
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

c.JSON(http.StatusOK, deleted{{ $alias.UpSingular }})
}
{{ end }}

func (gcc *GeneratedCrudController) Register{{ $alias.UpPlural }}(rg *gin.RouterGroup) {
rg.GET("/{{ $alias.DownPlural }}/:id", gcc.Get{{ $alias.UpSingular }}ByID)
rg.GET("/{{ $alias.DownPlural }}", gcc.Get{{ $alias.UpPlural }})
rg.PUT("/{{ $alias.DownPlural }}/:id", gcc.Update{{ $alias.UpSingular }}ByID)
rg.DELETE("/{{ $alias.DownPlural }}/:id", gcc.Delete{{ $alias.UpSingular }}ByID)
rg.DELETE("/{{ $alias.DownPlural }}", gcc.BulkDelete{{ $alias.UpPlural }}ByIDs)
{{ if $soft -}}
rg.POST("/{{ $alias.DownPlural }}/:id/unDelete", gcc.UnDelete{{ $alias.UpSingular }}ByID)
{{- end -}}
}