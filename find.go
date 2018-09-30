package sqlhandler

import "github.com/clipperhouse/typewriter"

var find = &typewriter.Template{
	Name: "Find",
	Text: `
// Find finds a {{.Type}} that matches the given conditions
func (h *{{.HandlerName}}) Find(where ...interface{}) (*{{.Type}}, error) {
	var res {{.Type}}

	if err := h.db.Find(&res, where...).Limit(1).Error; err != nil {
		return nil, err
	}

	return &res, nil
}
`}
