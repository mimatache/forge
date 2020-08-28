package command

import (
	"fmt"
	"github.com/mimatache/forge/internal/manifest"
)

func NewCommandList(forgeryName manifest.ForgeryName, forgeries map[manifest.ForgeryName]manifest.Forgery) (manifest.Forgeries, error) {
	commands := executables{}
	err := commands.populateCommandList(forgeryName, forgeries)
	return commands.commands, err
}

type executables struct {
	commands manifest.Forgeries
}

func (e *executables) populateCommandList(forgeryName manifest.ForgeryName, forgeries map[manifest.ForgeryName]manifest.Forgery) error {
	forgery, ok := forgeries[forgeryName]
	if !ok {
		return fmt.Errorf("undefined forgery requested: %s", forgeryName)
	}
	for _, v := range forgery.Pre {
		err := e.populateCommandList(v, forgeries)
		if err != nil {
			return err
		}
	}

	e.commands = append(e.commands, forgery)
	return nil
}
