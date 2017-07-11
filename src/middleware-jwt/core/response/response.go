// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

// core
package response

import (
	"encoding/json"
	liberr "middleware-jwt/lib/error"
	"net/http"
)

type Body struct {
	Meta       Meta        `json:"meta"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Meta struct {
	Code   int             `json:"code"`
	Errors []liberr.Errors `json:"errors"`
	Url    string          `json:"url"`
	Method string          `json:"method"`
}

type Pagination struct {
	nextUrl   string `json:"next_url"`
	nextMaxId string `json:"next_max_id"`
}

type Headers map[string]string

// A maneira como funciona o http é o código de resposta enviado como um cabeçalho.
// Portanto, devemos definir o status da resposta antes de enviar o corpo, caso contrário,
// net/http assumirá a resposta sendo como (StatusOK) e configurá-lo para 200.
// Sendo um comportamento que está no pacote net/http
func Header(w http.ResponseWriter, status int, message interface{}, headers Headers) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

func Envelope(errAPI liberr.ErrorsAPI, status int, data interface{}) Body {
	envelope := Body{}

	envelope.Meta.Errors = errAPI.Errors
	envelope.Meta.Url = errAPI.Url
	envelope.Meta.Method = errAPI.Method
	envelope.Meta.Code = status
	envelope.Data = data

	return envelope
}
