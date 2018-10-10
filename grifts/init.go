package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/obiknows/obinnas_abstract/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
