// +build !libcontainer,!runc

package supervisor

import (
	"errors"
    "fmt"
	"github.com/docker/containerd/runtime"
)

func newRuntime(stateDir string) (runtime.Runtime, error) {
	fmt.Println("###### newRuntime prepared #######\r\n")
	return nil, errors.New("unsupported platform")
}
