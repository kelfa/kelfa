package filesystem

import (
	"os"
	"path/filepath"
	"time"
)

func (d *DP) DataEndTime() (*time.Time, error) {
	newest := "0000-01-01"

	err := filepath.Walk(d.folder, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".log" {
			return nil
		}
		if info.Name() > newest {
			newest = info.Name()
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	t, err := time.Parse("2006-01-02.log", newest)
	if err != nil {
		return nil, err
	}
	t = t.Add(time.Hour*24 - time.Second)
	return &t, nil
}
