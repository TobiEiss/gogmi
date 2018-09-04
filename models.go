package goGMI

type Suppliers []Supplier

type Supplier struct {
	PrimUID      string `json:"prim_uid"`
	Name         string `json:"name"`
	SupplierType string `json:"supplier_type"`
	Note         string `json:"note"`
	Tags         string `json:"tags"`
}
