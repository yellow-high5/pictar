package cmd

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cropCmd)
}

var cropCmd = &cobra.Command{
	Use:   "crop",
	Short: "Cuts out a rectangular region.",
	Long:  "https://godoc.org/github.com/disintegration/imaging#Crop",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		src, err := imaging.Open(filePath)
		if err != nil {
			log.Fatalf("No such file path: %v", filePath)
		}

		// TODO: crop size and anchor should be alternative
		dst := imaging.CropCenter(src, 100, 100)

		err = imaging.Save(dst, "./result.png")
		if err != nil {
			log.Fatalf("Failed to save image: %v", err)
		}

	},
}
