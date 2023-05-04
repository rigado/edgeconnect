package main

import (
	"flag"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/rigado/edgeconnect/api"
)

var (
	ip        = flag.String("ip", "", "ip address of the gateway")
	wpVersion = flag.String("wp-version", "", "wirepas version to check for")
)

func main() {
	flag.Parse()

	if *ip == "" {
		log.Fatal("ip address is required, use -ip")
	}

	if *wpVersion == "" {
		log.Fatal("wirepas version is required, use -wp-version")
	}

	ec := api.NewApi(*ip)

	attempts := 10
	needsClearAndRestart := false
	for attempts > 0 {
		modes, err := ec.Mode()
		if err != nil {
			if strings.Contains(err.Error(), "received not ok status") {
				log.Warnf("edge connect is not available; trying again in 30 seconds")
				attempts--
				time.Sleep(30 * time.Second)
				continue
			}
		}

		m, ok := modes["radio0"]
		if !ok {
			log.Fatal("radio not found modes output")
		}

		if m.Custom {
			log.Infof("radio has custom firmware programmed. cleaingup edge connect and restarting it")
			needsClearAndRestart = true
		} else {
			log.Infof("radio does not have custom firmware programmed: %s", m.Mode)
			break
		}

		if needsClearAndRestart {
			err := ec.ClearAndRestart(true) //also clears the manifest to force re-flash of radio firmware
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if needsClearAndRestart {
		//wait for a while for edge connect to restart
		time.Sleep(2 * time.Minute) // this may need to be tweaked or handled in another way
	}

	//get the modes again and verify that the radio is in the correct mode
	for attempts > 0 {
		modes, err := ec.Mode()
		if err != nil {
			if strings.Contains(err.Error(), "received not ok status") {
				log.Warnf("edge connect is not available; trying again in 30 seconds")
				attempts--
				time.Sleep(30 * time.Second)
				continue
			}
		}

		m, ok := modes["radio0"]
		if !ok {
			log.Fatal("radio not found in modes output")
		}

		if m.Mode != *wpVersion {
			log.Fatalf("the radio is not in the correct mode. expected: %s, actual: %s", *wpVersion, m.Mode)
		}
	}
}
