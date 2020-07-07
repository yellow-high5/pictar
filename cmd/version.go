package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of pictar",
	Long:  `All software has versions. This is pictar's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pictar version pictar1.0.0")
	},
}
