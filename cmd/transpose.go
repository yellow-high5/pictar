package cmd

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"github.com/yellow-high5/pictar/helper"
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
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			var filePath []string

			if b, err := cmd.Flags().GetBool("directory"); b && err == nil {
				filePath = helper.Dirwalk(args[0])
			} else {
				filePath = args[0:]
			}

			processing := func(filePath string) error {
				src, err := imaging.Open(filePath)
				if err != nil {
					log.Fatalf("No such file path: %v", filePath)
					return err
				}

				dst := imaging.Transpose(src)

				return helper.SaveFile(filePath, dst, cmd)
			}

			return helper.SaveMultiFile(processing, filePath)

		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
