package gogmi

import (
	"encoding/json"
	"strconv"
)

// PrimUID is just a workaround, cause' there is a bug in the API.
// /getSuppliers returns an string instead an int.
type PrimUID int

func (prim *PrimUID) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(prim))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*prim = PrimUID(i)
	return nil
}

// Supplier is the model for GMI-supplier
type Supplier struct {
	PrimUID      PrimUID `json:"prim_uid"`
	Name         string  `json:"name"`
	SupplierType string  `json:"supplier_type"`
	Note         string  `json:"note"`
	Tags         string  `json:"tags"`
}

// Suppliers is a slice of suppliers
type Suppliers []Supplier

// Invoice represent a record
type Invoice struct {
	PrimUID        string `json:"prim_uid,omitempty"`
	SupplierUID    string `json:"supplier_uid,omitempty"`
	InvoiceNumber  string `json:"invoice_number,omitempty"`
	InvoiceDate    string `json:"invoice_date,omitempty"`
	InvoiceDueDate string `json:"invoice_due_date,omitempty"`
	NetAmount      string `json:"net_amount,omitempty"`
	Vat            string `json:"vat,omitempty"`
	GrossAmount    string `json:"gross_amount,omitempty"`
	Currency       string `json:"currency,omitempty"`
	IsArchived     string `json:"is_archived,omitempty"`
	IsOcrCompleted int    `json:"is_ocr_completed,omitempty"`
	Tags           string `json:"tags,omitempty"`
	Note           string `json:"note,omitempty"`
	Source         string `json:"source,omitempty"`
	Filename       string `json:"filename,omitempty"`
	FileSize       string `json:"file_size,omitempty"`
	PaymentStatus  string `json:"payment_status,omitempty"`
	PaymentMethod  string `json:"payment_method,omitempty"`
	PaymentDetails struct {
		Iban              string `json:"iban,omitempty"`
		Bic               string `json:"bic,omitempty"`
		AccountHolderName string `json:"account_holder_name,omitempty"`
		AccountNumber     string `json:"account_number,omitempty"`
		BankName          string `json:"bank_name,omitempty"`
		BankAddress       string `json:"bank_address,omitempty"`
		SortCode          string `json:"sort_code,omitempty"`
		RoutingNumber     string `json:"routing_number,omitempty"`
		IfscCode          string `json:"ifsc_code,omitempty"`
		RoutingCode       string `json:"routing_code,omitempty"`
	} `json:"payment_details,omitempty"`
}

// RecordsRack holds all records
type RecordsRack struct {
	Invoices   []Invoice `json:"records"`
	TotalCount string    `json:"total_count"`
	Start      int       `json:"start"`
	Offset     int       `json:"offset"`
}

// Countries is a slice of Country
type Countries []Country

// Country represent a country
type Country struct {
	PrimUID     string `json:"prim_uid"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	Vat         string `json:"vat"`
	IsEu        string `json:"is_eu"`
}
