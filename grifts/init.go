package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/mrNobody95/adminyar/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
