# WAT

Simplify command execution on the shell even more to get you a one-liner with stdout/stderr capturing (`/exec`)
Advanced topic, introduce a command/strategy patter to abstract if want to run a command locally/remote using ssh/on a docker container (`/runnte/`)

## Usage

See the examples at [go-shell-cli-quickstarter](https://github.com/EugenMayer/go-shell-cli-quickstarter/blob/master/cmd/myexec.go#L21)

```go

package mystuff

import (
	"github.com/eugenmayer/go-exec/exec/v1"
)


if stdout, stderr, err := exec.Run("echo hi"); err != nil {
    log.Print(stdout)
    log.Print(stderr)
    log.Fatal(err)
}

if stdout, stderr, err := exec.Run("echo ho"); err != nil {
    log.Print(stdout)
    log.Print(stderr)
    log.Fatal(err)
}

// continue..
``` 
