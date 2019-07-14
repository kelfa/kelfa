package filesystem

import (
	"os"
	"path/filepath"
	"time"
)

func (d *DP) DataBeginTime() (*time.Time, error) {
	oldest := "9999-12-31.log"

	err := filepath.Walk(d.folder, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".log" {
			return nil
		}
		if info.Name() < oldest {
			oldest = info.Name()
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	t, err := time.Parse("2006-01-02.log", oldest)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
