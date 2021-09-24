{{- $alias := .Aliases.Table .Table.Name -}}
{{- $orig_tbl_name := .Table.Name -}}
{{- $canSoftDelete := .Table.CanSoftDelete -}}
{{- $soft := and .AddSoftDeletes $canSoftDelete }}

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

pks := strings.Split(id, ";")

if len(pks) != {{ len .Table.PKey.Columns }} {
api.NewAPIError("invalid-id", http.StatusBadRequest, fmt.Sprintf("Expected {{ len .Table.PKey.Columns }} ids, got %d", len(pks))).Respond(c)
return
}

qms := []qm.QueryMod{}

{{ range $i, $pk := .Table.PKey.Columns }}
    qms = append(qms, qm.Where("{{ $pk }} = ?", pks[{{ $i }}]))
{{ end }}

existing{{ $alias.UpSingular }}, err := models.{{ $alias.UpPlural }}(qms...).OneG(c.Request.Context())
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
