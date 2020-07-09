package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

var cropCmd = &cobra.Command{
	Use:   "crop",
	Short: "Cuts out a rectangular region.",
	Long:  "https://godoc.org/github.com/disintegration/imaging#Crop",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		width, _ := strconv.Atoi(args[0])
		height, _ := strconv.Atoi(args[1])
		filePath := args[2]
		src, err := imaging.Open(filePath)
		if err != nil {
			log.Fatalf("No such file path: %v", filePath)
		}

		// TODO: crop size and anchor should be alternative
		dst := imaging.CropCenter(src, width, height)

		err = imaging.Save(dst, fmt.Sprintf("./result.%s", cmd.Flags().Lookup("ext").Value))
		if err != nil {
			log.Fatalf("Failed to save image: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(cropCmd)
}
