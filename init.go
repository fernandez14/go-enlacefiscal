package efiscal

import (
	"github.com/jmcvetta/napping"

	"errors"
	"time"
)

type API struct {
	user       string
	key        string
	production bool
}

const EF_VERSION = "5.0"
const EF_DECIMALS = 2
const EF_SERVICE = "https://api.enlacefiscal.com/rest/v1/"

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

func (module API) Sign(invoice *Invoice) (map[string]interface{}, error) {

	var res Response

	invoice.Prepare()

	url := EF_SERVICE + "comprobantes/emitir"
	response, err := napping.Post(url, invoice, &res, nil)

	if err != nil {
		return res.Body, err
	}

	if response.Status() != 200 {
		return res.Body, errors.New("EnlaceFiscal invalid response. Did not get http 200")
	}

	if res.IsError() {
		return res.Body, res.GetError()
	}

	return res.Body, nil
}
