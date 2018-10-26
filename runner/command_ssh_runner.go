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
	err = runner.SshApi.ConnectAndSession()
	if err != nil {
		if runner.SshApi.Session != nil {
			runner.SshApi.Session.Close()
		}
		return
	}

	if runner.Verbose {
		fmt.Println(fmt.Sprintf("Using SSH. Running on server: %s. Running command: %s", runner.SshApi.Host, cmdStr))
	}
	err = runner.SshApi.Session.Run(cmdStr)
	if err != nil {
		if runner.SshApi.Session != nil {
			runner.SshApi.Session.Close()
		}
		return runner.SshApi.GetStdOut(), runner.SshApi.GetStdOut(), err
	}
	// else
	runner.SshApi.Session.Close()
	return runner.SshApi.GetStdOut(), runner.SshApi.GetStdOut(), nil
}
