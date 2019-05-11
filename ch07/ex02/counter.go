package counter

import (
	"io"
)

type wrapper struct {
	inner io.Writer
	count *int64
}

func (wrap *wrapper) Write(p []byte) (int, error) {
	*wrap.count += int64(len(p))
	return wrap.inner.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var count int64
	wrap := &wrapper{inner: w, count: &count}

	return wrap, &count
}
