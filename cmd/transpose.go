package cmd

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(transposeCmd)
}

var transposeCmd = &cobra.Command{
	Use:   "transpose",
	Short: "Flips the image horizontally and rotates 90 degrees counter-clockwise.",
	Long:  "https://godoc.org/github.com/disintegration/imaging#Transpose",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		src, err := imaging.Open(filePath)
		if err != nil {
			log.Fatalf("No such file path: %v", filePath)
		}

		dst := imaging.Transpose(src)

		err = imaging.Save(dst, "./result.png")
		if err != nil {
			log.Fatalf("Failed to save image: %v", err)
		}

	},
}
