package filesystem

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
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

func (d *DP) GetDataPoints(from time.Time, to time.Time) (datapoints []objects.DataPoint, err error) {
	logs, err := d.getData(from, to)
	if err != nil {
		return nil, err
	}
	for _, log := range logs {
		dt, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT%sZ", log["date"], log["time"]))
		if err != nil {
			return nil, err
		}
		et, err := time.ParseDuration(fmt.Sprintf("%ss", log["time-taken"]))
		if err != nil {
			return nil, err
		}
		requestSize, err := strconv.Atoi(log["cs-bytes"])
		if err != nil {
			return nil, err
		}
		responseSize, err := strconv.Atoi(log["sc-bytes"])
		if err != nil {
			return nil, err
		}
		responseCode, err := strconv.Atoi(log["sc-status"])
		if err != nil {
			return nil, err
		}
		datapoints = append(datapoints, objects.DataPoint{
			ID:                       log["x-edge-request-id"],
			DateTime:                 dt,
			ElapsedTime:              et,
			ClientIP:                 net.ParseIP(log["c-ip"]),
			ClientUserAgent:          log["cs(User-Agent)"],
			RequestDomain:            log["x-host-header"],
			RequestURI:               log["cs-uri-stem"],
			RequestURIQuery:          log["cs-uri-query"],
			RequestReferer:           log["cs(Referer)"],
			RequestMethod:            log["cs-method"],
			RequestCookie:            log["cs(Cookie)"],
			RequestSize:              requestSize,
			CommunicationProtocol:    log["cs-protocol"],
			CommunicationTLSProtocol: log["ssl-protocol"],
			CommunicationTLSCipher:   log["ssl-cipher"],
			ResponseSize:             responseSize,
			ResponseCode:             responseCode,
			CDNLocation:              log["x-edge-location"],
			CDNHost:                  log["cs(Host)"],
			CDNCache:                 log["x-edge-result-type"],
		})
	}
	return datapoints, nil
}
