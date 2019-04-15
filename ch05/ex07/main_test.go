package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/andreyvit/diff"
)

func TestSimpleDOM(t *testing.T) {
	in = strings.NewReader("<h1>title</h1><img src='path/to/image.png'></img>")
	out = new(bytes.Buffer)
	pretty()
	expected := `<html>
  <head/>
  <body>
    <h1>
      title
    </h1>
    <img src="path/to/image.png"/>
  </body>
</html>
`
	got := out.(*bytes.Buffer).String()
	if got != expected {
		t.Errorf("not expected:\n%v", diff.LineDiff(expected, got))
	}
}
