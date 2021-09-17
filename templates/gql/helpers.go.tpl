{{- $alias := .Aliases.Table .Table.Name -}}
{{- $orig_tbl_name := .Table.Name -}}
{{- $canSoftDelete := .Table.CanSoftDelete -}}
{{- $soft := and .AddSoftDeletes $canSoftDelete -}}

type CanUpdate{{ $alias.UpSingular }}Func func({{ $alias.DownSingular }} *models.{{ $alias.UpSingular }}, currentUser *models.User) bool

func Update{{ $alias.UpSingular }}Helper(ctx context.Context, id string, ud *model.Update{{ $alias.UpSingular }}, canUpdate CanUpdate{{ $alias.UpSingular }}Func) (*models.{{ $alias.UpSingular }}, error) {
existing, err := models.{{ $alias.UpPlural }}(qm.Where("id = ?", id)).OneG(ctx)
if err != nil {
return nil, err
}

if !canUpdate(existing, util.ExtractUser(ctx)) {
return nil, api.NewAPIError("missing-permissions", http.StatusForbidden, "Missing required permissions")
}

{{ range $column := .Table.Columns -}}
    {{- $colAlias := $alias.Column $column.Name -}}
    {{- $orig_col_name := $column.Name -}}
    {{- if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
    {{- else }}
        {{- if eq $orig_col_name "id" }}
        {{- else if eq $orig_col_name "created_at" -}}
        {{- else if eq $orig_col_name "updated_at" -}}
        {{- else if eq $orig_col_name "deleted_at" -}}
        {{- else if eq $orig_col_name "created_by_id" -}}
        {{- else if eq $column.Type "types.JSON" -}}
        {{- else if eq $column.Type "null.JSON" -}}
        {{- else -}}
            if ud.{{ $colAlias }} != nil {
            {{ if eq $column.Type "null.Int" }}
                existing.{{ $colAlias }} = null.IntFromPtr(ud.{{ $colAlias }})
            {{ else if eq $column.Type "null.Time" }}
                existing.{{ $colAlias }} = null.TimeFromPtr(ud.{{ $colAlias }})
            {{ else if eq $column.Type "null.Float64" }}
                existing.{{ $colAlias }} = null.Float64FromPtr(ud.{{ $colAlias }})
            {{ else if eq $column.Type "types.StringArray" }}
                existing.{{ $colAlias }} = util.StringPtrSliceToStringArray(ud.{{ $colAlias }})
            {{ else }}
                existing.{{ $colAlias }} = *ud.{{ $colAlias }}
            {{ end }}
            }
        {{- end }}
    {{ end -}}
{{- end -}}

_, err = existing.UpdateG(ctx, boil.Infer())
if err != nil {
return nil, err
}

return existing, nil
}

type Pre{{ $alias.UpSingular }}CreateFunc func({{ $alias.DownSingular }} *models.{{ $alias.UpSingular }}, currentUser *models.User) error

func Create{{ $alias.UpSingular }}(ctx context.Context, c *model.Create{{ $alias.UpSingular }}, pre Pre{{ $alias.UpSingular }}CreateFunc) (*models.{{ $alias.UpSingular }}, error) {
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
        {{- else if eq $orig_col_name "created_by_id" -}}
        {{- else if eq $column.Type "types.JSON" -}}
        {{- else if eq $column.Type "null.JSON" -}}
        {{ else if eq $orig_col_name "id" }}
            {{ if and (eq $orig_col_name "id") (eq $column.DBType "uuid") }}
                new{{ $alias.UpSingular }}.{{ $colAlias }} = uuid.Must(uuid.NewV4()).String()
            {{ else }}
                new{{ $alias.UpSingular }}.{{ $colAlias }} = ""
            {{ end }}
        {{else -}}
            {{ if or $column.Nullable $hasDefault }}
                if c.{{ $colAlias }} != nil {
                    {{ if eq $column.Type "null.Int" }}
                        new{{ $alias.UpSingular }}.{{ $colAlias }} = null.IntFromPtr(c.{{ $colAlias }})
                    {{ else if eq $column.Type "null.Time" }}
                        new{{ $alias.UpSingular }}.{{ $colAlias }} = null.TimeFromPtr(c.{{ $colAlias }})
                    {{ else if eq $column.Type "null.Float64" }}
                        new{{ $alias.UpSingular }}.{{ $colAlias }} = null.Float64FromPtr(c.{{ $colAlias }})
                    {{ else if eq $column.Type "types.StringArray" }}
                        new{{ $alias.UpSingular }}.{{ $colAlias }} = c.{{ $colAlias }}
                    {{ else }}
                        new{{ $alias.UpSingular }}.{{ $colAlias }} = *c.{{ $colAlias }}
                    {{ end }}
                }
            {{ else }}
                {{ if eq $column.Type "null.Int" }}
                    new{{ $alias.UpSingular }}.{{ $colAlias }} = null.IntFromPtr(c.{{ $colAlias }})
                {{ else if eq $column.Type "null.Time" }}
                    new{{ $alias.UpSingular }}.{{ $colAlias }} = null.TimeFromPtr(c.{{ $colAlias }})
                {{ else if eq $column.Type "null.Float64" }}
                    new{{ $alias.UpSingular }}.{{ $colAlias }} = null.Float64FromPtr(c.{{ $colAlias }})
                {{ else if eq $column.Type "types.StringArray" }}
                    new{{ $alias.UpSingular }}.{{ $colAlias }} = c.{{ $colAlias }}
                {{ else }}
                    new{{ $alias.UpSingular }}.{{ $colAlias }} = c.{{ $colAlias }}
                {{ end }}
            {{ end }}
        {{end -}}
    {{end -}}

    if err := pre(&new{{ $alias.UpSingular }}, util.ExtractUser(ctx)); err != nil {
        return nil, err
    }
    
    err := new{{ $alias.UpSingular }}.InsertG(ctx, boil.Infer())
    if err != nil {
        return nil, err
    }

    return &new{{ $alias.UpSingular }}, nil
}

func Create{{ $alias.UpSingular }}TX(ctx context.Context, tx *sql.Tx, c *model.Create{{ $alias.UpSingular }}, pre Pre{{ $alias.UpSingular }}CreateFunc) (*models.{{ $alias.UpSingular }}, error) {
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
    {{- else if eq $orig_col_name "created_by_id" -}}
    {{- else if eq $column.Type "types.JSON" -}}
    {{- else if eq $column.Type "null.JSON" -}}
    {{ else if eq $orig_col_name "id" }}
        {{ if and (eq $orig_col_name "id") (eq $column.DBType "uuid") }}
            new{{ $alias.UpSingular }}.{{ $colAlias }} = uuid.Must(uuid.NewV4()).String()
        {{ else }}
            new{{ $alias.UpSingular }}.{{ $colAlias }} = ""
        {{ end }}
    {{else -}}
        {{ if or $column.Nullable $hasDefault }}
            if c.{{ $colAlias }} != nil {
            {{ if eq $column.Type "null.Int" }}
                new{{ $alias.UpSingular }}.{{ $colAlias }} = null.IntFromPtr(c.{{ $colAlias }})
            {{ else if eq $column.Type "null.Time" }}
                new{{ $alias.UpSingular }}.{{ $colAlias }} = null.TimeFromPtr(c.{{ $colAlias }})
            {{ else if eq $column.Type "null.Float64" }}
                new{{ $alias.UpSingular }}.{{ $colAlias }} = null.Float64FromPtr(c.{{ $colAlias }})
            {{ else if eq $column.Type "types.StringArray" }}
                new{{ $alias.UpSingular }}.{{ $colAlias }} = c.{{ $colAlias }}
            {{ else }}
                new{{ $alias.UpSingular }}.{{ $colAlias }} = *c.{{ $colAlias }}
            {{ end }}
            }
        {{ else }}
            {{ if eq $column.Type "null.Int" }}
                new{{ $alias.UpSingular }}.{{ $colAlias }} = null.IntFromPtr(c.{{ $colAlias }})
            {{ else if eq $column.Type "null.Time" }}
                new{{ $alias.UpSingular }}.{{ $colAlias }} = null.TimeFromPtr(c.{{ $colAlias }})
            {{ else if eq $column.Type "null.Float64" }}
                new{{ $alias.UpSingular }}.{{ $colAlias }} = null.Float64FromPtr(c.{{ $colAlias }})
            {{ else if eq $column.Type "types.StringArray" }}
                new{{ $alias.UpSingular }}.{{ $colAlias }} = c.{{ $colAlias }}
            {{ else }}
                new{{ $alias.UpSingular }}.{{ $colAlias }} = c.{{ $colAlias }}
            {{ end }}
        {{ end }}
    {{end -}}
{{end -}}

if err := pre(&new{{ $alias.UpSingular }}, util.ExtractUser(ctx)); err != nil {
return nil, err
}

err := new{{ $alias.UpSingular }}.Insert(ctx, tx, boil.Infer())
if err != nil {
return nil, err
}

return &new{{ $alias.UpSingular }}, nil
}