package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
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
			// TODO: rewrite server.boot()
			r := gin.Default()
			r.GET("/", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pictar server startup!!",
				})
			})
			r.Run()

			return nil
		},
	}

	cc.baseBuilderCmd = b.newBaseBuilderCmd(cmd)

	return cc
}
