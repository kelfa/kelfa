package main

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/spf13/viper"
	"github.com/kelfa/kelfa/pkg/analytics"
	"github.com/kelfa/kelfa/pkg/dal"
	"github.com/kelfa/kelfa/pkg/dal/objects"
	"github.com/kelfa/kelfa/pkg/session"
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

	grgxp := regexp.MustCompile(`DuckDuck`)
	gbdps := make(DP)
	for i := range a.Sessions {
		if grgxp.MatchString(*a.Sessions[i].UserAgent) {
			gbdps.Add(&a.Sessions[i])
		}
	}

	for ua, gbdp := range gbdps {
		fmt.Printf("%s\n", ua)
		for ip, amount := range gbdp {
			fmt.Printf("  %v - %s\n", amount, ip)
		}
	}
}

type DP map[string]map[string]int

func (dp *DP) Add(s *session.Session) {
	for k, as := range *dp {
		if *s.UserAgent == k {
			for ip, amount := range as {
				if ip == s.IP.String() {
					(*dp)[k][ip] = amount + 1
					return
				}
				(*dp)[k][s.IP.String()] = 1
				return
			}
		}
	}
	(*dp)[*s.UserAgent] = map[string]int{s.IP.String(): 1}
}
