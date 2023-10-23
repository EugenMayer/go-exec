package utils

import (
	"fmt"
	"github.com/eugenmayer/go-exec/runner"
	"strings"
)

// Tar using the system cli for transport-agnostic packaging (works for ssh/docker shell calls)
func TarGz(runner runner.CommandRunner, baseFolder string, files []string, destFile string) (error) {
	_,_, err := runner.Run(fmt.Sprintf("tar -C %s -czf %s %s", baseFolder, destFile, strings.Join(files, " ")))
	if err != nil {
		return err
	}
	return nil
}
