package api

import (
	"github.com/rigado/edge-connect/edgeconnect"
	log "github.com/sirupsen/logrus"
)

type apiInterface struct {
	url string
}

//NewApi ...
func NewApi(url string) edgeconnect.API {

	u := "http://" + url + ":62307/rec/v1"
	log.Debugln("url: %s\n", u)

	return &apiInterface{u}
}
