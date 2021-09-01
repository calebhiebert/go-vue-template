{{- $alias := .Aliases.Table .Table.Name -}}

// Get{{ $alias.UpSingular }}ByID godoc
// @Summary Gets a single {{ $alias.UpSingular }} entity by their id
// @Produce json
// @Success 200 {object} {{ $alias.UpSingular }}
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
// @Success 200 {object} {{ $alias.UpSingular }}Slice
// @Router /crud/{{ $alias.DownPlural }} [get]
func (*GeneratedCrudController) Get{{ $alias.UpPlural }}(c *gin.Context) {
limit, offset := api.ExtractLimitOffset(c)

{{ $alias.DownPlural }}, err := models.{{ $alias.UpPlural }}(qm.Limit(limit), qm.Offset(offset)).AllG(c.Request.Context())
if err != nil {
api.APIErrorFromErr(err).Respond(c)
return
}

c.JSON(http.StatusOK, {{ $alias.DownPlural }})
}

func (gcc *GeneratedCrudController) Register{{ $alias.UpPlural }}(rg *gin.RouterGroup) {
rg.GET("/{{ $alias.DownPlural }}/:id", gcc.Get{{ $alias.UpSingular }}ByID)
rg.GET("/{{ $alias.DownPlural }}", gcc.Get{{ $alias.UpPlural }})
}