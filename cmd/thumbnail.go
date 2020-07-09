package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

var (
	thumbnailCmd = &cobra.Command{
		Use:   "thumbnail",
		Short: "Scales the image up or down",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Thumbnail",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			width, _ := strconv.Atoi(args[0])
			height, _ := strconv.Atoi(args[1])
			filePath := args[2]
			src, err := imaging.Open(filePath)
			if err != nil {
				log.Fatalf("No such file path: %v", filePath)
			}

			dst := imaging.Thumbnail(src, width, height, imaging.Gaussian)

			err = imaging.Save(dst, fmt.Sprintf("./result.%s", cmd.Flags().Lookup("ext").Value))
			if err != nil {
				log.Fatalf("Failed to save image: %v", err)
			}

		},
	}
)

func init() {
	rootCmd.AddCommand(thumbnailCmd)
}
