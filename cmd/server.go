package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yellow-high5/pictar/server"
)

type serverCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newServerCmd() *serverCmd {
	cc := &serverCmd{}

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Startup API server of pictar",
		Long:  `Server provides an endpoint for processing images.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			server.Boot()
			return nil
		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
