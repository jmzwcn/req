package req

import (
	"bytes"
	"github.com/jmzwcn/req/v3/internal/tests"
	"log"
	"testing"
)

func TestLogger(t *testing.T) {
	buf := new(bytes.Buffer)
	l := NewLogger(buf, "", log.Ldate|log.Lmicroseconds)
	c := tc().SetLogger(l)
	c.SetProxyURL(":=\\<>ksfj&*&sf")
	tests.AssertContains(t, buf.String(), "error", true)
	buf.Reset()
	c.R().SetOutput(nil)
	tests.AssertContains(t, buf.String(), "warn", true)
}
