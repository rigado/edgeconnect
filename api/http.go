package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//HandleResponse ...
func HandleResponse(resp *http.Response, v interface{}) error {
	res, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return fmt.Errorf("error reading response body: ", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received not ok status: %d - %s", resp.StatusCode, string(res))
	}

	err = json.Unmarshal(res, &v)

	if err != nil {
		return err
	}

	return nil

}

//HandleRespString ...
func HandleRespString(resp *http.Response, s interface{}) (string, error) {

	res, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return "", fmt.Errorf("error reading response body; ", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received not ok status: %d - %s", resp.StatusCode, string(res))
	}

	err = json.Unmarshal(res, &s)

	if err != nil {
		return "", err
	}

	return string(res), nil

}
