package tests

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/datastore"
)

func NewTestApp(optTestDataDir ...string) (*string, error) {
	var testDataDir string
	if len(optTestDataDir) == 0 || optTestDataDir[0] == "" {
		// fallback to the default test data directory
		_, currentFile, _, _ := runtime.Caller(0)
		testDataDir = filepath.Join(path.Dir(currentFile), "data")
	} else {
		testDataDir = optTestDataDir[0]
	}

	tempDir, err := TempDirClone(testDataDir)
	if err != nil {
		return nil, err
	}

	// load data dir and db connections
	if err := core.Bootstrap(tempDir); err != nil {
		return nil, err
	}
	datastore.InitTestDB(tempDir)
	return &tempDir, nil
}

// TempDirClone creates a new temporary directory copy from the
// provided directory path.
//
// It is the caller's responsibility to call `os.RemoveAll(tempDir)`
// when the directory is no longer needed!
func TempDirClone(dirToClone string) (string, error) {
	tempDir, err := os.MkdirTemp("", "purebackend_test_*")
	if err != nil {
		return "", err
	}

	// copy everything from testDataDir to tempDir
	if err := copyDir(dirToClone, tempDir); err != nil {
		return "", err
	}

	return tempDir, nil
}

// -------------------------------------------------------------------
// Helpers
// -------------------------------------------------------------------

func copyDir(src string, dest string) error {
	if err := os.MkdirAll(dest, os.ModePerm); err != nil {
		return err
	}

	sourceDir, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceDir.Close()

	items, err := sourceDir.Readdir(-1)
	if err != nil {
		return err
	}

	for _, item := range items {
		fullSrcPath := filepath.Join(src, item.Name())
		fullDestPath := filepath.Join(dest, item.Name())

		var copyErr error
		if item.IsDir() {
			copyErr = copyDir(fullSrcPath, fullDestPath)
		} else {
			copyErr = copyFile(fullSrcPath, fullDestPath)
		}

		if copyErr != nil {
			return copyErr
		}
	}

	return nil
}

func copyFile(src string, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return err
	}

	return nil
}
