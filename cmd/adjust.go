package cmd

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type adjustCmd struct {
	saturation float64
	contrast   float64
	brightness float64
	gamma      float64
	sigmoid    float64

	*baseBuilderCmd
}

func (b *commandsBuilder) newAdjustCmd() *adjustCmd {
	cc := &adjustCmd{}

	cmd := &cobra.Command{
		Use:   "adjust",
		Short: "Adjust saturation, contrast, brightness, gamma, sigmoid, LUT",
		Long:  "https://godoc.org/github.com/disintegration/imaging",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0:]

			processing := func(filePath string) error {
				src, err := imaging.Open(filePath)
				if err != nil {
					log.Fatalf("No such file path: %v", filePath)
				}

				// -100 =< saturation =< 100
				if cc.saturation != 0 {
					src = imaging.AdjustSaturation(src, cc.saturation)
				}

				// -100 =< contrast =< 100
				if cc.contrast != 0 {
					src = imaging.AdjustContrast(src, cc.contrast)
				}

				// -100 =< brightness =< 100
				if cc.brightness != 0 {
					src = imaging.AdjustBrightness(src, cc.brightness)
				}

				// gammma > 0
				if cc.gamma != 1.0 {
					src = imaging.AdjustGamma(src, cc.gamma)
				}

				// 0~1, -10 =< sigmoid =< 10
				if cc.sigmoid != 0 {
					src = imaging.AdjustSigmoid(src, 0.5, cc.sigmoid)
				}

				dst := src

				return saveFile(filePath, dst, cmd)
			}

			return saveMultiFile(processing, filePath)

		},
	}

	cmd.Flags().Float64Var(&cc.saturation, "saturation", 0, "The depth or intensity of color within an image")
	cmd.Flags().Float64Var(&cc.contrast, "contrast", 0, "The difference in luminance or colour that makes an object")
	cmd.Flags().Float64Var(&cc.brightness, "brightness", 0, "The overall lightness or darkness of the image")
	cmd.Flags().Float64Var(&cc.gamma, "gamma", 1.0, "The value indicating the response characteristics of image gradation")
	cmd.Flags().Float64Var(&cc.sigmoid, "sigmoid", 0, "The value that determines contrast enhancement")

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
