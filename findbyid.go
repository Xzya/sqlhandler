package sqlhandler

import "github.com/clipperhouse/typewriter"

var findByID = &typewriter.Template{
	Name: "FindByID",
	Text: `
// FindByID finds a {{.Type}} with the given ID
func (h *{{.HandlerName}}) FindByID(ID uint) (*{{.Type}}, error) {
	var res {{.Type}}

	if err := h.db.Find(&res, "id = ?", ID).Limit(1).Error; err != nil {
		return nil, err
	}

	return &res, nil
}
`}
