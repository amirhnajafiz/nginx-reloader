package fetch

import (
	"os/exec"
)

// fetch controller downloads the files from the given address.
// it uses wget command.
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
	// use wget to download the content from the given address
	// the -P flag specifies the directory prefix where all retrieved files and directories will be saved to
	cmd := exec.Command("wget", "-P", c.localDir, "-r", address)

	// execute the wget command
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return c.callback()
}
