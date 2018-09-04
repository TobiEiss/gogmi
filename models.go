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
