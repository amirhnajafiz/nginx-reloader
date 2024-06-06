package fetch

import (
	"os"
	"testing"
)

func callback() error {
	os.RemoveAll("./tmp/fetch")

	return nil
}

func TestFetchController(t *testing.T) {
	ctl := New("./tmp/fetch", callback)

	if err := ctl.GetFiles("https://github.com/amirhnajafiz/playbooks/archive/refs/heads/main.zip"); err != nil {
		t.Error(err)
	}
}
