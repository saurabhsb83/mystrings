package mystrings

import (
	"bytes"
	"fmt"
)

func Reverse(s string) string {

	result := ""
	for _, v := range s {

		result = string(v) + result
	}
	return result

}

func Reversev2(s string) {

	result := ""
	for _, v := range s {

		result = string(v) + result
	}

	buf := bytes.NewBufferString(result)
	fmt.Fprintln(buf)

}
