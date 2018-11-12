package runner

import (
	"fmt"
	"github.com/eugenmayer/go-exec/exec"
)

type DockerCommandRunner struct {
	Container string
	Verbose bool
}

func (runner DockerCommandRunner) Run(cmdStr string) (stdOut string, stdErr string, err error) {
	if runner.Verbose {
		fmt.Println(fmt.Sprintf("Using Docker. Running on container: %s. Running command: %s", runner.Container, cmdStr))
	}
	return exec.Run(fmt.Sprintf("docker exec %s %s", runner.Container, cmdStr), runner.Verbose)
}
