package fetch

import (
	"fmt"
	"testing"
)

func callback() error {
	fmt.Println("done.")

	return nil
}

func TestFetchController(t *testing.T) {
	ctl := New("./tmp", callback)

	if err := ctl.GetFiles(""); err != nil {
		t.Error(err)
	}
}
