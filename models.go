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

type Record struct {
	PrimUID        string `json:"prim_uid"`
	SupplierUID    string `json:"supplier_uid"`
	InvoiceNumber  string `json:"invoice_number"`
	InvoiceDate    string `json:"invoice_date"`
	InvoiceDueDate string `json:"invoice_due_date"`
	NetAmount      string `json:"net_amount"`
	Vat            string `json:"vat"`
	GrossAmount    string `json:"gross_amount"`
	Currency       string `json:"currency"`
	IsArchived     string `json:"is_archived"`
	IsOcrCompleted int    `json:"is_ocr_completed"`
	Tags           string `json:"tags"`
	Note           string `json:"note"`
	Source         string `json:"source"`
	Filename       string `json:"filename"`
	FileSize       string `json:"file_size"`
	PaymentStatus  string `json:"payment_status"`
	PaymentMethod  string `json:"payment_method"`
	PaymentDetails struct {
		Iban              string `json:"iban"`
		Bic               string `json:"bic"`
		AccountHolderName string `json:"account_holder_name"`
		AccountNumber     string `json:"account_number"`
		BankName          string `json:"bank_name"`
		BankAddress       string `json:"bank_address"`
		SortCode          string `json:"sort_code"`
		RoutingNumber     string `json:"routing_number"`
		IfscCode          string `json:"ifsc_code"`
		RoutingCode       string `json:"routing_code"`
	} `json:"payment_details,omitempty"`
}

type RecordsRack struct {
	Records    []Record `json:"records"`
	TotalCount string   `json:"total_count"`
	Start      int      `json:"start"`
	Offset     int      `json:"offset"`
}
