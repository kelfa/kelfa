package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"go.kelfa.io/kelfa/pkg/analytics"
	"go.kelfa.io/kelfa/pkg/dal"
	"go.kelfa.io/kelfa/pkg/dal/objects"
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

	ds, err := dal.New("filesystem", objects.BackendOptions{Path: viper.GetString("data_folder")})
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
	a, err := analytics.New(&ds, *fromTime, *toTime)
	if err != nil {
		log.Fatalf("%v", err)
	}
	ss := a.GetSessions(viper.GetDuration("session_inactivity_timeout"))

	uas := UAS{}
	for _, s := range ss.Sessions {
		if s.Crawler {
			continue
		}
		uas.Add(*s.UserAgent)
	}
	sort.Slice(uas.UAs, func(i, j int) bool {
		return uas.UAs[i].Count < uas.UAs[j].Count
	})
	spew.Dump(uas)
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
