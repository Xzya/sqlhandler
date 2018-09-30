package sqlhandler

import (
	"github.com/clipperhouse/typewriter"
	"io"
)

func init() {
	err := typewriter.Register(NewSQLHandlerWriter())
	if err != nil {
		panic(err)
	}
}

func HandlerName(typ typewriter.Type) string {
	return typ.Name + "Handler"
}

type SQLHandlerWriter struct{}

func NewSQLHandlerWriter() *SQLHandlerWriter {
	return &SQLHandlerWriter{}
}

func (sw *SQLHandlerWriter) Name() string {
	return "sqlhandler"
}

func (sw *SQLHandlerWriter) Imports(typ typewriter.Type) (result []typewriter.ImportSpec) {
	result = append(result, typewriter.ImportSpec{
		Path: "github.com/jinzhu/gorm",
	})
	return
}

func (sw *SQLHandlerWriter) Write(w io.Writer, typ typewriter.Type) error {
	tag, found := typ.FindTag(sw)

	if !found {
		return nil
	}

	// start with the sqlhandler template
	tmpl, err := templates.ByTag(typ, tag)

	if err != nil {
		return err
	}

	m := model{
		Type:        typ,
		HandlerName: HandlerName(typ),
	}

	if err := tmpl.Execute(w, m); err != nil {
		return err
	}

	for _, v := range tag.Values {
		m := model{
			Type:          typ,
			HandlerName:   HandlerName(typ),
			TypeParameters: v.TypeParameters,
			TagValue:      v,
		}

		tmpl, err := templates.ByTagValue(typ, v)

		if err != nil {
			return err
		}

		if err := tmpl.Execute(w, m); err != nil {
			return err
		}
	}

	return nil
}
