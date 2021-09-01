
type GeneratedCrudController struct {}

func RegisterGeneratedCrud(rg *gin.RouterGroup) {
    gcc := GeneratedCrudController{}

    {{ range $table := .Tables }}
        gcc.Register{{ titleCase $table.Name }}(rg)
    {{- end }}
}