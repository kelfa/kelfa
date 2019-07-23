package crawlerflagger_test

import (
	"testing"

	"go.kelfa.io/kelfa/pkg/crawlerflagger"
)

func BenchmarkExactMatch(b *testing.B) {
	uas := []struct {
		name string
		ua   string
	}{
		{"first", "Googlebot/2.1 (+http://www.google.com/bot.html)"},
		{"line1000", "ZoominfoBot (zoominfobot at zoominfo dot com)"},
		{"line2000", "Mozilla/5.0 (Windows NT 6.1; WOW64) SkypeUriPreview Preview/0.5"},
		{"line3000", "Companybook-Crawler (+https://www.companybooknetworking.com/)"},
		{"line4000", "Mozilla/5.0 (compatible; Nmap Scripting Engine; https://nmap.org/book/nse.html)"},
		{"miss", "non-existant-uas"},
	}
	for _, ua := range uas {
		b.Run(ua.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c := crawlerflagger.ExactMatch(ua.ua)
				_ = c
			}
		})
	}
}
