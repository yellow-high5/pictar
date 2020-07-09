package cmd

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rotateCmd)
}

var (
	angle     float64
	rotateCmd = &cobra.Command{
		Use:   "rotate",
		Short: "Rotates an image by the given angle counter-clockwise.",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Rotate",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			angle, _ = strconv.ParseFloat(args[0], 64)
			filePath := args[1]
			src, err := imaging.Open(filePath)
			if err != nil {
				log.Fatalf("No such file path: %v", filePath)
			}

			// TODO: angle and bgcolor should be alternative
			dst := imaging.Rotate(src, angle, color.Transparent)

			err = imaging.Save(dst, fmt.Sprintf("./result.%s", cmd.Flags().Lookup("ext").Value))
			if err != nil {
				log.Fatalf("Failed to save image: %v", err)
			}

		},
	}
)
