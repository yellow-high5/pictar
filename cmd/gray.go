package cmd

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type grayCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newGrayCmd() *grayCmd {
	cc := &grayCmd{}

	cmd := &cobra.Command{
		Use:   "gray",
		Short: "Generate gray scale version.",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Grayscale",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0:]

			processing := func(filePath string) error {
				src, err := imaging.Open(filePath)
				if err != nil {
					log.Fatalf("No such file path: %v", filePath)
					return err
				}

				dst := imaging.Grayscale(src)

				return saveFile(filePath, dst, cmd)
			}

			return saveMultiFile(processing, filePath)

		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
