package parse

import (
	"io"
	"os"
	"strings"
)

func Read(filePath string) (io.ReadCloser, error) {
	file := strings.Replace(filePath, "/", string(os.PathSeparator), -1)
	return os.Open(file)
}
