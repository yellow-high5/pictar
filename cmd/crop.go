package cmd

import (
	"log"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type cropCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newCropCmd() *cropCmd {
	cc := &cropCmd{}

	cmd := &cobra.Command{
		Use:   "crop",
		Short: "Cuts out a rectangular region.",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Crop",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			width, _ := strconv.Atoi(args[0])
			height, _ := strconv.Atoi(args[1])
			filePath := args[2]

			src, err := imaging.Open(filePath)
			if err != nil {
				log.Fatalf("No such file path: %v", filePath)
				return err
			}

			dst := imaging.CropCenter(src, width, height)

			return saveFile(filePath, dst, cmd)

		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
