package mtalk

import (
	"fmt"
	"net"
	"os"
	"path"
	"testing"
	"time"
)

const (
	TMP_DIR   = "tmp"
	FILE_NAME = "foo.txt"
)

func TestMtalk(t *testing.T) {
	setupFile(t)

	go Listen(8081)

	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		t.Fatalf("conn error: %v", err)
	}

	filePath := fooPath()
	cmd := fmt.Sprintf("cmd rm %s \n", filePath)
	conn.Write([]byte(cmd))
	// see a better way to wait
	time.Sleep(2 * time.Second)

	_, err = os.Stat(filePath)
	if !os.IsNotExist(err) {
		t.Errorf("%s was not removed", FILE_NAME)
	}
}

func setupFile(t *testing.T) {
	t.Helper()
	_, err := os.Stat(TMP_DIR)
	if os.IsNotExist(err) {
		err := os.Mkdir(TMP_DIR, 0755)
		if err != nil {
			t.Fatalf("mkdir error: %v", err)
		}
	}
	f, err := os.Create(fooPath())
	if err != nil {
		t.Fatalf("%s was not created", FILE_NAME)
	}
	f.Close()
}

func fooPath() string {
	return path.Join(TMP_DIR, string(os.PathSeparator), FILE_NAME)
}
