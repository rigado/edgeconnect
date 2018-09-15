package command

import (
	"fmt"

	"github.com/rigado/edge-connect/edgeconnect/api"
	"github.com/spf13/cobra"
)

const defaultURL = "192.168.240.220"

var scanning bool

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Enable and disable scanning for BLE devices if bt5.0 is available",
	Run:   nil,
}

var scanOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Disables scanning",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			url = defaultURL
		}
		ec := api.NewApi(url)

		r, err := ec.ModeFor("radio0")
		if r.Mode != "bt5.0" {
			fmt.Println("the mode for radio0 is not bt5.0")
			fmt.Println("run: ecapi modes set --r radio0 -m bt5.0 -u <url>")
		}

		err = ec.SetScanning(false)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

var scanOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Disables scanning",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			url = defaultURL
		}
		ec := api.NewApi(url)

		r, err := ec.ModeFor("radio0")
		if r.Mode != "bt5.0" {
			fmt.Println("the mode for radio0 is not bt5.0")
			fmt.Println("run: ecapi modes set --r radio0 -m bt5.0 -u <url>")
		}

		err = ec.SetScanning(true)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func scanCommand(root *cobra.Command) {
	scanOffCmd.Flags().StringVarP(&url, "url", "u", "", "URL to send the request")
	scanOffCmd.MarkFlagRequired("url")
	scanOnCmd.Flags().StringVarP(&url, "url", "u", "", "URL to send the request")
	scanOnCmd.MarkFlagRequired("url")

	scanCmd.AddCommand(scanOffCmd)
	scanCmd.AddCommand(scanOnCmd)

	rootCmd.AddCommand(scanCmd)
}
