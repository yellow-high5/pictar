package cmd

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type flipCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newFlipCmd() *flipCmd {
	cc := &flipCmd{}

	cmd := &cobra.Command{
		Use:   "flip",
		Short: "Flips the image.",
		Long:  "https://godoc.org/github.com/disintegration/imaging#FlipH or https://godoc.org/github.com/disintegration/imaging#FlipV",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			direction := args[0]

			var filePath []string

			if b, err := cmd.Flags().GetBool("directory"); b && err == nil {
				filePath = dirwalk(args[1])
			} else {
				filePath = args[1:]
			}

			processing := func(filePath string) error {
				src, err := imaging.Open(filePath)
				if err != nil {
					log.Fatalf("No such file path: %v", filePath)
					return err
				}

				if direction == "horizon" {
					src = imaging.FlipH(src)
				}

				if direction == "vertical" {
					src = imaging.FlipV(src)
				}

				dst := src

				return saveFile(filePath, dst, cmd)
			}

			return saveMultiFile(processing, filePath)

		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
