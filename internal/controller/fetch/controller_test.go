package fetch

import (
	"testing"
)

func callback() error {
	return nil
}

func TestFetchController(t *testing.T) {
	ctl := New("./tmp/fetch", callback)

	if err := ctl.GetFiles("https://github.com/amirhnajafiz/playbooks/archive/refs/heads/main.zip"); err != nil {
		t.Error(err)
	}
}
