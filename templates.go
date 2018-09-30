package sqlhandler

import "github.com/clipperhouse/typewriter"

// a convenience for passing values into templates; in MVC it'd be called a view model
type model struct {
	Type           typewriter.Type
	HandlerName    string
	TypeParameters []typewriter.Type
	typewriter.TagValue
}

var templates = typewriter.TemplateSlice{
	handler,

	create,

	find,
	findAll,
	findByID,
	findAssociation,

	delete,
	deleteByID,
}
