package analytics

import "go.kelfa.io/pkg/session"

func (a *Analytics) DataByResourceFormat(ignoreErrors bool) *session.ResourceTypes {
	rt := session.ResourceTypes{}
	for _, s := range a.Sessions {
		rt.Add(s.DataByResourceFormat(ignoreErrors))
	}
	return &rt
}
