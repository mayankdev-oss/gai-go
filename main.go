package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cob = &cobra.Command{
	Use:   "gai",
	Short: "This command helps in initializing gai tool to explain the codebase",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gai! The AI codebase explainer.")
	},
}

func main() {
	Cob.Execute()

}
