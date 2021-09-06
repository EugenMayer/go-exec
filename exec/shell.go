package exec

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Start a command. You will need to call .run() yourself. Use this for non-blocking execution like live error streaming.
func Start(cmdStr string, workDir string, verbose bool) (*exec.Cmd, error) {
	if verbose {
		fmt.Println(fmt.Sprintf("Running: %s", cmdStr))
	}

	// Using this variant as we have pipes or stdout redirect at some point
	cmd := exec.Command("sh", "-c", cmdStr)
	err := cmd.Start()

	if err != nil {
		return nil, nil
	}

	// since we use dash, we could have a very basic env without proper paths. Ensure we at least have the usual extension point
	// for extra binaries
	cmd.Env = append(os.Environ(), "PATH="+os.Getenv("PATH")+":/usr/local/bin")
	cmd.Dir = workDir
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	return cmd, err
}

// Run a command right away
func Run(cmdStr string, verbose bool) (stdOut string, stdErr string, err error) {
	return RunInDir(cmdStr, "", verbose)
}

// RunInDir RunDir a command in a specific working directory
func RunInDir(cmdStr string, workingDir string, verbose bool) (stdOut string, stdErr string, err error) {
	var cmd, startErr = Start(cmdStr, workingDir, verbose)
	if startErr != nil {
		return "", "", startErr
	}

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

// RunFatal Run a command and fatal if an error occurs ( non exit 0 )
func RunFatal(cmdStr string, verbose bool) (stdOut string, stdErr string) {
	stdOut,stdErr, err := Run(cmdStr, verbose)
	if err != nil {
		log.Fatal(err)
	}

	return
}
