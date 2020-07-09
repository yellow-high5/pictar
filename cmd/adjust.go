package cmd

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

var (
	saturation float64
	contrast   float64
	brightness float64
	gamma      float64
	sigmoid    float64
	adjustCmd  = &cobra.Command{
		Use:   "adjust",
		Short: "Adjust saturation, contrast, brightness, gamma, sigmoid, LUT",
		Long:  "https://godoc.org/github.com/disintegration/imaging",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0]
			src, err := imaging.Open(filePath)
			if err != nil {
				log.Fatalf("No such file path: %v", filePath)
			}

			// -100 =< saturation =< 100
			if saturation != 0 {
				src = imaging.AdjustSaturation(src, saturation)
			}

			// -100 =< contrast =< 100
			if contrast != 0 {
				src = imaging.AdjustContrast(src, contrast)
			}

			// -100 =< brightness =< 100
			if brightness != 0 {
				src = imaging.AdjustBrightness(src, brightness)
			}

			// gammma > 0
			if gamma != 1.0 {
				src = imaging.AdjustGamma(src, gamma)
			}

			// 0~1, -10 =< sigmoid =< 10
			if sigmoid != 0 {
				src = imaging.AdjustSigmoid(src, 0.5, sigmoid)
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

	adjustCmd.Flags().Float64VarP(&saturation, "saturation", "a", 0, "The depth or intensity of color within an image")
	adjustCmd.Flags().Float64VarP(&contrast, "contrast", "c", 0, "The difference in luminance or colour that makes an object")
	adjustCmd.Flags().Float64VarP(&brightness, "brightness", "b", 0, "The overall lightness or darkness of the image")
	adjustCmd.Flags().Float64VarP(&gamma, "gamma", "g", 1.0, "The value indicating the response characteristics of image gradation")
	adjustCmd.Flags().Float64VarP(&sigmoid, "sigmoid", "s", 0, "The value that determines contrast enhancement")

	rootCmd.AddCommand(adjustCmd)
}
