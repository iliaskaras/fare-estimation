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
	Long: `A Fare Estimation tool that provides command utilities for:

1. Cleans out erroneous coordinates out of ride coordinate content files.
2. Evaluate the ride fares.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
