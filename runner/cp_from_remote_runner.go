package runner

import (
	"github.com/eugenmayer/go-sshclient/sshwrapper"
	"os"
)

type CopyFromRemoteToLocalRunner struct {
	SshApi  *sshwrapper.SshApi
	Verbose bool
}

func (runner CopyFromRemoteToLocalRunner) Copy(source string, dest string) (stdOut string, stdErr string, err error) {
	runner.SshApi.Connect()

	if session, err := runner.SshApi.Client.NewSession(); err != nil {
		return "", "", err
	} else {
		runner.SshApi.Session = session
	}

	r, err := runner.SshApi.Session.Output("dd if=" + source)
	if err != nil {
		return "", "", err
	}
	defer runner.SshApi.Session.Close()
	file, err := os.Create(dest)
	if err != nil {
		return "", "", err
	}
	//write to local file
	file.Write(r)
	return "", "", nil
}
