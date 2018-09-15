package command

import (
	"fmt"
	"log"

	prettyjson "github.com/hokaccha/go-prettyjson"
	"github.com/rigado/edgeconnect/api"
	"github.com/spf13/cobra"
)

var radio string
var name string
var version string
var file string

var firmwareCmd = &cobra.Command{
	Use:   "upload-fw",
	Short: "Uploads firmware to a radio",
	Run: func(cmd *cobra.Command, args []string) {
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

func firmwareCommand(root *cobra.Command) {
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
	rootCmd.AddCommand(firmwareCmd)
}
