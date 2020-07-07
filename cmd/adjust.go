package cmd

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(adjustCmd)
}

var adjustCmd = &cobra.Command{
	Use:   "adjust",
	Short: "Adjust saturation, contrast, brightness, gamma, sigmoid, LUT",
	Long:  "https://godoc.org/github.com/disintegration/imaging",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		src, err := imaging.Open(filePath)
		if err != nil {
			log.Fatalf("No such file path: %v", filePath)
		}

		// TODO: parameter should be specified
		dst := imaging.AdjustSaturation(src, -20)

		err = imaging.Save(dst, "./result.png")
		if err != nil {
			log.Fatalf("Failed to save image: %v", err)
		}

	},
}
