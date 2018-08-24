package main

import (
	"fmt"
	"os"

	"github.com/hokaccha/go-prettyjson"
	"github.com/rigado/edge-connect/edgeconnect/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const defaultURL = "192.168.240.220"

const demoVersion = "0.0.1"

var rootCmd = &cobra.Command{
	Run: nil,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of this demo Edge Connect API application",
	Long:  `Prints the version of this demo Edge Connect API application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Edge Connect API Demo: v", demoVersion)
	},
}
var mode string

var modesCmd = &cobra.Command{
	Use:   "modes",
	Short: "Get available mode information about edge connect",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			url = defaultURL
		}
		ec := api.NewApi(url)

		if radio != "" {
			r, err := ec.ModeFor(radio)
			if err != nil {
				fmt.Println(err)
				return
			}

			j, err := prettyjson.Marshal(r)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(j))
			return
		}

		radios, err := ec.Mode()
		if err != nil {
			fmt.Println(err)
			return
		}

		j, err := prettyjson.Marshal(radios)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(j))
	},
}

var setModeCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the mode for a radio",
	Long:  "Sets the mode for a radio",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			url = defaultURL
		}
		ec := api.NewApi(url)

		if radio == "" {
			radio = "radio0"
		}

		if mode == "" {
			log.Fatalf("empty mode argument")
			return
		}

		err := ec.SetModeFor(radio, mode)
		if err != nil {
			log.Fatalf("failed to set mode of radio %s to %s", radio, mode)
		}

		log.Infof("set %s mode to %s", radio, mode)
	},
}

var radio string
var name string
var version string
var file string
var url string

var firmwareCmd = &cobra.Command{
	Use:   "upload-fw",
	Short: "Uploads firmware to a radio",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			url = defaultURL
		}

		ec := api.NewApi(url)

		if radio != "" {
			_, err := ec.UploadFirmware(radio, name, version, file)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			radio = "radio0"
			_, err := ec.UploadFirmwareOld(name, version, file)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		radios, err := ec.Mode()
		if err != nil {
			log.Fatalln(err)
		}

		r := radios[radio]
		j, err := prettyjson.Marshal(r)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(j))
	},
}

func main() {
	rootCmd.AddCommand(modesCmd)
	rootCmd.AddCommand(firmwareCmd)
	rootCmd.AddCommand(versionCmd)
	modesCmd.AddCommand(setModeCmd)

	modesCmd.Flags().StringVarP(&radio, "radio", "r", "", "Mode information for specific radio")
	modesCmd.Flags().StringVarP(&url, "url", "u", "", "URL to send the request")
	modesCmd.MarkFlagRequired("url")

	setModeCmd.Flags().StringVarP(&radio, "radio", "r", "", "Radio on which to set the mode")
	setModeCmd.Flags().StringVarP(&mode, "mode", "m", "", "Mode to set")
	setModeCmd.Flags().StringVarP(&url, "url", "u", "", "URL to send the request")
	setModeCmd.MarkFlagRequired("mode")
	setModeCmd.MarkFlagRequired("url")
	setModeCmd.MarkFlagRequired("radio")

	firmwareCmd.Flags().StringVarP(&radio, "radio", "r", "", "Radio to which to upload firmware")
	firmwareCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the firmware")
	firmwareCmd.Flags().StringVarP(&version, "version", "v", "", "Version of the firmware")
	firmwareCmd.Flags().StringVarP(&file, "file", "f", "", "Firmware file in Intel HEX format")
	firmwareCmd.Flags().StringVarP(&url, "url", "u", "", "URL to send the request")
	firmwareCmd.MarkFlagRequired("name")
	firmwareCmd.MarkFlagRequired("file")
	firmwareCmd.MarkFlagRequired("version")
	firmwareCmd.MarkFlagRequired("url")
	firmwareCmd.MarkFlagRequired("radio")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
