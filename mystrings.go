package mystrings

import (
	"bytes"
	"fmt"
)

func Reversev1(s string) string {

	result := ""
	for _, v := range s {

		result = string(v) + result
	}
	return result

}

func Reverse(s string) {

	result := ""
	for _, v := range s {

		result = string(v) + result
	}

	buf := bytes.NewBufferString(result)
	fmt.Fprintln(buf)

}
