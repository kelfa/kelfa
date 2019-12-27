package main

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"time"

	"github.com/spf13/viper"
	"go.kelfa.io/kelfa/pkg/analytics"
	"go.kelfa.io/kelfa/pkg/dal"
	"go.kelfa.io/kelfa/pkg/dal/objects"
	"go.kelfa.io/kelfa/pkg/session"
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

	grgxp := regexp.MustCompile(`Googlebot`)
	hrgxp := regexp.MustCompile(`(\.googlebot\.com\.|\.googleusercontent\.com\.)$`)
	gbdps := make(DP)
	for _, s := range a.Sessions {
		if grgxp.MatchString(*s.UserAgent) {
			gbdps.Add(&s)
		}
	}

	var total int
	var totalValid int
	for ua, gbdp := range gbdps {
		var valid int
		var unmatched []string
		for _, dp := range gbdp {
			host, err := net.LookupAddr(dp.IP.String())
			if err != nil {
				fmt.Printf("error: %v", host)
			}
			if len(host) == 0 {
				unmatched = append(unmatched, dp.IP.String())
				continue
			}
			if hrgxp.MatchString(host[0]) {
				valid++
			} else {
				unmatched = append(unmatched, fmt.Sprintf("%s (%s)", dp.IP.String(), host[0]))
			}
		}
		fmt.Printf("User Agent: %s\nPassed Google validation:%v/%v\n", ua, valid, len(gbdp))
		if len(unmatched) > 0 {
			fmt.Println("Unmatched hosts:")
			for _, uh := range unmatched {
				fmt.Printf("    %v\n", uh)
			}
		}
		fmt.Println("---")
		totalValid = totalValid + valid
		total = total + len(gbdp)
	}
	fmt.Printf("Total: %v/%v\n", totalValid, total)
}

type DP map[string][]Session

type Session struct {
	Date time.Time
	IP   net.IP
	Ok   bool
}

func (dp *DP) Add(s *session.Session) {
	for k, as := range *dp {
		if *s.UserAgent == k {
			(*dp)[k] = append(as, Session{
				Date: *s.Begin,
				IP:   *s.IP,
			})
			return
		}
	}
	(*dp)[*s.UserAgent] = []Session{
		{
			Date: *s.Begin,
			IP:   *s.IP,
		},
	}
}
