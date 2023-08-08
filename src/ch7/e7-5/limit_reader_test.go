package limit_reader

import (
	"fmt"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	s := "Hello, Reader!"
	readLength := 6
	reader := strings.NewReader(s)
	limited := LimitReader(reader, int64(readLength))

	b := make([]byte, 64)
	n, err := limited.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	if readLength != n {
		t.Fatalf("Read %d instead of %d bytes", n, readLength)
	}
	fmt.Println(string(b))
}

func TestLimitReader2(t *testing.T) {
	s := "Hello, Reader!"
	readLength := len(s) + 1
	reader := strings.NewReader(s)
	limited := LimitReader(reader, int64(readLength))

	b := make([]byte, 64)
	n, err := limited.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(s) {
		t.Fatalf("Read %d instead of %d bytes", n, readLength)
	}
	fmt.Println(string(b))
}
