
type GeneratedCrudController struct {}

func RegisterGeneratedCrud(rg *gin.RouterGroup) {
    gcc := GeneratedCrudController{}

    {{ range $table := .Tables }}
        gcc.Register{{ titleCase $table.Name }}(rg)
    {{- end }}

    rg.GET("/_schema", func(c *gin.Context) {
        c.JSON(200, AdminInfo)
    })
}