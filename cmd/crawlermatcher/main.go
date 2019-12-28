package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/spf13/viper"
	"go.kelfa.io/pkg/analytics"
	"go.kelfa.io/pkg/crawlerflagger"
	"go.kelfa.io/pkg/dal"
	"go.kelfa.io/pkg/dal/objects"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.kelfa")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: %s \n", err)
		return
	}

	ds, err := dal.NewDataSource("filesystem", objects.BackendOptions{Path: viper.GetString("data_folder")})
	if err != nil {
		log.Fatalf("%v", err)
	}
	fromTime, err := ds.DataBeginTime()
	if err != nil {
		log.Fatalf("%v", err)
	}
	toTime, err := ds.DataEndTime()
	if err != nil {
		log.Fatalf("%v", err)
	}
	a, err := analytics.New(&ds, *fromTime, *toTime, analytics.None, time.Hour)
	if err != nil {
		log.Fatalf("%v", err)
	}

	uas := UAS{}
	for _, s := range a.Sessions {
		if s.Crawler {
			continue
		}
		if c := crawlerflagger.RegexpMatch(*s.UserAgent); c != nil {
			uas.Add(*s.UserAgent)
		}
	}
	sort.Slice(uas.UAs, func(i, j int) bool {
		return uas.UAs[i].Count < uas.UAs[j].Count
	})
	for _, ua := range uas.UAs {
		fmt.Printf("%v - %s\n", ua.Count, ua.UserAgent)
	}
}

type UA struct {
	UserAgent string
	Count     int
}

type UAS struct {
	UAs []UA
}

func (ss *UAS) Add(ua string) {
	for k, s := range ss.UAs {
		if s.UserAgent == ua {
			ss.UAs[k].Count++
			return
		}
	}
	ns := UA{
		UserAgent: ua,
		Count:     1,
	}
	ss.UAs = append(ss.UAs, ns)
}
