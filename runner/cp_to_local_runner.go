package runner

import (
	"github.com/eugenmayer/go-sshclient/sshwrapper"
	"io"
	"os"
)

type CopyToLocalFromLocalRunner struct {
	SshApi  *sshwrapper.SshApi
	Verbose bool
}

func NewCopyToLocalFromLocalRunner(sshApi  *sshwrapper.SshApi) CopyToLocalFromLocalRunner {
	return CopyToLocalFromLocalRunner{
		SshApi: sshApi,
	}
}

func (runner CopyToLocalFromLocalRunner) Copy(source string, dest string) (stdOut string, stdErr string, err error) {
	in, err := os.Open(source)
	if err != nil {
		return "", "", err
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return "", "", err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return "", "", err
	}
	return "", "", out.Close()
}
