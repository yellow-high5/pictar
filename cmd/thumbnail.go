package cmd

import (
	"log"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"github.com/yellow-high5/pictar/helper"
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
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			width, _ := strconv.Atoi(args[0])
			height, _ := strconv.Atoi(args[1])

			var filePath []string

			if b, err := cmd.Flags().GetBool("directory"); b && err == nil {
				filePath = helper.Dirwalk(args[2])
			} else {
				filePath = args[2:]
			}

			processing := func(filePath string) error {
				src, err := imaging.Open(filePath)
				if err != nil {
					log.Fatalf("No such file path: %v", filePath)
					return err
				}

				dst := imaging.Thumbnail(src, width, height, helper.GetFilter(cmd.Flags().Lookup("filter").Value.String()))

				return helper.SaveFile(filePath, dst, cmd)
			}

			return helper.SaveMultiFile(processing, filePath)
		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
