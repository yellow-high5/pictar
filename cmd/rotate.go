package cmd

import (
	"image/color"
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rotateCmd)
}

var rotateCmd = &cobra.Command{
	Use:   "rotate",
	Short: "Rotates an image by the given angle counter-clockwise.",
	Long:  "https://godoc.org/github.com/disintegration/imaging#Rotate",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		src, err := imaging.Open(filePath)
		if err != nil {
			log.Fatalf("No such file path: %v", filePath)
		}

		// TODO: angle and bgcolor should be alternative
		dst := imaging.Rotate(src, 90, color.Transparent)

		err = imaging.Save(dst, "./result.png")
		if err != nil {
			log.Fatalf("Failed to save image: %v", err)
		}

	},
}
