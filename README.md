# go-enlacefiscal
Golang EnlaceFiscal.com API wrapper

## Summary

Enlacefiscal.com API golang native wrapper

## Installation

```go get github.com/fernandez14/go-enlacefiscal```

## Usage

For further information about implementing and using this wrapper please go to [GoDoc](https://godoc.org/github.com/fernandez14/go-enlacefiscal) documentation.

```go
api := efiscal.API{"USERNAME", "AUTH_KEY", false} // Set third param to true in Production

// Signing RFC, Series and Folio (refer to EnlaceFiscal docs)
invoice := api.Invoice("XAXX010101000", "TEST", "1")

item := efiscal.Item{
	Quantity:    1,
	Value:       15.0,
	Unit:        "servicio",
	Description: "Servicio de Prueba",
}

// Add an item using Item struct
invoice.AddItem(item)

// Transfer current items IVA to the list of taxes
invoice.TransferIVA(16)

// One time payment using bank transfer
invoice.SetPayment(&efiscal.PAY_ONE_TIME_TRANSFER)

// Target RFC & customer
invoice.SetReceiver(&efiscal.Receiver{"XAXX010101000", "Publico en General"})

// Sign invoice
data, err := api.Sign(invoice)

if err != nil {
  panic(err)
}
```
