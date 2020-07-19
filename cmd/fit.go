package cmd

import (
	"log"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type fitCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newFitCmd() *fitCmd {
	cc := &fitCmd{}

	cmd := &cobra.Command{
		Use:   "fit",
		Short: "Fit the specified maximum width and height and returns the transform",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Fit",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			width, _ := strconv.Atoi(args[0])
			height, _ := strconv.Atoi(args[1])

			var filePath []string

			if b, err := cmd.Flags().GetBool("directory"); b && err == nil {
				filePath = dirwalk(args[2])
			} else {
				filePath = args[2:]
			}

			processing := func(filePath string) error {
				src, err := imaging.Open(filePath)
				if err != nil {
					log.Fatalf("No such file path: %v", filePath)
				}

				// TODO: width and height should be alternative
				dst := imaging.Thumbnail(src, width, height, getFilter(cmd.Flags().Lookup("filter").Value.String()))

				return saveFile(filePath, dst, cmd)
			}

			return saveMultiFile(processing, filePath)

		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
