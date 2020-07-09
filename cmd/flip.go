package cmd

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

var (
	horizontally bool
	vertically   bool
	flipCmd      = &cobra.Command{
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

			if horizontally {
				src = imaging.FlipH(src)
			}

			if vertically {
				src = imaging.FlipV(src)
			}

			dst := src

			err = imaging.Save(dst, fmt.Sprintf("./result.%s", cmd.Flags().Lookup("ext").Value))
			if err != nil {
				log.Fatalf("Failed to save image: %v", err)
			}

		},
	}
)

func init() {
	rootCmd.AddCommand(flipCmd)

	flipCmd.Flags().BoolVar(&horizontally, "horizontally", false, "")
	flipCmd.Flags().BoolVar(&vertically, "vertically", false, "")
}
