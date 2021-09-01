type TableField struct {
Name     string `json:"name"`
Type     string `json:"type"`
HasDefault  bool `json:"has_default"`
Nullable bool   `json:"nullable"`
}