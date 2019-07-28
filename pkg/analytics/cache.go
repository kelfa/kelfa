package analytics

import "go.kelfa.io/kelfa/pkg/session"

func (a *Analytics) CacheStats() *session.CacheStats {
	cs := session.CacheStats{}
	for _, ds := range a.Data {
		for _, s := range ds.Sessions {
			cs.Add(s.CacheStats())
		}
	}
	return &cs
}
