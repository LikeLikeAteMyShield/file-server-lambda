package filefetch

import (
	"io"
	"bytes"
)

func FetchFileContent(filename string) (io.Reader, error) {
	return bytes.NewBufferString(""), nil
}
