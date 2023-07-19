// Exercise 7.5: The LimitReader function in the io package accepts an io.Reader r and a
// number of bytes n, and returns another Reader that reads from r but reports an end-of-file
// condition after n bytes. Implement it.
// func LimitReader(r io.Reader, n int64) io.Reader

package limit_reader

import "io"

type LimitedReader struct {
	r              io.Reader
	remainingBytes int64
}

func (lr LimitedReader) Read(data []byte) (n int, err error) {
	if len(data) > int(lr.remainingBytes) {
		return lr.r.Read(data[:lr.remainingBytes])
	}
	return lr.r.Read(data)
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
