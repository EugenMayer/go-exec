package runner

import (
	"github.com/eugenmayer/go-sshclient/sshwrapper"
)



type CopyToRemoteFromLocalRunner struct {
	SshApi  *sshwrapper.SshApi
	Verbose bool
}

func NewCopyToRemoteFromLocalRunner(sshApi  *sshwrapper.SshApi) CopyToRemoteFromLocalRunner {
	return CopyToRemoteFromLocalRunner{
		SshApi: sshApi,
	}
}

func (runner CopyToRemoteFromLocalRunner) Copy(source string, dest string) (stdOut string, stdErr string, err error) {
	err = runner.SshApi.CopyToRemote(source, dest)
	return "", "", err
}
