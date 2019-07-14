package analytics

import (
	"net/http"

	"go.kelfa.io/kelfa/pkg/filepath"
)

func (a *Analytics) DataByResourceFormat(ignoreErrors bool) map[string]int {
	exts := make(map[string]int)
	for _, log := range a.dataPoints {
		if ignoreErrors {
			if log.ResponseCode != http.StatusOK {
				continue
			}
		}
		e := filepath.Ext(log.RequestURI, "html")
		if val, ok := exts[e]; ok {
			exts[e] = val + 1
		} else {
			exts[e] = 1
		}
	}
	return exts
}
