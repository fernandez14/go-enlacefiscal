package efiscal

import (
	"time"
)

type API struct {
	user       string
	key        string
	production bool
}

const EF_VERSION = "5.0"
const EF_DECIMALS = 2

func (module API) Invoice(rfc, series, folio string) *Invoice {

	mode := "debug"

	if module.production {
		mode = "produccion"
	}

	invoice := &Invoice{
		mode:      mode,
		version:   EF_VERSION,
		subtotal:  0.0,
		discounts: 0.0,
		total:     0.0,
		decimals:  EF_DECIMALS,
		Series:    series,
		Folio:     folio,
		Emitted:   time.Now(),
		RFC:       rfc,
	}

	return invoice
}

func (module API) Sign(invoice *Invoice) (interface{}, error) {

	invoice.Prepare()

	return nil, nil
}
