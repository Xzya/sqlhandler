package sqlhandler

import "github.com/clipperhouse/typewriter"

var findAssociation = &typewriter.Template{
	Name: "FindAssociation",
	Text: `
{{ $params := . }}
{{ range $index, $TypeParameter := .TypeParameters }}
// Find{{$params.Type}}{{$TypeParameter.LongName}}Association finds the {{$params.Type}}'s {{$TypeParameter}} association
func Find{{$params.Type}}{{$TypeParameter.LongName}}Association(db *gorm.DB, item *{{$params.Type}}) error {
	return db.Model(item).Association("{{$TypeParameter}}").Find(&item.{{$TypeParameter}}).Error
}
{{ end }}
`,
}
