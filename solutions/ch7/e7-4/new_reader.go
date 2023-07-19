// Exercise 7.4: The strings.NewReader function returns a value that satisfies the io.Reader
// interface (and others) by reading from its argument, a string. Implement a simple version of
// NewReader yourself, and use it to make the HTML parser (§5.2) take input from a string.

package new_reader

import "io"

type CustomReader struct {
	s string
}

func (r CustomReader) Read(p []byte) (n int, e error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		e = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &CustomReader{s}
}
