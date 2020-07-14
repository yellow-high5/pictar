package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type thumbnailCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newThumbnailCmd() *thumbnailCmd {
	cc := &thumbnailCmd{}

	cmd := &cobra.Command{
		Use:   "thumbnail",
		Short: "Scales the image up or down",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Thumbnail",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			width, _ := strconv.Atoi(args[0])
			height, _ := strconv.Atoi(args[1])
			filePath := args[2]

			src, err := imaging.Open(filePath)
			if err != nil {
				log.Fatalf("No such file path: %v", filePath)
				return err
			}

			dst := imaging.Thumbnail(src, width, height, imaging.Gaussian)

			err = imaging.Save(dst, fmt.Sprintf("./result.%s", cmd.Flags().Lookup("extention").Value))
			if err != nil {
				log.Fatalf("Failed to save image: %v", err)
				return err
			}

			return nil

		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
