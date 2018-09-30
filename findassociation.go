package sqlhandler

import "github.com/clipperhouse/typewriter"

var findAssociation = &typewriter.Template{
	Name: "FindAssociation",
	Text: `
{{ $TypeParameter := (index .TypeParameters 0) }}
// Find{{.Type}}{{$TypeParameter.LongName}}Association finds the {{.Type}}'s {{$TypeParameter}} association
func Find{{.Type}}{{$TypeParameter.LongName}}Association(db *gorm.DB, item *{{.Type}}) error {
	return db.Model(item).Association("{{$TypeParameter}}").Find(&item.{{$TypeParameter}}).Error
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, but no constraints on that type
		{},
	},
}
