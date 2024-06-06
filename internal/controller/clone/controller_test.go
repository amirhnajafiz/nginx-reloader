package clone

import (
	"testing"
)

func callback() error {
	return nil
}

func TestCloneController(t *testing.T) {
	ctl := New("./tmp/clone", callback)

	if err := ctl.GetFiles("https://github.com/amirhnajafiz/playbooks.git"); err != nil {
		t.Error(err)
	}
}
