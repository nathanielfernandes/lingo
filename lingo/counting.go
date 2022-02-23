package lingo

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func LineCounter(r io.Reader) (uint32, error) {
	var count uint32 = 0
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}
	return count, nil
}

func GetLineCount(fp string) uint32 {
	f, err := os.Open(fp)
	defer f.Close()

	if err != nil {
		return 0
	}

	c, err := LineCounter(f)
	if err != nil {
		return c
	}

	return c
}
