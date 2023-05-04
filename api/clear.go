package api

import (
	"fmt"
	"net/http"
)

func (a *apiInterface) ClearAndRestart(clearManifest bool) error {
	endpoint := "/clearAndRestart"
	if clearManifest {
		endpoint = endpoint + "?manifest=true"
	}
	resp, err := http.Get(a.url + endpoint)
	if err != nil {
		return fmt.Errorf("error performing clearAndRestart: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http error clearAndRestart: %v", resp.Status)
	}

	return nil
}
