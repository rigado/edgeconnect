package command

import (
	"fmt"
	"os"

	"github.com/rigado/edgeconnect/api"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset a radio using the hardware reset pin",
	Run: func(cmd *cobra.Command, args []string) {
		ec := api.NewApi(url)

		if radio == "" {
			fmt.Println("radio cannot be empty")
			os.Exit(1)
		}

		err := ec.ResetRadio(radio)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%s was reset\n", radio)
	},
}

func resetCommand(root *cobra.Command) {
	resetCmd.Flags().StringVarP(&radio, "radio", "r", "", "Radio to which to upload firmware")
	resetCmd.MarkFlagRequired("radio")
	resetCmd.Flags().StringVarP(&url, "url", "u", "", "URL to send the request")
	resetCmd.MarkFlagRequired("url")
	rootCmd.AddCommand(resetCmd)
}
