package sqlhandler

import "github.com/clipperhouse/typewriter"

var handler = &typewriter.Template{
	Name: "sqlhandler",
	Text: `
// {{.HandlerName}} handles all {{.Type}} database operations
type {{.HandlerName}} struct {
	db *gorm.DB
}

func New{{.HandlerName}}(db *gorm.DB) *{{.HandlerName}}{
	return &{{.HandlerName}}{
		db,
	}
}
`}
