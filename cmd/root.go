package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile         string
	resampleFilter  string
	outputExtension string
	saveImagePath   string
	rootCmd         = &cobra.Command{
		Use:   "pictar",
		Short: "An image processing CLI.",
		Long:  `Pictar is an image processing tool made by Golang. It is designed to be used as both an API and a CLI. 📽`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			fmt.Println("hello, pictar")
		},
	}
)

func init() {
	// cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	viper.SetDefault("license", "apache")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
