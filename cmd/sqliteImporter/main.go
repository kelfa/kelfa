package main

import (
	"fmt"
	"log"

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
	sql, err := dal.NewWritableDataSource("sqlite", objects.BackendOptions{})
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
	for _, dp := range dps {
		if err := sql.AddDataPoint(dp); err != nil {
			log.Fatalf("%v", err)
		}
	}
}
