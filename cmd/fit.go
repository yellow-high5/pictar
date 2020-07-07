package cmd

import (
	"log"

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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		src, err := imaging.Open(filePath)
		if err != nil {
			log.Fatalf("No such file path: %v", filePath)
		}

		// TODO: width and height should be alternative
		dst := imaging.Fit(src, 200, 100, imaging.Gaussian)

		err = imaging.Save(dst, "./result.png")
		if err != nil {
			log.Fatalf("Failed to save image: %v", err)
		}

	},
}
