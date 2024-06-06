package clone

import (
	"os"
	"testing"
)

func callback() error {
	os.RemoveAll("./tmp/clone")

	return nil
}

func TestCloneController(t *testing.T) {
	ctl := New("./tmp/clone", callback)

	if err := ctl.GetFiles("https://github.com/amirhnajafiz/playbooks.git"); err != nil {
		t.Error(err)
	}
}
