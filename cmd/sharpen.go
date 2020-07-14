package cmd

import (
	"fmt"
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
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0]

			src, err := imaging.Open(filePath)
			if err != nil {
				log.Fatalf("No such file path: %v", filePath)
				return err
			}

			dst := imaging.Sharpen(src, cc.sigma)

			err = imaging.Save(dst, fmt.Sprintf("./result.%s", cmd.Flags().Lookup("extention").Value))
			if err != nil {
				log.Fatalf("Failed to save image: %v", err)
				return err
			}

			return nil
		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	cmd.PersistentFlags().Float64Var(&cc.sigma, "sigma", 0, "")

	return cc
}
