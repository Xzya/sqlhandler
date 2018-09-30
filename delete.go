package sqlhandler

import "github.com/clipperhouse/typewriter"

var delete = &typewriter.Template{
	Name: "Delete",
	Text: `
// Delete deletes a {{.Type}} that matches the given conditions
func (h *{{.HandlerName}}) Delete(value interface{}, where ...interface{}) error {
	return h.db.Delete(value, where...).Error
}
`}
