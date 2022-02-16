/*
Package cmd
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fare-estimation",
	Short: "A tool for estimating ride fares",
	Long: `A Fare Estimation tool that provides command for producing the
fare estimations of rides.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
