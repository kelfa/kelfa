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
		{"case0", "Googlebot/2.1 (+http://www.google.com/bot.html)"},
		{"case101", "Summify (Summify/1.0.1; +http://summify.com)"},
		{"case200", "www.deadlinkchecker.com XMLHTTP/1.0"},
		{"case300", "Mozilla/5.0 (compatible; BomboraBot/1.0; +http://www.bombora.com/bot)"},
		{"case400", "Clickagy Intelligence Bot v2"},
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

func BenchmarkRegExpMatch(b *testing.B) {
	uas := []struct {
		name string
		ua   string
	}{
		{"case0", "Googlebot/2.1 (+http://www.google.com/bot.html)"},
		{"case101", "Summify (Summify/1.0.1; +http://summify.com)"},
		{"case200", "www.deadlinkchecker.com XMLHTTP/1.0"},
		{"case300", "Mozilla/5.0 (compatible; BomboraBot/1.0; +http://www.bombora.com/bot)"},
		{"case400", "Clickagy Intelligence Bot v2"},
		{"miss", "non-existant-uas"},
	}
	for _, ua := range uas {
		b.Run(ua.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c := crawlerflagger.RegexpMatch(ua.ua)
				_ = c
			}
		})
	}
}
