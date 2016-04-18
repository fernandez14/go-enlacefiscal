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
	ref, re := self.Body["numeroReferencia"].(int)

	if exists && re {
		code, ce := message["codigoError"].(string)

		if ce {
			description, de := message["descripcionError"].(map[string]interface{})

			if de {
				text, te := description["texto"].(string)

				if te {

					return &ResponseError{code, ref, text}
				}
			}
		}
	}

	return errors.New("Response has no errors on it.")
}

type ResponseError struct {
	Code        string
	Reference   int
	Description string
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("%d - Code: %s - %s", e.Reference, e.Code, e.Description)
}
