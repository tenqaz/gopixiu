package util

import (
	ut "github.com/go-playground/universal-translator"
	"strings"
)

var Trans ut.Translator

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
