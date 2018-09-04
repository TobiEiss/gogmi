package gogmi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GMI is your client
type GMI struct {
	APIVersion string
	APIKey     string
}

func (gmi *GMI) do(path string, methode string, in map[string]interface{}, out interface{}) error {
	// Create client
	client := &http.Client{}

	// add key
	if in == nil {
		in = map[string]interface{}{}
	}
	in["api_key"] = gmi.APIKey

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
func (gmi *GMI) ListSuppliers() (suppliers Suppliers, err error) {
	err = gmi.do("listSuppliers", http.MethodGet, nil, &suppliers)
	return
}

// GetSupplier returns a specific supplier
func (gmi *GMI) GetSupplier(primUID int) (supplier Supplier, err error) {
	err = gmi.do("getSupplier", http.MethodGet, map[string]interface{}{"supplier_id": primUID}, &supplier)
	return
}
