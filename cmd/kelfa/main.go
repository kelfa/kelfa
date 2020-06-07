package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.kelfa.io/pkg/lib"
)

var ds *lib.DataSource

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.kelfa")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: %s \n", err)
		return
	}
	ds, err = lib.New(fmt.Sprintf("%s/.kelfa/database.db", os.Getenv("HOME")))
	if err != nil {
		fmt.Printf("fatal error loading datasource: %s \n", err)
		return
	}

	rootCmd.PersistentFlags().StringP("from", "f", defaultFromDate(), "from")
	rootCmd.PersistentFlags().StringP("to", "t", defaultToDate(), "to")
	rootCmd.PersistentFlags().StringP("data", "d", "", "data folder")

	if err := viper.BindPFlag("data_folder", rootCmd.PersistentFlags().Lookup("data")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("period_from", rootCmd.PersistentFlags().Lookup("from")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("period_to", rootCmd.PersistentFlags().Lookup("to")); err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use: "kelfa",
}

func defaultFromDate() string {
	dd, _ := ds.DataEndTime()
	d := dd.AddDate(0, 0, -6)
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
}

func defaultToDate() string {
	d, _ := ds.DataEndTime()
	return d.Format(time.RFC3339)
}
