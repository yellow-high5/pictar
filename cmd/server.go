package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Startup API server of pictar",
	Long:  `Server provides an endpoint for processing images.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pictar server startup!!",
			})
		})
		r.Run()
	},
}
