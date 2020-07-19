package cmd

import (
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
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			var filePath []string

			if b, err := cmd.Flags().GetBool("directory"); b && err == nil {
				filePath = dirwalk(args[0])
			} else {
				filePath = args[0:]
			}

			processing := func(filePath string) error {
				src, err := imaging.Open(filePath)
				if err != nil {
					log.Fatalf("No such file path: %v", filePath)
					return err
				}

				dst := imaging.Transverse(src)

				return saveFile(filePath, dst, cmd)
			}

			return saveMultiFile(processing, filePath)

		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
