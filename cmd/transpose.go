package cmd

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type transposeCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newTransposeCmd() *transposeCmd {
	cc := &transposeCmd{}

	cmd := &cobra.Command{
		Use:   "transpose",
		Short: "Flips the image horizontally and rotates 90 degrees counter-clockwise.",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Transpose",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0]

			src, err := imaging.Open(filePath)
			if err != nil {
				log.Fatalf("No such file path: %v", filePath)
				return err
			}

			dst := imaging.Transpose(src)

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
