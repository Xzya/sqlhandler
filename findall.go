package sqlhandler

import "github.com/clipperhouse/typewriter"

var findAll = &typewriter.Template{
	Name: "FindAll",
	Text: `
// FindAll finds all {{.Type}}
func (h {{.HandlerName}}) FindAll() ([]{{.Type}}, error) {
	var results []{{.Type}}

	if err := h.db.Model(&{{.Type}}{}).Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
`}
