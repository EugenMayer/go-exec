package runner

import (
	"fmt"
	"github.com/eugenmayer/go-sshclient/sshwrapper"
)

type SshCommandRunner struct {
	SshApi  *sshwrapper.SshApi
	Verbose bool
}

func (runner SshCommandRunner) Run(cmdStr string) (stdOut string, stdErr string, err error) {
	if runner.Verbose {
		fmt.Println(fmt.Sprintf("Using SSH. Running on server: %s. Running command: %s", runner.SshApi.Host, cmdStr))
	}

	stdOut, stdErr, err = runner.SshApi.Run(cmdStr)
	return
}
