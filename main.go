package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
	"strings"
)

type requestData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	

	cert, err := tls.LoadX509KeyPair("./certificate.crt", "./private.key")
	if err != nil {
		panic(err)
	}

	customTLSConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: customTLSConfig,
		},
	}

	hittingURL := "https://apideveloper.rblbank.com/test/sb/rbl/api/v1/sendernotification/rpay"
	data := requestData{
		Name:  "razorpay",
		Email: "razorpay@gmail.com",
	}
	payloadData, error := json.Marshal(data)
	if error != nil {
		panic(error)
	}
	payload := strings.NewReader(string(payloadData))

	req, err := http.NewRequest("POST", hittingURL, payload)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	log.Println("SUCCESS", "handshake completed")

	fmt.Println(string(responseBody))
}
