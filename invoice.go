package efiscal

import (
	"time"
)

// Invoice is a representation of a sign request to EnlaceFiscal.com
type Invoice struct {
	mode      string    `json:"modo"`
	version   string    `json:"versionEF"`
	subtotal  float64   `json:"subTotal"`
	discounts float64   `json:"descuentos"`
	total     float64   `json:"total"`
	decimals  int       `json:"numeroDecimales"`
	Series    string    `json:"serie"`
	Folio     string    `json:"folioInterno"`
	Emitted   time.Time `json:"fechaEmision"`
	RFC       string    `json:"rfc"`

	// Nested structures
	Payment  *Payment  `json:"DatosDePago,omitempty"`
	Receiver *Receiver `json:"Receptor,omitempty"`
	Items    *Items    `json:"Partidas,omitempty"`
	Taxes    *Taxes    `json:"Impuestos,omitempty"`
}

// Sets the current invoice representation payment settings
func (self *Invoice) SetPayment(v *Payment) {
	self.Payment = v
}

// Sets the current invoice representation receiver information (target rfc)
func (self *Invoice) SetReceiver(v *Receiver) {
	self.Receiver = v
}

// Sets the current invoice emitted time
func (self *Invoice) SetEmitted(t time.Time) {
	self.Emitted = t
}

func (self *Invoice) GetSubtotal() float64 {
	return self.subtotal
}

func (self *Invoice) GetTotal() float64 {
	return self.total
}

// Add an Item to the current invoice representation
func (self *Invoice) AddItem(i Item) {

	if self.Items == nil {
		self.Items = &Items{make([]Item, 0)}
	}

	if i.Total == 0.0 {
		i.Total = float64(i.Quantity) * i.Value
	}

	self.Items.List = append(self.Items.List, i)
	self.subtotal += i.Total
	self.total += i.Total
}

// Transfers IVA Taxes to the current representation
func (self *Invoice) TransferIVA(rate float64) {

	self.prepareTaxes()

	tax := Tax{
		Type:  "IVA",
		Rate:  rate,
		Total: (self.total / 100 * rate),
	}

	self.total += tax.Total
	self.Taxes.Transfers = append(self.Taxes.Transfers, tax)
}

func (self *Invoice) prepareTaxes() {

	if self.Taxes == nil {
		self.Taxes = &Taxes{make([]Tax, 0)}
	}
}

// Prepare the Invoice representation for signin
func (self *Invoice) Prepare() {

}
