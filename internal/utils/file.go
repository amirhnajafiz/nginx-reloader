package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// MoveDir moves all files from src to dest.
func MoveDir(srcDir, destDir string) error {
	// ensure the destination directory exists
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating destination directory: %v", err)
	}

	// move files from srcDir to destDir
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() { // check if it's a file and not a directory
			destPath := filepath.Join(destDir, filepath.Base(path))
			err := MoveFile(path, destPath)
			if err != nil {
				return fmt.Errorf("failed to move %s to %s: %v", path, destPath, err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking through files: %v", err)
	}

	return nil
}

// MoveFile moves a file from src to dest and removes the src file.
func MoveFile(src, dest string) error {
	// open the source file
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	// create the destination file
	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	// copy the contents
	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	// close the source file
	err = source.Close()
	if err != nil {
		return err
	}

	// remove the source file
	err = os.Remove(src)
	if err != nil {
		return err
	}

	return nil
}
