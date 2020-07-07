package main

import (
	"github.com/yellow-high5/pictar/cmd"
)

func main() {
	cmd.Execute()
	// resp := cmd.Execute()

	// if resp.Err != nil {
	// 	if resp.IsUserError() {
	// 		resp.Cmd.Println("")
	// 		resp.Cmd.Println(resp.Cmd.UsageString())
	// 	}
	// 	os.Exit(-1)
	// }

}
