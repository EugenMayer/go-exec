package runner

import (
	"github.com/eugenmayer/go-sshclient/sshwrapper"
)

type CopyToRemoteFromLocalRunner struct {
	SshApi  *sshwrapper.SshApi
	Verbose bool
}

func (runner CopyToRemoteFromLocalRunner) Copy(source string, dest string) (stdOut string, stdErr string, err error) {
	runner.SshApi.ConnectAndSession()
	err = runner.SshApi.CopyToRemote(source, dest)
	runner.SshApi.Session.Close()
	return "", "", err
}
