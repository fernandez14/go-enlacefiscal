package efiscal

type Payment struct {
	Method string `json:"metodoDePago"`
	Type   string `json:"formaDePago"`
}

var PAY_ONE_TIME_TRANSFER Payment = Payment{
	Method: "Transferencia Electronica",
	Type:   "Pago en una sola exhibición",
}
