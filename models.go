package gogmi

// Supplier is the model for GMI-supplier
type Supplier struct {
	PrimUID      string `json:"prim_uid"`
	Name         string `json:"name"`
	SupplierType string `json:"supplier_type"`
	Note         string `json:"note"`
	Tags         string `json:"tags"`
}

// Suppliers is a slice of suppliers
type Suppliers []Supplier
