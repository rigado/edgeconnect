package command

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Run: nil,
}

//InitCommands initializes the command set
func InitCommands() *cobra.Command {
	versionCommand(rootCmd)
	modesCommand(rootCmd)
	scanCommand(rootCmd)
	firmwareCommand(rootCmd)

	return rootCmd
}
