package filesystem

import (
	"os"
	"path/filepath"
	"time"

	"go.kelfa.io/elf"
	"go.kelfa.io/kelfa/dal/objects"
)

type DP struct {
	folder string
	from   time.Time
	to     time.Time
	data   []map[string]string
}

// TODO: Check if folder exists
func New(bo objects.BackendOptions) (*DP, error) {
	var err error
	d := DP{
		folder: bo.Path,
		from:   bo.From,
		to:     bo.To,
	}
	d.data, err = d.getData()
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// TODO: Analyse if a signature change improves the situation
func (d *DP) getRelatedDays(from time.Time, to time.Time) []time.Time {
	uf := from.UTC()
	t := time.Date(uf.Year(), uf.Month(), uf.Day(), 0, 0, 0, 0, time.UTC)
	var o []time.Time
	for {
		if to.Before(t) {
			break
		}
		o = append(o, t)
		t = t.Add(time.Hour * 24)
	}
	return o
}

func (d *DP) getData() ([]map[string]string, error) {
	days := d.getRelatedDays(d.from, d.to)
	first := days[0]
	last := days[len(days)-1]
	data := make([]map[string]string, 0)
	for _, day := range days {
		f, err := os.Open(filepath.Join(d.folder, day.Format("2006-01-02.log")))
		if err != nil {
			return nil, err
		}
		r := elf.NewReader(f)
		logs, err := r.ReadAll()
		if err != nil {
			return nil, err
		}
		if day == first {
			fromTime := d.from.Format("15:04:05")
			filteredLogs := make([]map[string]string, 0)
			for _, log := range logs {
				if log["time"] >= fromTime {
					filteredLogs = append(filteredLogs, log)
				}
			}
			logs = filteredLogs
		}
		if day == last {
			toTime := d.to.Format("15:04:05")
			filteredLogs := make([]map[string]string, 0)
			for _, log := range logs {
				if log["time"] <= toTime {
					filteredLogs = append(filteredLogs, log)
				}
			}
			logs = filteredLogs
		}
		data = append(data, logs...)
	}
	return data, nil
}
