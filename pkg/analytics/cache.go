package analytics

import "go.kelfa.io/kelfa/pkg/session"

func (a *Analytics) CacheStats() *session.CacheStats {
	cs := session.CacheStats{}
	for _, s := range a.Sessions {
		cs.Add(s.CacheStats())
	}
	return &cs
}
