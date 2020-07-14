package cmd

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

type blurCmd struct {
	sigma float64

	*baseBuilderCmd
}

func (b *commandsBuilder) newBlurCmd() *blurCmd {
	cc := &blurCmd{}

	cmd := &cobra.Command{
		Use:   "blur",
		Short: "Generate blured version.",
		Long:  "https://godoc.org/github.com/disintegration/imaging#Blur",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0]

			src, err := imaging.Open(filePath)
			if err != nil {
				log.Fatalf("No such file path: %v", filePath)
				return err
			}

			// TODO: sigma should be specified
			log.Printf("sigma is %v", cc.sigma)
			dst := imaging.Blur(src, cc.sigma)

			err = imaging.Save(dst, fmt.Sprintf("./result.%s", cmd.Flags().Lookup("extention").Value))
			if err != nil {
				log.Fatalf("Failed to save image: %v", err)
				return err
			}

			return nil
		},
	}

	cmd.Flags().Float64VarP(&cc.sigma, "sigma", "s", 0, "The value that determines contrast enhancement")

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
