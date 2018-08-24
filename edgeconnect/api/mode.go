package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/rigado/edge-connect/edgeconnect"

	log "github.com/sirupsen/logrus"
)

//todo: DRY this up
//...
func (a *apiInterface) Mode() (map[string]edgeconnect.Radio, error) {

	resp, err := http.Get(a.url + "/modes")
	if err != nil {
		return nil, fmt.Errorf("error performing getting modes: %s", err.Error())
	}
	defer resp.Body.Close()

	var radios map[string]edgeconnect.Radio
	err = HandleResponse(resp, &radios)

	if err != nil {
		return nil, err
	}

	return radios, nil

}

//...
func (a *apiInterface) SetMode(mode string) error {

	log.Debugln("Swtiching to: ", mode, ". . .")

	m := edgeconnect.Modes{Mode: mode}

	jb, err := json.Marshal(m)
	bytes := bytes.NewBuffer(jb)

	resp, err := http.Post(a.url+"/modes", "application/json", bytes)
	if err != nil {
		return fmt.Errorf("error setting mode for radio0: %s", err.Error())
	}
	defer resp.Body.Close()

	var s edgeconnect.Status
	err = HandleResponse(resp, &s)

	if err != nil {
		return err
	}

	if s.Status != "Success" {
		return fmt.Errorf("failed to set mode: %s", s.Status)
	}

	return nil
}

func (a *apiInterface) SetModeFor(radio string, mode string) error {

	log.Debugf("Swtiching %s to mode %s\n", radio, mode)

	m := edgeconnect.Modes{Mode: mode}

	jb, err := json.Marshal(m)
	bytes := bytes.NewBuffer(jb)

	resp, err := http.Post(a.url+"/modes/"+radio, "application/json", bytes)
	if err != nil {
		return fmt.Errorf("error setting mode for radio0: %s", err.Error())
	}
	defer resp.Body.Close()

	var s edgeconnect.Status
	err = HandleResponse(resp, &s)

	if err != nil {
		return err
	}

	if s.Status != "Success" {
		return fmt.Errorf("failed to set %s to mode %s", s.Status)
	}

	return nil
}

func (a *apiInterface) ModeFor(radio string) (edgeconnect.Radio, error) {
	res, err := http.Get(a.url + "/modes/" + radio)

	if err != nil {
		return edgeconnect.Radio{}, fmt.Errorf("error getting mode for %s: %s", radio, err.Error())
	}

	var r edgeconnect.Radio
	err = HandleResponse(res, &r)

	if err != nil {
		return edgeconnect.Radio{}, err
	}

	return r, nil
}

func (a *apiInterface) UploadFirmware(radio, name, version, file string) (string, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodyWriter.WriteField("mode", name)
	bodyWriter.WriteField("version", version)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("firmware", file)
	if err != nil {
		fmt.Println("error writing to buffer")
		return "", err
	}

	// open file handle
	fh, err := os.Open(file)
	if err != nil {
		fmt.Println("error opening file")
		return "", err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return "", err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	url := a.url + "/modes/firmware/" + radio
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %d - %s", resp.StatusCode, string(respBody))
	}

	fmt.Println(resp.Status)
	fmt.Println(string(respBody))

	return "", nil
}

func (a *apiInterface) UploadFirmwareOld(name, version, file string) (string, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodyWriter.WriteField("mode", name)
	bodyWriter.WriteField("version", version)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("firmware", file)
	if err != nil {
		fmt.Println("error writing to buffer")
		return "", err
	}

	// open file handle
	fh, err := os.Open(file)
	if err != nil {
		fmt.Println("error opening file")
		return "", err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return "", err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	url := a.url + "/modes/firmware"
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %d - %s", resp.StatusCode, string(respBody))
	}

	fmt.Println(resp.Status)
	fmt.Println(string(respBody))

	return "", nil
}
