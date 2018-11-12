package exec

import (
	"bufio"
	"fmt"
	"github.com/eugenmayer/go-exec/runner"
	"strings"
)

func FilesInFolder(runner runner.CommandRunner, srcPath string) ([]string, error) {
	stdout,_, err := runner.Run(fmt.Sprintf("find %s -type f", srcPath))
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