package xstrings

import "bytes"

func NullTerminated(data []byte) string {
	index := bytes.IndexByte(data, 0)
	if index < 0 {
		index = len(data)
	}

	return string(data[:index])
}
