package fetch

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// fetch controller downloads the files from the given address.
// it uses wget command.
type controller struct {
	localDir string
	filename string
	callback func(string) error
}

func New(ld, fn string, cb func(string) error) *controller {
	return &controller{
		localDir: ld,
		filename: fn,
		callback: cb,
	}
}

func (c controller) GetFiles(address string) error {
	// get the file name from URL and determine the command to use for extraction
	filename := filepath.Base(address)
	if len(c.filename) > 0 {
		filename = c.filename
	}

	// download the file using wget
	if err := exec.Command("wget", address, "-P", c.localDir).Run(); err != nil {
		return fmt.Errorf("error downloading file: %v", err)
	}

	// determine how to extract based on file extension
	if strings.HasSuffix(filename, ".zip") {
		// unzip the file
		unzipCmd := exec.Command("unzip", filepath.Join(c.localDir, filename), "-d", c.localDir)
		if err := unzipCmd.Run(); err != nil {
			return fmt.Errorf("error unzipping file: %v", err)
		}
	} else if strings.HasSuffix(filename, ".tar.gz") {
		// extract tar.gz file
		tarCmd := exec.Command("tar", "xzf", filepath.Join(c.localDir, filename), "-C", c.localDir)
		if err := tarCmd.Run(); err != nil {
			return fmt.Errorf("error extracting tar.gz file: %v", err)
		}
	}

	return c.callback(filepath.Join(c.localDir, filename))
}
