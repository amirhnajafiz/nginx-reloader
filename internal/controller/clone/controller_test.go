package clone

import (
	"fmt"
	"testing"
)

func callback() error {
	fmt.Println("done.")

	return nil
}

func TestCloneController(t *testing.T) {
	ctl := New("./tmp", callback)

	if err := ctl.GetFiles(""); err != nil {
		t.Error(err)
	}
}
