package sqlhandler

import "github.com/clipperhouse/typewriter"

var findAssociations = &typewriter.Template{
	Name: "FindAssociations",
	Text: `
// Find{{.Type}}Associations finds the {{.Type}}'s associations
func Find{{.Type}}Associations(db *gorm.DB, item *{{.Type}}) error {
{{ $params := . }}
{{ range $index, $TypeParameter := .TypeParameters }}
	if err := Find{{$params.Type}}{{$TypeParameter}}Association(db, item); err != nil {
		return err
	}
{{ end }}
	return nil
}
`,
}
