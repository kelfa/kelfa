package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
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
	fs, err := dal.New("filesystem", objects.BackendOptions{Path: viper.GetString("data_folder")})
	if err != nil {
		log.Fatalf("%v", err)
	}
	sql, err := dal.NewWritableDataSource("sqlite", objects.BackendOptions{Path: fmt.Sprintf("%s/.kelfa/database.db", os.Getenv("HOME"))})
	if err != nil {
		log.Fatalf("%v", err)
	}
	fromTime, err := fs.DataBeginTime()
	if err != nil {
		log.Fatalf("%v", err)
	}
	toTime, err := fs.DataEndTime()
	if err != nil {
		log.Fatalf("%v", err)
	}
	dps, err := fs.GetDataPoints(*fromTime, *toTime)
	if err != nil {
		log.Fatalf("%v", err)
	}
	dpssql, err := sql.GetDataPoints(*fromTime, *toTime)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, dp := range dps {
		var present bool
		for _, dpsql := range dpssql {
			if dp.ID == dpsql.ID {
				present = true
				break
			}
		}
		if !present {
			fmt.Printf("Going to add %s %s...", dp.DateTime, dp.ID)
			if err := sql.AddDataPoint(dp); err != nil {
				log.Fatalf("%v", err)
			}
			fmt.Printf(" done\n")
		} else {
			fmt.Printf("Not going to add %s %s. Already present\n", dp.DateTime, dp.ID)
		}
	}
}
