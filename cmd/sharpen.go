package cmd

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type sharpenCmd struct {
	sigma float64

	*baseBuilderCmd
}

func (b *commandsBuilder) newSharpenCmd() *sharpenCmd {
	cc := &sharpenCmd{}

	cmd := &cobra.Command{
		Use:   "sharpen",
		Short: "Generate sharpened version.",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Sharpen",
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

				dst := imaging.Sharpen(src, cc.sigma)

				return saveFile(filePath, dst, cmd)
			}

			return saveMultiFile(processing, filePath)

		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	cmd.PersistentFlags().Float64Var(&cc.sigma, "sigma", 0, "")

	return cc
}
