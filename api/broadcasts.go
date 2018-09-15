package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	sse "astuart.co/go-sse"
	"github.com/rigado/edgeconnect"

	log "github.com/sirupsen/logrus"
)

//...
func (a *apiInterface) Broadcasts() (map[string]edgeconnect.Broadcast, error) {

	res, err := http.Get(a.url + "/broadcasts")
	if err != nil {
		return nil, fmt.Errorf("error getting broadcasts: %s", err.Error())
	}

	var bcs map[string]edgeconnect.Broadcast
	err = HandleResponse(res, &bcs)

	if err != nil {
		return nil, err
	}

	return bcs, nil

}

func (a *apiInterface) BroadcastForDevice(addr string) (edgeconnect.Broadcast, error) {

	res, err := http.Get(a.url + "/broadcasts/" + addr)

	if err != nil {
		return edgeconnect.Broadcast{}, fmt.Errorf("error getting broadcast for %s: %s", addr, err.Error())
	}

	var bc edgeconnect.Broadcast
	err = HandleResponse(res, &bc)

	if err != nil {
		return edgeconnect.Broadcast{}, err
	}

	return bc, nil

}

//...
func (a *apiInterface) PrintBroadcast(b edgeconnect.Broadcast) {

	log.Debugf("---------------------------------------------------------------------------------------------------------------------")
	log.Debugln("Company:", b.Company)
	log.Debugln("Address:", b.Address, "  Type:", b.Type, "  Rssi:", b.Rssi, "  Last Seen:", b.Seen)

	if b.Company == "Apple" {
		log.Debugln("UUID:", b.Ibeacon.UUID, "| Major:", b.Ibeacon.Major, "| Minor:", b.Ibeacon.Minor, "| TxPower: ", b.Ibeacon.TxPower)
	}

	log.Debugln("UID:", b.Eddystone.UID.UID, " TxPower:", b.Eddystone.UID.TxPower)
	log.Debugln("url:", b.Eddystone.URL.URL, "TxPower:", b.Eddystone.URL.TxPower)
	log.Debugln("Version:", b.Eddystone.TLM.Version, " Battery:", b.Eddystone.TLM.Battery, " Temperature:", b.Eddystone.TLM.Temperature, " Advertising Count:", b.Eddystone.TLM.AdvertisingCount, "Sec: ", b.Eddystone.TLM.SecCount)
}

func (a *apiInterface) BroadcastEvents(stop chan struct{}) error {
	url := a.url + "/broadcastEvents"
	event := make(chan *sse.Event)

	go sse.Notify(url, event)

	for {
		select {
		case e := <-event:

			res, err := ioutil.ReadAll(e.Data)

			if err != nil {
				return fmt.Errorf("error reading response body: %s", err.Error())
			}
			var v edgeconnect.Broadcast
			err = json.Unmarshal(res, &v)

			if err != nil {
				return err
			}

			a.PrintBroadcast(v)
		case <-stop:
			return nil
		}
	}
}

func (a *apiInterface) BroadcastEventsForDevice(addr string, stop chan struct{}) error {
	url := (a.url + "/broadcasts/" + addr + "/events")
	event := make(chan *sse.Event)

	go sse.Notify(url, event)

	for {
		select {
		case e := <-event:
			res, err := ioutil.ReadAll(e.Data)

			if err != nil {
				log.Fatalln("Error reading http Response body; ", err.Error())
			}
			var v edgeconnect.Broadcast
			err = json.Unmarshal(res, &v)

			if err != nil {
				log.Fatalln("Error unmarshalling JSON data; ", err.Error())
			}

			a.PrintBroadcast(v)
		case <-stop:
			return nil
		}
	}

}
