package mtalk

import (
	"net"
	"os"
	"testing"
	"time"
)

func TestMtalk(t *testing.T) {
	f, err := os.Create("tmp/foo.txt")
	if err != nil {
		t.Fatal("foo.txt was not created")
	}
	f.Close()

	go Listen(8081)

	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		t.Fatalf("conn error: %v", err)
	}

	conn.Write([]byte("cmd rm tmp/foo.txt \n"))
	// see a better way to wait
	time.Sleep(2 * time.Second)

	_, err = os.Stat("tmp/foo.txt")
	if !os.IsNotExist(err) {
		t.Errorf("foo.txt was not removed")
	}
}
