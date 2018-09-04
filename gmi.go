package goGMI

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// GMI is your client
type GMI struct {
	APIVersion string
}

func (gmi *GMI) do(path string, methode string, in, out interface{}) error {
	// Create client
	client := &http.Client{}
	// Turn the struct into JSON bytes
	b, _ := json.Marshal(&in)
	// Post JSON request to FreshDesk
	req, _ := http.NewRequest(methode, fmt.Sprintf("https://api.getmyinvoices.com/accounts/%s/%s", gmi.APIVersion, path), bytes.NewReader(b))
	req.Header.Add("Content-type", "application/json")
	res, e := client.Do(req)
	if e != nil {
		return e
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	newStr := buf.String()

	fmt.Printf(newStr)

	// Check the status
	if res.StatusCode != 200 {
		return errors.New("Freshdesk server didn't like the request")
	}
	// Grab the JSON response
	if e = json.NewDecoder(res.Body).Decode(out); e != nil {
		return e
	}
	return nil
}

// ListSuppliers give a list of all suppliers
func (gmi *GMI) ListSuppliers() error {
	var response interface{}
	err := gmi.do("listSuppliers", http.MethodGet, map[string]string{"api_key": "x0ql-l1zb-ha18-8vhz-1jzg-0oho-oh7s"}, &response)
	log.Println(response)
	return err
}
