package formatter

import (
	"io"
)

// JSON is a writer that formats/colorizes JSON without decoding it.
// If the stream of bytes does not start with {, the formatting is disabled.
type JSON struct {
	Out       io.Writer
	Scheme    ColorScheme
	inited    bool
	disabled  bool
	last      byte
	lastQuote byte
	isValue   bool
	level     int
	buf       []byte
}

func (j *JSON) Write(p []byte) (n int, err error) {
	var result = Color(Pretty(p), nil)
	n, err = j.Out.Write(result)
	if err != nil || n != len(result) {
		return
	}
	return len(p), nil
}
