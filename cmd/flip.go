package cmd

import (
	"fmt"
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
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			direction := args[0]
			filePath := args[1]

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
