{{- $alias := .Aliases.Table .Table.Name -}}
{{- $orig_tbl_name := .Table.Name -}}
{{- $canSoftDelete := .Table.CanSoftDelete -}}
{{- $soft := and .AddSoftDeletes $canSoftDelete }}

func (gcc *GeneratedCrudController) Register{{ $alias.UpPlural }}(rg *gin.RouterGroup) {
rg.GET("/{{ $alias.DownPlural }}/:id", gcc.Get{{ $alias.UpSingular }}ByID)
rg.GET("/{{ $alias.DownPlural }}", gcc.Get{{ $alias.UpPlural }})
rg.GET("/{{ $alias.DownPlural }}/minisearch", gcc.Minisearch{{ $alias.UpPlural }})
rg.PUT("/{{ $alias.DownPlural }}/:id", gcc.Update{{ $alias.UpSingular }}ByID)
rg.POST("/{{ $alias.DownPlural }}", gcc.Create{{ $alias.UpSingular }})
rg.DELETE("/{{ $alias.DownPlural }}/:id", gcc.Delete{{ $alias.UpSingular }}ByID)
rg.DELETE("/{{ $alias.DownPlural }}", gcc.BulkDelete{{ $alias.UpPlural }}ByIDs)
{{ if $soft -}}
rg.POST("/{{ $alias.DownPlural }}/:id/unDelete", gcc.UnDelete{{ $alias.UpSingular }}ByID)
{{- end -}}
}