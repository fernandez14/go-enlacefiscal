package efiscal

import (
	"errors"
	"fmt"
)

type Response struct {
	Body map[string]interface{} `json:"AckEnlaceFiscal"`
}

func (self Response) IsError() bool {

	_, exists := self.Body["mensajeError"]

	return exists
}

func (self Response) GetError() error {

	message, exists := self.Body["mensajeError"].(map[string]interface{})
	ref, re := self.Body["numeroReferencia"]

	if exists && re {
		code, ce := message["codigoError"]

		if ce {
			description, de := message["descripcionError"].(map[string]interface{})

			if de {
				text, te := description["texto"]

				if te {

					return &ResponseError{code, ref, text}
				}
			}
		}
	}

	return errors.New(message["codigoError"].(string))
}

type ResponseError struct {
	Code        interface{}
	Reference   interface{}
	Description interface{}
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("%v - Code: %v - %v", e.Reference, e.Code, e.Description)
}
