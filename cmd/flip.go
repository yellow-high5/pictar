package cmd

import (
	"errors"
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"github.com/yellow-high5/pictar/helper"
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
				filePath = helper.Dirwalk(args[1])
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
				} else if direction == "vertical" {
					src = imaging.FlipV(src)
				} else {
					log.Fatalf("Cannot available such a option")
					return errors.New("Cannot available such a option")
				}

				dst := src

				return helper.SaveFile(filePath, dst, cmd)
			}

			return helper.SaveMultiFile(processing, filePath)

		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
