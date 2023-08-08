package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestOutline(t *testing.T) {
	_, err := captureStdout(testOutline)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	// fmt.Println("Pass")
}

func testOutline() error {
	input := `
<html>
<!-- This is a comment -->
<body>
	<p class="something" id="short"><span class="special">hi</span></p><br/>
</body>
</html>
`
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return err
	}
	forEachNode(doc, startElement, endElement)
	return nil
}

func captureStdout(f func() error) (string, error) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String(), err
}
