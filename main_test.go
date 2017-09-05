package file

import (
	"os"
	"path/filepath"
	"testing"
)

var TestDir = filepath.Join(os.TempDir(), "csvq_file_test")

func GetTestFilePath(filename string) string {
	return filepath.Join(TestDir, filename)
}

func TestMain(m *testing.M) {
	os.Exit(run(m))
}

func run(m *testing.M) int {
	defer teardown()

	setup()
	return m.Run()
}

func setup() {
	if _, err := os.Stat(TestDir); err == nil {
		os.RemoveAll(TestDir)
	}

	if _, err := os.Stat(TestDir); os.IsNotExist(err) {
		os.Mkdir(TestDir, 0755)
	}
}

func teardown() {
	if _, err := os.Stat(TestDir); err == nil {
		os.RemoveAll(TestDir)
	}
}
