package uniq

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
3
4
5`

var testOkRes = `1
2
3
4
5
`

var testFail = `
1
2
45
3`

func TestOk(t *testing.T) {

	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := uniq(in, out)

	if err != nil {
		t.Errorf("test for OK failed - error")
	}

	if out.String() != testOkRes {
		t.Errorf("test for OK failed - results not match")
	}
}

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := uniq(in, out)

	if err == nil {
		t.Errorf("test for OK failed - error")
	}
}
