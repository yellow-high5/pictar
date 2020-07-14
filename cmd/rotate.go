package cmd

import (
	"fmt"
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
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			angle, _ := strconv.ParseFloat(args[0], 64)
			filePath := args[1]

			src, err := imaging.Open(filePath)
			if err != nil {
				log.Fatalf("No such file path: %v", filePath)
				return err
			}

			// TODO: angle and bgcolor should be alternative
			dst := imaging.Rotate(src, angle, color.Transparent)

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
