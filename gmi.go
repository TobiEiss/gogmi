package gogmi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
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
		return errors.New("server didn't like the request")
	}
	// Grab the JSON response
	if e = json.NewDecoder(res.Body).Decode(out); e != nil {
		return e
	}
	return nil
}

// ListSuppliers give a list of all suppliers
func (gmi *GMI) ListSuppliers() (suppliers Suppliers, err error) {
	err = gmi.do("listSuppliers", http.MethodPost, nil, &suppliers)
	return
}

// GetSupplier returns a specific supplier
func (gmi *GMI) GetSupplier(primUID int) (supplier Supplier, err error) {
	err = gmi.do("getSupplier", http.MethodPost, map[string]interface{}{"supplier_id": primUID}, &supplier)
	return
}

// ListInvoices returns all invoices
func (gmi *GMI) ListInvoices() (invoices []Invoice, err error) {
	var rack RecordsRack
	err = gmi.do("listInvoices", http.MethodPost, nil, &rack)
	invoices = rack.Invoices
	return
}

// ListInvoicesFilterByDate returns all invoices with a filter by date
func (gmi *GMI) ListInvoicesFilterByDate(startDate time.Time) (invoices []Invoice, err error) {
	var rack RecordsRack
	in := map[string]interface{}{}
	in["start_date_filter"] = startDate.Format("2006-01-02")
	err = gmi.do("listInvoices", http.MethodPost, nil, &rack)
	invoices = rack.Invoices
	return
}

// GetInvoice returns specific invoice
func (gmi *GMI) GetInvoice(primUID PrimUID) (rack interface{}, err error) {
	err = gmi.do("getInvoice", http.MethodPost, map[string]interface{}{"invoice_prim_uid": primUID}, &rack)
	return
}

// GetCountries returns a slice of countries
func (gmi *GMI) GetCountries() (countries Countries, err error) {
	err = gmi.do("getCountries", http.MethodPost, nil, &countries)
	return
}
