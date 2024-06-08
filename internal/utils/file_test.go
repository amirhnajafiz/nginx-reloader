package utils

import "testing"

func generateFiles(size int, path string) {

}

func checkExistance(path string) int {
	return 0
}

func removeFiles(path string) {

}

func TestMoveFiles(t *testing.T) {
	src := ""
	dest := ""
	size := 0

	generateFiles(size, src)

	MoveDir(src, dest)

	if count := checkExistance(dest); count != size {
		t.Errorf("moved %d out of %d\n", count, size)
	}

	removeFiles(dest)
}
