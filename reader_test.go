package elff_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/selffa/elff"
)

var examples = []string{
	`#Version: 1.0
#Date: 12-Jan-1996 00:00:00
#Fields: time cs-method cs-uri
00:34:23 GET /foo/bar.html
12:21:16 GET /foo/bar.html
12:45:52 GET /foo/bar.html
12:57:34 GET /foo/bar.html`,
	`#Version: 1.0
#Fields: date time x-edge-location sc-bytes c-ip cs-method cs(Host) cs-uri-stem sc-status cs(Referer) cs(User-Agent) cs-uri-query cs(Cookie) x-edge-result-type x-edge-request-id x-host-header cs-protocol cs-bytes time-taken x-forwarded-for ssl-protocol ssl-cipher x-edge-response-result-type cs-protocol-version fle-status fle-encrypted-fields
2019-06-03      08:53:51        BOS50-C1        708     159.65.224.212  GET     d2bau79chpbdbg.cloudfront.net   /index.xml      304     -       NewsBlur%2520Feed%2520Fetcher%2520-%25201%2520subscriber%2520-%2520http://www.newsblur.com/site/7072202/fabio-alessandro-fale-locatis-blog%2520(Mozilla/5.0%2520(Macintosh;%2520Intel%2520Mac%2520OS%2520X%252010_12_3)%2520AppleWebKit/537.36%2520(KHTML,%2520like%2520Gecko)%2520Chrome/56.0.2924.87%2520Safari/537.36)     -       -       Hit     OpEFs4f8-BiSQBK9fWLCyc8ueiLtORLV9l-4TY_DpngcR7JWDmeUkg==        fale.io https   591   0.014    -       TLSv1.2 ECDHE-RSA-AES128-GCM-SHA256     Hit     HTTP/1.1        -       -
2019-06-03      08:53:51        BOS50-C1        713     159.65.224.212  GET     d2bau79chpbdbg.cloudfront.net   /index.xml      304     -       NewsBlur%2520Feed%2520Fetcher%2520-%25201%2520subscriber%2520-%2520http://www.newsblur.com/site/7072202/fabio-alessandro-fale-locatis-blog%2520(Mozilla/5.0%2520(Macintosh;%2520Intel%2520Mac%2520OS%2520X%252010_12_3)%2520AppleWebKit/537.36%2520(KHTML,%2520like%2520Gecko)%2520Chrome/56.0.2924.87%2520Safari/537.36)     -       -       Hit     zPyO_tGB_BEqcjqHfDuX0lUXA2t6jj5alpzPdw2Vd1ZosKxiaAUk3A==        fale.io https   559   0.028    -       TLSv1.2 ECDHE-RSA-AES128-GCM-SHA256     Hit     HTTP/1.1        -       -`,
}

func TestReadHeaders(t *testing.T) {
	for _, example := range examples {
		r := elff.NewReader(strings.NewReader(example))
		d, err := r.ReadAll()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(r.Version)
		fmt.Println(r.Fields)
		fmt.Println(r.Date)
		spew.Dump(d)
	}
}
