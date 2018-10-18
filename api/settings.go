package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rigado/edgeconnect"
)

//...
func (a *apiInterface) IsScanning() (bool, error) {

	resp, err := http.Get(a.url + "/settings")
	if err != nil {
		return false, fmt.Errorf("error getting scan state: %s", err.Error())
	}
	defer resp.Body.Close()

	var data edgeconnect.Settings
	err = HandleResponse(resp, &data)

	if err != nil {
		return false, err
	}

	return data.Scanning, nil

}

//...
func (a *apiInterface) SetScanning(scan bool) error {

	s := edgeconnect.Settings{Scanning: scan}

	b, err := json.Marshal(s)
	bytes := bytes.NewBuffer(b)

	resp, err := http.Post(a.url+"/settings", "application/json", bytes)
	if err != nil {
		return fmt.Errorf("error setting scan state: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("error setting scan state: %d - %s", resp.StatusCode, string(msg))
	}

	return nil
}

func (a *apiInterface) ResetRadio(radio string) error {
	payload := struct{}{}

	b, err := json.Marshal(payload)
	data := bytes.NewBuffer(b)

	ep := fmt.Sprintf("%s/settings/%s/reset", a.url, radio)
	resp, err := http.Post(ep, "application/json", data)
	if err != nil {
		return fmt.Errorf("error reseting %s: %s", radio, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("radio reset server error: %d", resp.StatusCode)
	}

	return nil
}
