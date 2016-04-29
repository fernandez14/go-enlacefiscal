package efiscal

type Receiver struct {
	RFC     string   `json:"rfc"`
	Name    string   `json:"nombre"`
	Address *Address `json:"DomicilioFiscal,omitempty"`
}

type Address struct {
	Street       string `json:"calle"`
	Ext          string `json:"noExterior"`
	Int          string `json:"noInterior"`
	Neighborhood string `json:"colonia"`
	Locality     string `json:"localidad"`
	Town         string `json:"municipio"`
	State        string `json:"estado"`
	Country      string `json:"pais"`
	Zipcode      string `json:"cp"`
}
