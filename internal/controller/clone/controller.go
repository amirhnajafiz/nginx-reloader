package clone

import (
	"fmt"
	"os/exec"
)

// clone controller clones into a git repository to get the files
// from the given address.
type controller struct {
	localDir string
	callback func(string) error
}

func New(ld string, cb func(string) error) *controller {
	return &controller{
		localDir: ld,
		callback: cb,
	}
}

func (c controller) GetFiles(address string) error {
	// clone the Git repository
	cmd := exec.Command("git", "clone", address, c.localDir)

	// execute the git clone command
	msg, err := cmd.CombinedOutput()
	if err != nil {
		if len(msg) > 0 {
			err = fmt.Errorf("message: %s, err: %s", msg, err)
		}

		return fmt.Errorf("failed to run git command: %v", err)
	}

	return c.callback(c.localDir)
}
