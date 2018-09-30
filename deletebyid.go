package sqlhandler

import "github.com/clipperhouse/typewriter"

var deleteByID = &typewriter.Template{
	Name: "DeleteByID",
	Text: `
// Delete deletes a {{.Type}} with the given ID
func (h *{{.HandlerName}}) DeleteByID(ID uint) error {
	return h.db.Delete(&{{.Type}}{}, "id = ?", ID).Limit(1).Error
}
`}
