package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

	rootCmd.PersistentFlags().StringP("from", "f", defaultFromDate(), "from")
	rootCmd.PersistentFlags().StringP("to", "t", defaultToDate(), "to")
	rootCmd.PersistentFlags().StringP("data", "d", "", "data folder")

	viper.BindPFlag("data_folder", rootCmd.PersistentFlags().Lookup("data"))
	viper.BindPFlag("period_from", rootCmd.PersistentFlags().Lookup("from"))
	viper.BindPFlag("period_to", rootCmd.PersistentFlags().Lookup("to"))
}

var rootCmd = &cobra.Command{
	Use: "kelfa",
}

func defaultFromDate() string {
	d := time.Now().AddDate(0, 0, -8)
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
}

func defaultToDate() string {
	d := time.Now().AddDate(0, 0, -2)
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, time.UTC).Format(time.RFC3339)
}
