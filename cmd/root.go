var rootCmd = &cobra.Command{
  Use:   "pictar",
  Short: "pictar is a image processing tool.",
  Long: `An image processing tool built with
                cobra and imaging.`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}