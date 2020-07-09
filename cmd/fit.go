package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fitCmd)
}

var fitCmd = &cobra.Command{
	Use:   "fit",
	Short: "Fit the specified maximum width and height and returns the transform",
	Long:  "https://godoc.org/github.com/disintegration/imaging#Fit",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		width, _ := strconv.Atoi(args[0])
		height, _ := strconv.Atoi(args[1])
		filePath := args[2]
		src, err := imaging.Open(filePath)
		if err != nil {
			log.Fatalf("No such file path: %v", filePath)
		}

		// TODO: width and height should be alternative
		dst := imaging.Fit(src, width, height, imaging.Gaussian)

		err = imaging.Save(dst, fmt.Sprintf("./result.%s", cmd.Flags().Lookup("ext").Value))
		if err != nil {
			log.Fatalf("Failed to save image: %v", err)
		}

	},
}
