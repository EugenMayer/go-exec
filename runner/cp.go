package runner

type CopyRunner interface {
	Copy(source string, dest string) (stdOut string, stdErr string, err error)
}
