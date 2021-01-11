package exec

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Run a command right away
func Run(cmdStr string, verbose bool) (stdOut string, stdErr string, err error) {
	return RunInDir(cmdStr, "", verbose)
}

// RunDir a command in a specific working directory
func RunInDir(cmdStr string, workingDir string, verbose bool) (stdOut string, stdErr string, err error) {
	if verbose {
		fmt.Println(fmt.Sprintf("Running: %s", cmdStr))
	}

	// Using this variant as we have pipes or stdout redirect at some point
	cmd := exec.Command("sh", "-c", cmdStr)
	// since we use dash, we could have a very basic env without proper paths. Ensure we at least have the usual extension point
	// for extra binaries
	cmd.Env = append(os.Environ(), "PATH="+os.Getenv("PATH")+":/usr/local/bin")
	cmd.Dir = workingDir
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()

	if err != nil {
		err = errors.New(err.Error() + "\n" + stderr.String())
	}

	return stdout.String(), stderr.String(), err
}

// Run a command and fatal if an error occurs ( non exit 0 )
func RunFatal(cmdStr string, verbose bool) (stdOut string, stdErr string) {
	stdOut,stdErr, err := Run(cmdStr, verbose)
	if err != nil {
		log.Fatal(err)
	}

	return
}
