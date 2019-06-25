package elff_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/selffa/elff"
)

var ex = `#Version: 1.0
#Date: 12-Jan-1996 00:00:00
#Fields: time cs-method cs-uri
00:34:23 GET /foo/bar.html
12:21:16 GET /foo/bar.html
12:45:52 GET /foo/bar.html
12:57:34 GET /foo/bar.html`

func TestReadHeaders(t *testing.T) {
	r := elff.NewReader(strings.NewReader(ex))
	_, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.Version)
	fmt.Println(r.Fields)
	fmt.Println(r.Date)
}
