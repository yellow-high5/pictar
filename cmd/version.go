package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type versionCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newVersionCmd() *versionCmd {
	cc := &versionCmd{}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of pictar",
		Long:  `All software has versions. This is pictar's`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("pictar version pictar1.0.0")
			return nil
		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
