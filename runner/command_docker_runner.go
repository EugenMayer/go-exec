package runner

import (
	"fmt"
	"github.com/eugenmayer/go-exec/exec"
)

type DockerCommandRunner struct {
	Container string
	Shell string // used to run your command in a shell docker exec %CONTAINER %SHELL -c '%CMD'
	CommandPattern string // use this to run your own run patter, should look like docker exec %s %s (first is container name, second your command, shell is no used here)
	Verbose bool
}

func NewDockerCommandRunner(container string) DockerCommandRunner {
	return DockerCommandRunner{
		Container: container,
		Shell: "bash",
		CommandPattern: "",
	}
}

func (runner DockerCommandRunner) Run(cmdStr string) (stdOut string, stdErr string, err error) {
	if runner.Verbose {
		fmt.Println(fmt.Sprintf("Using Docker. Running on container: %s. Running command: %s", runner.Container, cmdStr))
	}
	// TODO: warn the user of the usage of single ticks or escape them here
	var cmd string

	if runner.CommandPattern != "" {
		cmd = fmt.Sprintf(runner.CommandPattern, runner.Container, cmdStr)
	} else {
		cmd = fmt.Sprintf("docker exec %s %s -c '%s'", runner.Container, runner.Shell, cmdStr)
	}

	return exec.Run(cmd, runner.Verbose)
}
