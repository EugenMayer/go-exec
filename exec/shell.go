package exec

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// Start a command. You will need to call .run() yourself. Use this for non-blocking execution like live error streaming.
func Start(cmdStr string, workDir string, verbose bool) (cmd *exec.Cmd, stdOut io.ReadCloser, stdErr io.ReadCloser, err error ) {
	if verbose {
		fmt.Println(fmt.Sprintf("Running: %s", cmdStr))
	}

	// Using this variant as we have pipes or stdout redirect at some point
	cmd = exec.Command("sh", "-c", cmdStr)

	if err != nil {
		return nil, nil,nil, err
	}

	// since we use dash, we could have a very basic env without proper paths. Ensure we at least have the usual extension point
	// for extra binaries
	cmd.Env = append(os.Environ(), "PATH="+os.Getenv("PATH")+":/usr/local/bin")
	cmd.Dir = workDir
	stdOut, err = cmd.StdoutPipe()
	if err != nil {
		return nil, nil,nil, err
	}

	stdErr, err = cmd.StderrPipe()
	if err != nil {
		return nil, nil,nil, err
	}

	err = cmd.Start()

	return cmd, stdOut, stdErr, nil
}

// Run a command right away
func Run(cmdStr string, verbose bool) (stdOut string, stdErr string, err error) {
	return RunInDir(cmdStr, "", verbose)
}

// RunInDir RunDir a command in a specific working directory
func RunInDir(cmdStr string, workingDir string, verbose bool) (stdOut string, stdErr string, err error) {
	var cmd, stdOutStream, stdErrStream, startErr = Start(cmdStr, workingDir, verbose)
	if startErr != nil {
		return "", "", startErr
	}

	stdOut = readBuffer(stdOutStream)
	stdErr = readBuffer(stdErrStream)

	err = cmd.Wait()

	if err != nil {
		err = errors.New(err.Error() + "\n" + stdErr)
	}

	return stdOut, stdErr, err
}

// RunFatal Run a command and fatal if an error occurs ( non exit 0 )
func RunFatal(cmdStr string, verbose bool) (stdOut string, stdErr string) {
	stdOut,stdErr, err := Run(cmdStr, verbose)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func readBuffer(buffer io.ReadCloser) string {
	scanner := bufio.NewScanner(buffer)
	scanner.Split(bufio.ScanWords)
	var out = ""
	for scanner.Scan() {
		m := scanner.Text()
		out = out + fmt.Sprintln(m)
	}
	return out
}
