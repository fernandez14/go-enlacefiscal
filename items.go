package efiscal

type Items struct {
	List []Item `json:"Partida"`
}

type Item struct {
	Quantity    int     `json:"cantidad"`
	Unit        string  `json:"unidad"`
	Description string  `json:"descripcion"`
	Value       float64 `json:"valorUnitario"`
	Total       float64 `json:"importe"`
}
