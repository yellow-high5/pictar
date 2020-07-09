package cmd

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

var (
	sigma      float64
	sharpenCmd = &cobra.Command{
		Use:   "sharpen",
		Short: "Generate sharpened version.",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Sharpen",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0]
			src, err := imaging.Open(filePath)
			if err != nil {
				log.Fatalf("No such file path: %v", filePath)
			}

			dst := imaging.Sharpen(src, sigma)

			err = imaging.Save(dst, "./result.png")
			if err != nil {
				log.Fatalf("Failed to save image: %v", err)
			}

		},
	}
)

func init() {
	rootCmd.AddCommand(sharpenCmd)

	sharpenCmd.Flags().Float64Var(&sigma, "sigma", 0, "")
}
