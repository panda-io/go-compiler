package ir

import (
	"fmt"
	"io"
)

// --- [ Formatted I/O writer ] ------------------------------------------------

// FmtWriter is a formatted I/O writer.
//
// A formatted I/O writer keeps track of the total number of bytes written to w
// and the first non-nil error encountered.
type FmtWriter struct {
	// underlying io.Writer.
	w io.Writer
	// Number of bytes written to w.
	size int64
	// First non-nil error encountered.
	err error
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string. It returns the
// number of bytes written and any write error encountered.
func (fw *FmtWriter) Fprint(a ...interface{}) (n int, err error) {
	if fw.err != nil {
		// early return if a previous error has been encountered.
		return 0, nil
	}
	n, err = fmt.Fprint(fw.w, a...)
	fw.size += int64(n)
	fw.err = err
	return n, err
}

// Fprintf formats according to a format specifier and writes to w. It returns
// the number of bytes written and any write error encountered.
func (fw *FmtWriter) Fprintf(format string, a ...interface{}) (n int, err error) {
	if fw.err != nil {
		// early return if a previous error has been encountered.
		return 0, nil
	}
	n, err = fmt.Fprintf(fw.w, format, a...)
	fw.size += int64(n)
	fw.err = err
	return n, err
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended. It
// returns the number of bytes written and any write error encountered.
func (fw *FmtWriter) Fprintln(a ...interface{}) (n int, err error) {
	if fw.err != nil {
		// early return if a previous error has been encountered.
		return 0, nil
	}
	n, err = fmt.Fprintln(fw.w, a...)
	fw.size += int64(n)
	fw.err = err
	return n, err
}
