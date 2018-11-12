package utils

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

func TarGz(runner runner.CommandRunner, baseFolder string, files []string, destFile string) (error) {
	_,_, err := runner.Run(fmt.Sprintf("cd %s && tar cfz %s -%s", baseFolder, destFile, strings.Join(files, " ")))
	if err != nil {
		return err
	}
	return nil
}