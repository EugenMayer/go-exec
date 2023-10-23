package utils

import (
	"bufio"
	"fmt"
	"github.com/eugenmayer/go-exec/runner"
	"strings"
)

// return all files in a folder
func FilesInFolder(runner runner.CommandRunner, srcPath string, filenamesOnly bool) ([]string, error) {
	var cmd = fmt.Sprintf("find %s -type f", srcPath)
	if filenamesOnly {
		cmd = cmd + " -printf '%f\n'"
	}
	stdout,_, err := runner.Run(cmd)
	if err != nil {
		return []string{}, err
	}

	scanner := bufio.NewScanner(strings.NewReader(stdout))
	var filePaths []string
	for scanner.Scan() {
		filePaths = append(filePaths, scanner.Text())
	}

	return filePaths, nil
}
