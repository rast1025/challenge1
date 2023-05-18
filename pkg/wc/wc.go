package wc

import (
	"bytes"
	"io"
	"strings"
	"unicode/utf8"
)

func GetByteCount(data []byte) int {
	return len(data)
}

func GetLineCount(data []byte) int {
	return countLines(bytes.NewReader(data))
}

func GetWordCount(data []byte) int {
	return len(strings.Fields(string(data)))
}

func GetCharCount(data []byte) int {
	return utf8.RuneCount(data)
}

func countLines(r io.Reader) int {
	var count, read int
	var err error
	var targetByte = []byte("\n")
	buffer := make([]byte, 32*1024)

	for {
		read, err = r.Read(buffer)
		if err != nil {
			break
		}
		count += bytes.Count(buffer[:read], targetByte)
	}

	if err == io.EOF {
		return count
	}
	panic(err)
}
