package cmd

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(transverseCmd)
}

var transverseCmd = &cobra.Command{
	Use:   "transverse",
	Short: "Flips the image vertically and rotates 90 degrees counter-clockwise.",
	Long:  "https://godoc.org/github.com/disintegration/imaging#Transverse",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		src, err := imaging.Open(filePath)
		if err != nil {
			log.Fatalf("No such file path: %v", filePath)
		}

		dst := imaging.Transverse(src)

		err = imaging.Save(dst, fmt.Sprintf("./result.%s", cmd.Flags().Lookup("ext").Value))
		if err != nil {
			log.Fatalf("Failed to save image: %v", err)
		}

	},
}
