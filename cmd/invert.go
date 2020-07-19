package cmd

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type invertCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newInvertCmd() *invertCmd {
	cc := &invertCmd{}

	cmd := &cobra.Command{
		Use:   "invert",
		Short: "Generate invert color version.",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Invert",
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

				dst := imaging.Invert(src)

				return saveFile(filePath, dst, cmd)
			}

			return saveMultiFile(processing, filePath)

		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
