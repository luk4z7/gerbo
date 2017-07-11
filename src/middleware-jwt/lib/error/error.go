// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

// lib
package error

type ErrorsAPI struct {
	Errors []Errors `json:"errors"`
	Url    string   `json:"url"`
	Method string   `json:"method"`
}

type Errors struct {
	ParameterName string `json:"parameter_name"`
	Type          string `json:"type"`
	Message       string `json:"message"`
}

type Err struct {
	Name string
}

func (e *Err) Error() string {
	return e.Name
}

// Return panic when error != nil
func Check(e error, m string) {
	if e != nil {
		if m == "" {
			panic(m)
		} else {
			panic(e)
		}
	}
}