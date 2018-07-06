package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(name, version, filename, targetIP string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodyWriter.WriteField("mode", name)
	bodyWriter.WriteField("version", version)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("firmware", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	url := "http://" + targetIP + ":62307/rec/v1/modes/firmware"
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
	return nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: fw-upload-test [name] [version] [firmware-file] [target-ip]")
		fmt.Println("    name: The name of the firmware being loaded (not the filename)")
		fmt.Println("    version: The version of the firmware (e.g. 1.0; 1.0.2; 2.3+fixes)")
		fmt.Println("    firmware-file: The path to the firmware Intel HEX file (cannot be binary at this time)")
		fmt.Println("    target-ip: The IP address of the gateway running edge-connect. The default is 127.0.0.1.")
		return
	}

	name := os.Args[1]
	version := os.Args[2]
	firmwareFile := os.Args[3]

	targetIP := "127.0.0.1"
	if len(os.Args) > 4 {
		targetIP = os.Args[4]
	}

	postFile(name, version, firmwareFile, targetIP)
}
