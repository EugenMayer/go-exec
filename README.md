# WAT

Simplify command execution on the shell even more to get you a one-liner with stdout/stderr capturing (`/exec`)
Also advanced, introduce a command/strategy patter to abstract if want to run a command locally or remote using ssh (`/runnter/`)

## Usage

Also see the example at [go-antibash-boilerplate](https://github.com/EugenMayer/go-antibash-boilerplate/blob/master/cmd/myexec.go#L21)

```go

package mystuff

import (
	"github.com/eugenmayer/go-exec/exec"
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