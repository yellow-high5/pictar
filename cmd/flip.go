package cmd

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(flipCmd)
}

var flipCmd = &cobra.Command{
	Use:   "flip",
	Short: "Flips the image.",
	Long:  "https://godoc.org/github.com/disintegration/imaging#FlipH or https://godoc.org/github.com/disintegration/imaging#FlipV",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		src, err := imaging.Open(filePath)
		if err != nil {
			log.Fatalf("No such file path: %v", filePath)
		}

		// TODO: flip direction should be alternative(horizontally or vertically)
		dst := imaging.FlipH(src)

		err = imaging.Save(dst, "./result.png")
		if err != nil {
			log.Fatalf("Failed to save image: %v", err)
		}

	},
}
