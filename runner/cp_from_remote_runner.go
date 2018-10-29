package runner

import (
	"github.com/eugenmayer/go-sshclient/sshwrapper"
)

type CopyFromRemoteToLocalRunner struct {
	SshApi  *sshwrapper.SshApi
	Verbose bool
}

func (runner CopyFromRemoteToLocalRunner) Copy(source string, dest string) (stdOut string, stdErr string, err error) {
	err = runner.SshApi.CopyFromRemote(source, dest)
	return runner.SshApi.GetStdOut(), runner.SshApi.GetStdErr(), err
}
