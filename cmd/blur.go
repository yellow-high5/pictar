package cmd

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

var (
	blurCmd = &cobra.Command{
		Use:   "blur",
		Short: "Generate blured version.",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Blur",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0]
			src, err := imaging.Open(filePath)
			if err != nil {
				log.Fatalf("No such file path: %v", filePath)
			}

			// TODO: sigma should be specified
			log.Printf("sigmoid is %v", sigmoid)
			dst := imaging.Blur(src, sigmoid)

			err = imaging.Save(dst, fmt.Sprintf("./result.%s", cmd.Flags().Lookup("ext").Value))
			if err != nil {
				log.Fatalf("Failed to save image: %v", err)
			}

		},
	}
)

func init() {
	rootCmd.AddCommand(blurCmd)

	blurCmd.Flags().Float64VarP(&sigmoid, "sigmoid", "s", 0, "The value that determines contrast enhancement")
}
