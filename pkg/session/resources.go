package session

import (
	"net/http"

	"github.com/kelfa/kelfa/pkg/filepath"
)

type ResourceTypes map[string]int

func (r *ResourceTypes) Add(a *ResourceTypes) {
	for k, v := range *a {
		if _, ok := (*r)[k]; ok {
			(*r)[k] = (*r)[k] + v
		} else {
			(*r)[k] = v
		}
	}
}

func (s *Session) DataByResourceFormat(ignoreErrors bool) *ResourceTypes {
	exts := ResourceTypes{}
	for _, log := range s.DataPoints {
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
	return &exts
}
