/*
Package cmd
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/iliaskaras/fare-estimation/app/files"
	baseAppErrors "github.com/iliaskaras/fare-estimation/app/infrastructure/errors"
	"github.com/spf13/cobra"
	"os"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleans out erroneous coordinates from a file containing ride information.",
	Long: `Detects and filters out invalid coordinate entries from a file containing information about rides.
The command produce a new filtered file with the invalid entries filtered out.

File input: A file that consisting of a list of tuples of the form (id_ride, lat, lng, timestamp),
representing the position of the taxi-cab during a ride. 
Two consecutive tuples in this file, p1 & p2 forms a segment S, which represent a single ride.

Invalid entry: Is considered the second part of a segment's tuple, more specifically the p2 row,
that the calculated segment's speed U is above 100km per hour.`,
	Run: func(cmd *cobra.Command, args []string) {

		filePath, _ := cmd.Flags().GetString("filepath")

		fileService, err := files.GetFileService(filePath)

		if err != nil {
			if errors.Is(err, files.UnsupportedFileType) {
				fmt.Printf(err.Error())
				os.Exit(1)
			}
		}

		err = fileService.Read(filePath)
		if err != nil {
			if errors.Is(err, baseAppErrors.InvalidInputError) {
				fmt.Printf(err.Error())
			}
			fmt.Printf(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().StringP(
		"filepath", "f", "", "The file path contains the ride information",
	)
}
