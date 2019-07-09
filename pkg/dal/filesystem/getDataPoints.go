package filesystem

import (
	"os"
	"path/filepath"
	"time"

	"go.kelfa.io/elf"
	"go.kelfa.io/kelfa/pkg/dal/objects"
)

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

func (d *DP) getData(from time.Time, to time.Time) ([]map[string]string, error) {
	days := d.getRelatedDays(from, to)
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
			fromTime := from.Format("15:04:05")
			filteredLogs := make([]map[string]string, 0)
			for _, log := range logs {
				if log["time"] >= fromTime {
					filteredLogs = append(filteredLogs, log)
				}
			}
			logs = filteredLogs
		}
		if day == last {
			toTime := to.Format("15:04:05")
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

func (d *DP) GetDataPoints(from time.Time, to time.Time) ([]objects.DataPoint, error) {
	logs, err := d.getData(from, to)
	if err != nil {
		return nil, err
	}
	for _, log := range logs {

	}
	return nil, nil
}
