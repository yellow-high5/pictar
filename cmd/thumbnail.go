package cmd

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(thumbnailCmd)
}

var thumbnailCmd = &cobra.Command{
	Use:   "thumbnail",
	Short: "Scales the image up or down",
	Long:  "https://godoc.org/github.com/disintegration/imaging#Thumbnail",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		src, err := imaging.Open(filePath)
		if err != nil {
			log.Fatalf("No such file path: %v", filePath)
		}

		// TODO width and height should be alternative
		dst := imaging.Thumbnail(src, 200, 200, imaging.Gaussian)

		err = imaging.Save(dst, "./result.png")
		if err != nil {
			log.Fatalf("Failed to save image: %v", err)
		}

	},
}
