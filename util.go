package iptables

import (
	"bytes"
	"fmt"
)

type LineBuffer struct {
	b     bytes.Buffer
	lines int
}

// Write takes a list of arguments, each a string or []string, joins all the
// individual strings with spaces, terminates with newline, and writes to buf.
// Any other argument type will panic.
func (buf *LineBuffer) Write(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			buf.b.WriteByte(' ')
		}
		switch x := arg.(type) {
		case string:
			buf.b.WriteString(x)
		case []string:
			for j, s := range x {
				if j > 0 {
					buf.b.WriteByte(' ')
				}
				buf.b.WriteString(s)
			}
		case []byte:
			buf.b.Write(x)
		default:
			panic(fmt.Sprintf("unknown argument type: %T", x))
		}
	}
	buf.b.WriteByte('\n')
	buf.lines++
}

func (buf *LineBuffer) NewWrite(strArrays ...[]string) {
	for i, strArray := range strArrays {
		if i > 0 {
			buf.b.WriteByte(' ')
		}
		for j, str := range strArray {
			if j > 0 {
				buf.b.WriteByte(' ')
			}
			buf.b.WriteString(str)

		}
	}
	buf.b.WriteByte('\n')
	buf.lines++
}

// WriteRune writes bytes to buffer, and terminates with newline.
func (buf *LineBuffer) WriteRune(r rune) {
	buf.b.WriteRune(r)
	//buf.b.WriteByte('\n')
	buf.lines++
}

// WriteString writes bytes to buffer, and terminates with newline.
func (buf *LineBuffer) WriteString(s string) {
	buf.b.WriteString(s)
	//buf.b.WriteByte('\n')
	buf.lines++
}

// Reset clears buf
func (buf *LineBuffer) Reset() {
	buf.b.Reset()
	buf.lines = 0
}

// Bytes returns the contents of buf as a []byte
func (buf *LineBuffer) Bytes() []byte {
	return buf.b.Bytes()
}
