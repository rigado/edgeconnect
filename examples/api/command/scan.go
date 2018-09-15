package command

import (
	"fmt"
	"log"

	"github.com/rigado/edgeconnect/api"
	"github.com/spf13/cobra"
)

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

		fmt.Println("scanning disabled")
	},
}

var scanOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Enables scanning",
	Run: func(cmd *cobra.Command, args []string) {
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

		fmt.Println("scanning enabled")
	},
}

var scanStateCmd = &cobra.Command{
	Use:   "state",
	Short: "Retrieves the current scanning state",
	Run: func(cmd *cobra.Command, args []string) {
		ec := api.NewApi(url)

		r, err := ec.ModeFor("radio0")
		if r.Mode != "bt5.0" {
			fmt.Println("the mode for radio0 is not bt5.0")
			fmt.Println("run: ecapi modes set --r radio0 -m bt5.0 -u <url>")
		}

		state, err := ec.IsScanning()
		if err != nil {
			fmt.Println(err)
			return
		}

		if state {
			log.Println("scanning on")
		} else {
			log.Println("scanning off")
		}
	},
}

func scanCommand(root *cobra.Command) {
	scanOffCmd.Flags().StringVarP(&url, "url", "u", "", "URL to send the request")
	scanOffCmd.MarkFlagRequired("url")
	scanOnCmd.Flags().StringVarP(&url, "url", "u", "", "URL to send the request")
	scanOnCmd.MarkFlagRequired("url")
	scanStateCmd.Flags().StringVarP(&url, "url", "u", "", "URL to send the request")
	scanStateCmd.MarkFlagRequired("url")

	scanCmd.AddCommand(scanStateCmd)
	scanCmd.AddCommand(scanOffCmd)
	scanCmd.AddCommand(scanOnCmd)

	rootCmd.AddCommand(scanCmd)
}
