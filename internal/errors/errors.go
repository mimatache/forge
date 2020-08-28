package errors

import (
	"fmt"
	"github.com/mimatache/forge/internal/manifest"
)

type RecursiveCommand struct {
	Command manifest.ForgeryName
}

func (r *RecursiveCommand) Error() string {
	return fmt.Sprintf("%s forgery has a recursive dependency", r.Command)
}
