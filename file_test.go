package file

import (
	"os"
	"runtime"
	"testing"
)

func TestOpen(t *testing.T) {
	var err error

	notexistpath := GetTestFilePath("notexist.txt")
	_, err = OpenForRead(notexistpath)
	if err == nil {
		t.Fatal("no error, want IOError")
	}
	if _, ok := err.(*IOError); !ok {
		t.Fatal("error is not a IOError")
	}

	_, err = OpenForUpdate(notexistpath)
	if err == nil {
		t.Fatal("no error, want IOError")
	}
	if _, ok := err.(*IOError); !ok {
		t.Fatal("error is not a IOError")
	}

	_, err = OpenNBForRead(notexistpath)
	if err == nil {
		t.Fatal("no error, want IOError")
	}
	if _, ok := err.(*IOError); !ok {
		t.Fatal("error is not a IOError")
	}

	_, err = OpenNBForUpdate(notexistpath)
	if err == nil {
		t.Fatal("no error, want IOError")
	}
	if _, ok := err.(*IOError); !ok {
		t.Fatal("error is not a IOError")
	}

	switch runtime.GOOS {
	case "darwin", "dragonfly", "freebsd", "linux", "netbsd", "openbsd", "windows":
		Timeout = 0.1

		shpath := GetTestFilePath("lock_sh.txt")
		expath := GetTestFilePath("lock_ex.txt")

		shfp, _ := os.OpenFile(shpath, os.O_CREATE, 0600)
		shfp.Close()
		exfp, _ := os.OpenFile(expath, os.O_CREATE, 0600)
		exfp.Close()

		shfp1, err := OpenForRead(shpath)
		defer Close(shfp1)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		exfp1, err := OpenForUpdate(expath)
		defer Close(exfp1)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		shfp2, err := OpenNBForRead(shpath)
		defer Close(shfp2)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		exfp2, err := OpenNBForUpdate(expath)
		defer Close(exfp2)
		if err == nil {
			t.Fatal("no error, want error for duplicate exclusive lock")
		}
		if _, ok := err.(*TimeoutError); !ok {
			t.Fatal("error is not a TimeoutError")
		}

		err = Lock(nil, SHARED_LOCK)
		if err == nil {
			t.Fatal("no error, want error for invalid file descriptor")
		}
		if _, ok := err.(*LockError); !ok {
			t.Fatal("error is not a LockError")
		}
		err = TryLock(nil, SHARED_LOCK)
		if err == nil {
			t.Fatal("no error, want error for invalid file descriptor")
		}
		if _, ok := err.(*LockError); !ok {
			t.Fatal("error is not a LockError")
		}
	}
}

func TestClose(t *testing.T) {
	path := GetTestFilePath("closetest.txt")
	fp, _ := OpenForCreate(path)
	err := fp.Close()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	err = Close(fp)
	if err == nil {
		t.Fatal("no error, want error for invalid file descriptor")
	}
	if _, ok := err.(*LockError); !ok {
		t.Fatal("error is not a LockError")
	}
}
