package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

const demoVersion = "0.0.6"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of this demo Edge Connect API application",
	Long:  `Prints the version of this demo Edge Connect API application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Edge Connect API Demo: v", demoVersion)
	},
}

func versionCommand(root *cobra.Command) {
	rootCmd.AddCommand(versionCmd)
}
