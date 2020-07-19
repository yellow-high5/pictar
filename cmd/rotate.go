package cmd

import (
	"image/color"
	"log"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type rotateCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newRotateCmd() *rotateCmd {
	cc := &rotateCmd{}

	cmd := &cobra.Command{
		Use:   "rotate",
		Short: "Rotates an image by the given angle counter-clockwise.",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Rotate",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			angle, _ := strconv.ParseFloat(args[0], 64)

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

				// TODO: angle and bgcolor should be alternative
				dst := imaging.Rotate(src, angle, color.Transparent)

				return saveFile(filePath, dst, cmd)
			}

			return saveMultiFile(processing, filePath)

		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
