package shell

import (
	"context"
	"io"
	"os/exec"
)

type Executor struct {
	Ctx context.Context
	Out io.Writer
	Err io.Writer
}

func (e *Executor) Execute(command string) error {
	cmd := exec.Command(SHELL, EXEC, command)
	cmd.Stdout = e.Out
	cmd.Stderr = e.Err
	return cmd.Run()
}
