package utils

import (
	"bufio"
	"fmt"
	"github.com/eugenmayer/go-exec/runner"
	"log"
	"os"
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

// Tar using the system cli for transport-agnostic packaging (works for ssh/docker shell calls)
func TarGz(runner runner.CommandRunner, baseFolder string, files []string, destFile string) (error) {
	_,_, err := runner.Run(fmt.Sprintf("tar -C %s -czf %s %s", baseFolder, destFile, strings.Join(files, " ")))
	if err != nil {
		return err
	}
	return nil
}

// ask the user for confirmation
func ConfirmQuestion(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}