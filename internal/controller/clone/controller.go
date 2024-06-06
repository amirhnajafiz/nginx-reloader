package clone

import (
	"os/exec"
)

// clone controller clones into a git repository to get the files
// from the given address.
type controller struct {
	localDir string
	callback func() error
}

func New(ld string, cb func() error) *controller {
	return &controller{
		localDir: ld,
		callback: cb,
	}
}

func (c controller) GetFiles(address string) error {
	// clone the Git repository
	cmd := exec.Command("git", "clone", address, c.localDir)

	// execute the git clone command
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return c.callback()
}
