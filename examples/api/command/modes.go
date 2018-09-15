package command

import (
	"fmt"

	prettyjson "github.com/hokaccha/go-prettyjson"
	"github.com/rigado/edge-connect/edgeconnect/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var mode string

var modesCmd = &cobra.Command{
	Use:   "modes",
	Short: "Get mode information about the radios connected to a cascade gateway",
	Run:   nil,
}

var getModeCmd = &cobra.Command{
	Use:   "get",
	Short: "Get mode for a radio",
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

func modesCommand(root *cobra.Command) {
	modesCmd.AddCommand(setModeCmd)
	modesCmd.AddCommand(getModeCmd)
	getModeCmd.Flags().StringVarP(&radio, "radio", "r", "", "Mode information for specific radio")
	getModeCmd.Flags().StringVarP(&url, "url", "u", "", "URL to send the request")
	getModeCmd.MarkFlagRequired("url")

	setModeCmd.Flags().StringVarP(&radio, "radio", "r", "", "Radio on which to set the mode")
	setModeCmd.Flags().StringVarP(&mode, "mode", "m", "", "Mode to set")
	setModeCmd.Flags().StringVarP(&url, "url", "u", "", "URL to send the request")
	setModeCmd.MarkFlagRequired("mode")
	setModeCmd.MarkFlagRequired("url")
	setModeCmd.MarkFlagRequired("radio")

	rootCmd.AddCommand(modesCmd)
}
