package efiscal

import (
	"github.com/jmcvetta/napping"

	"errors"
	"net/url"
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

func Boot(user, key string, production bool) API {
	
	return API{user, key, production}	
}

func (module API) Invoice(rfc, series, folio string) *Invoice {

	mode := "debug"

	if module.production {
		mode = "produccion"
	}

	invoice := &Invoice{
		Mode:      mode,
		Version:   EF_VERSION,
		Subtotal:  0.0,
		Discounts: 0.0,
		Total:     0.0,
		Decimals:  EF_DECIMALS,
		Series:    series,
		Folio:     folio,
		Emitted:   JSONTime(time.Now()),
		RFC:       rfc,
	}

	return invoice
}

func (module API) Sign(invoice *Invoice) (map[string]interface{}, error) {

	var res Response

	invoice.Prepare()

	s := napping.Session{
		Userinfo: url.UserPassword(module.user, module.key),
	}

	payload := struct {
		CFDI *Invoice `json:"CFDi"`
	}{
		invoice,
	}

	url := EF_SERVICE + "comprobantes/emitir"
	response, err := s.Post(url, &payload, &res, nil)

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
