package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WebhookRequest struct {
	TimeStamp string  `json:"TimeStamp"`
	Values    []Value `json:"Values"`
}

type Value struct {
	FieldName       string      `json:"FieldName"`
	IsNull          string      `json:"IsNull"`
	Item            interface{} `json:"Item"`
	ItemElementName string      `json:"ItemElementName"`
}

func main() {
	handler := http.HandlerFunc(webhookHandler)
	address := ""
	port := 5000
	serverAddressWithPort := fmt.Sprintf("%s:%d", address, port)

	http.Handle("/webhook", handler)
	http.ListenAndServe(serverAddressWithPort, nil)
	log.Fatal(fmt.Sprintf("The server is running on %s", serverAddressWithPort))
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Wasn't able to read the body. Err: %s", err)
	}

	// Deserialize body content
	var jsonBody WebhookRequest
	err = json.Unmarshal([]byte(body), &jsonBody)
	if err != nil {
		log.Fatalf("Can't handle the sent JSON with JSON unmarshal. Err: %s", err)
	}

	// Serialize data for nice console log output
	requestBody, err := json.MarshalIndent(jsonBody, "", "\t")
	if err != nil {
		log.Fatalf("Can't handle the sent JSON with JSON marshal. Err: %s", err)
	}

	log.Println("Request body:\n", string(requestBody))
}
