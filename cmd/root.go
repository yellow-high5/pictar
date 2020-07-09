package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Application string
	Debug       bool
}

var (
	configFile      string
	config          Config
	resampleFilter  string
	outputExtension string
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "config.yaml", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVar(&resampleFilter, "filter", "Gaussian", "specifies a resampling filter to be used for image resizing.")
	rootCmd.PersistentFlags().StringVar(&outputExtension, "ext", "png", "specifies the extension of the output file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}

		if err := viper.Unmarshal(&config); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	viper.AutomaticEnv()

}
