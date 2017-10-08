package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/requaos/ceremics/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
