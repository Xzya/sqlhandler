package sqlhandler

import "github.com/clipperhouse/typewriter"

var create = &typewriter.Template{
	Name: "Create",
	Text: `
// Create inserts a {{.Type}} into the database
func (h *{{.HandlerName}}) Create(value *{{.Type}}) error {
	return h.db.Create(value).Error
}
`}
