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

pks := strings.Split(id, ";")

if len(pks) != {{ len .Table.PKey.Columns }} {
api.NewAPIError("invalid-id", http.StatusBadRequest, fmt.Sprintf("Expected {{ len .Table.PKey.Columns }} ids, got %d", len(pks))).Respond(c)
return
}

qms := []qm.QueryMod{}

{{ range $i, $pk := .Table.PKey.Columns }}
    qms = append(qms, qm.Where("{{ $pk }} = ?", pks[{{ $i }}]))
{{ end }}

{{ $alias.UpSingular }}, err := models.{{ $alias.UpPlural }}(qms...).OneG(c.Request.Context())
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
// @Success 200 {object} APIGet{{ $alias.UpPlural }}Response
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
        case "{{ $orig_col_name }}.null":
        if v[0] == "true" {
        queryMods = append(queryMods, qm.Where("{{ $orig_col_name }} IS NULL"))
        } else {
        queryMods = append(queryMods, qm.Where("{{ $orig_col_name }} IS NOT NULL"))
        }
        {{ if eq $column.Type "int" }}
            {{ template "numeric_query_operators" $orig_col_name }}
        {{ else if eq $column.Type "time.Time" }}
            {{ template "numeric_query_operators" $orig_col_name }}
        {{ else if eq $column.Type "null.Int" }}
            {{ template "numeric_query_operators" $orig_col_name }}
        {{ else if eq $column.Type "null.Time" }}
            {{ template "numeric_query_operators" $orig_col_name }}
        {{ else if eq $column.Type "null.Float64" }}
            {{ template "numeric_query_operators" $orig_col_name }}
        {{ else if eq $column.Type "null.Bool" }}
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


type Get{{ $alias.UpPlural }}MinisearchResponse struct {
Items models.{{ $alias.UpSingular }}Slice `json:"items"`
Total int64            `json:"total"`
}

type APIGet{{ $alias.UpPlural }}MinisearchResponse struct {
Items []API{{ $alias.UpSingular }} `json:"items"`
Total int64            `json:"total"`
}

// Minisearch{{ $alias.UpPlural }} godoc
// @Summary Quick search on predefined fields for {{ $alias.UpPlural }}
// @Produce json
// @Success 200 {object} APIGet{{ $alias.UpPlural }}MinisearchResponse
// @Param q query string true "Query fields eg: ?q=Bobbo"
// @Router /crud/{{ $alias.DownPlural }}/minisearch [get]
func (*GeneratedCrudController) Minisearch{{ $alias.UpPlural }}(c *gin.Context) {
queryMods := []qm.QueryMod{}

ss := fmt.Sprintf("%%%s%%", strings.ReplaceAll(c.Query("q"), "%", "\\%"))

for _, field := range {{ $alias.UpPlural }}Admin.Fields {
if field.Config.MiniSearchable {
queryMods = append(queryMods, qm.Or(fmt.Sprintf("%s ILIKE ?", field.ID), ss))
}
}

count, err := models.{{ $alias.UpPlural }}(queryMods...).CountG(c.Request.Context())
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

limit, offset := api.ExtractLimitOffset(c)
queryMods = append(queryMods, qm.Limit(limit), qm.Offset(offset))

results, err := models.{{ $alias.UpPlural }}(queryMods...).AllG(c)
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

c.JSON(http.StatusOK, gin.H{"total": count, "items": results})
}
