package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type resizeCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newResizeCmd() *resizeCmd {
	cc := &resizeCmd{}

	cmd := &cobra.Command{
		Use:   "resize",
		Short: "Resizes the image to the specified width and height",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Resize",
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

			// TODO: width and height should be alternative
			dst := imaging.Resize(src, width, height, imaging.Gaussian)

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
