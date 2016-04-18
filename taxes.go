package efiscal

type Taxes struct {
	Transfers []Tax `json:"Traslados,omitempty"`
}

type Tax struct {
	Type  string  `json:"impuesto,omitempty"`
	Rate  float64 `json:"tasa,omitempty"`
	Total float64 `json:"importe,omitempty"`
}
