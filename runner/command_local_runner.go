package runner

import (
	"github.com/eugenmayer/go-exec/exec"
)

type LocalCommandRunner struct {
	Verbose bool
}

func (runner LocalCommandRunner) Run(cmdStr string) (stdOut string, stdErr string, err error) {
	return exec.Run(cmdStr, runner.Verbose)
}
