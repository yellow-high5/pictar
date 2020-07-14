package cmd

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type transverseCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newTransverseCmd() *transverseCmd {
	cc := &transverseCmd{}

	cmd := &cobra.Command{
		Use:   "transverse",
		Short: "Flips the image vertically and rotates 90 degrees counter-clockwise.",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Transverse",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0]

			src, err := imaging.Open(filePath)
			if err != nil {
				log.Fatalf("No such file path: %v", filePath)
				return err
			}

			dst := imaging.Transverse(src)

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
