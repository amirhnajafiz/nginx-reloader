package utils

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func generatePath(path ...string) {
	for _, item := range path {
		os.Mkdir(item, 0700)
	}
}

func generateFiles(size int, path string) {
	dummy := "dummy message"

	for i := 0; i < size; i++ {
		address := fmt.Sprintf("%s/file.%d", path, i)

		if err := os.WriteFile(address, []byte(dummy), 0700); err != nil {
			log.Println(err.Error())
		}
	}
}

func checkExistance(path string) int {
	files, _ := os.ReadDir(path)
	return len(files)
}

func removeFiles(path string) {
	os.RemoveAll(path)
}

func TestMoveFiles(t *testing.T) {
	src := "./tmp-1"
	dest := "./tmp-2"
	size := 10

	generatePath("./tmp-1", "./tmp-2")
	generateFiles(size, src)

	MoveDir(src, dest)

	if count := checkExistance(dest); count != size {
		t.Errorf("moved %d out of %d\n", count, size)
	}

	removeFiles(dest)
	removeFiles(src)
}
