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
		default:
			panic(fmt.Sprintf("unknown argument type: %T", x))
		}
	}
	buf.b.WriteByte('\n')
	buf.lines++
}

// WriteBytes writes bytes to buffer, and terminates with newline.
func (buf *LineBuffer) WriteBytes(bytes []byte) {
	buf.b.Write(bytes)
	buf.b.WriteByte('\n')
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

// Lines returns the number of lines in buf. Note that more precisely, this returns the
// number of times Write() or WriteBytes() was called; it assumes that you never wrote
// any newlines to the buffer yourself.
func (buf *LineBuffer) Lines() int {
	return buf.lines
}
