package runner

import (
	"fmt"
	"github.com/eugenmayer/go-exec/exec"
)

type DockerCommandRunner struct {
	Container string
	Shell string
	Verbose bool
}

func (runner DockerCommandRunner) Run(cmdStr string) (stdOut string, stdErr string, err error) {
	if runner.Verbose {
		fmt.Println(fmt.Sprintf("Using Docker. Running on container: %s. Running command: %s", runner.Container, cmdStr))
	}
	// TODO: warn the user of the usage of single ticks or escape them here
	return exec.Run(fmt.Sprintf("docker exec %s %s -c '%s'", runner.Container, runner.Shell, cmdStr), runner.Verbose)
}
