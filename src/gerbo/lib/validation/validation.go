// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

// lib
package validation

import (
	"github.com/fatih/structs"
	liberr "gerbo/lib/error"
	"reflect"
)

func isAlpha(r string) bool {
	str := []rune(r)
	if len(str) < 1 {
		return false
	}
	return int(str[0]) >= 48 && int(str[0]) <= 57
}

// Check the values the map passed
func MustBeNotEmpty(v interface{}, require func() []string, sub func() map[string][]string) (string, error) {
	required := require()
	subitems := sub()

	s := structs.New(v)
	for _, v := range required {
		name := s.Field(v)
		for key, values := range subitems {
			if key == v {
				for i:=0; i < len(values); i++ {
					if name.Field(values[i]).Value() == "" {
						return values[i], &liberr.Err{Name: values[i] + " - Parametro incorreto"}
					}
				}
			}
		}
		if name.Kind() == reflect.String {
			value := name.Value().(string)
			if !isAlpha(value) || value == ""{
				return v, &liberr.Err{Name: v + " - Parametro incorreto"}
			}
		}
	}
	return "", nil
}
