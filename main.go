package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	cert, error := tls.LoadX509KeyPair("./certificate.crt", "./private.key")
	if error != nil {
		panic(error)
	}

	customtlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: customtlsConfig,
		},
	}

	hittingURL := "https://apideveloper.rblbank.com/test/sb/rbl/api/v1/sendernotification/rpay"

	response, error := httpClient.Get(hittingURL)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	responseBody, error := ioutil.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}
	fmt.Println(string(responseBody))
}
