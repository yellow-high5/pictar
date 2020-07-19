package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type pictarCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newPictarCmd() *pictarCmd {
	cc := &pictarCmd{}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(&cobra.Command{
		Use:   "pictar",
		Short: "An image processing CLI.",
		Long:  `Pictar is an image processing tool made by Golang. It is designed to be used as both an API and a CLI. 📽`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	})

	cc.cmd.PersistentFlags().StringP("config", "c", "config.yaml", "config file")
	cc.cmd.PersistentFlags().StringP("extention", "e", "png", "specifies the extension of the output file")
	cc.cmd.PersistentFlags().StringP("filter", "f", "Gaussian", "specifies a resampling filter to be used for image resizing")
	cc.cmd.PersistentFlags().StringP("save", "s", ".", "file save destination path after image processing")
	cc.cmd.PersistentFlags().BoolP("directory", "D", false, "Whether to process images as directories")

	return cc
}

func Execute() {
	pictarCmd := newCommandsBuilder().addAll().build()
	cmd := pictarCmd.getCommand()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
