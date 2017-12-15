package main

import (
	"fmt"
	"strconv"
)

// FormatDefault - default formatter
func FormatDefault(value interface{}, _ map[string]string) string {
	return fmt.Sprintf("%v", value)
}

// FormatDiv - formatter with div with class
func FormatDiv(value interface{}, options map[string]string) string {
	var res string
	if cssClass, OK := options["class"]; OK {
		res = fmt.Sprintf(`<div class="%s">%v</div>`, cssClass, value)
	} else {
		res = fmt.Sprintf(`<div>%v</div>`, value)
	}

	return res
}

// FormatInt format column int
func FormatInt(value interface{}, _ map[string]string) string {
	return strconv.FormatInt(int64(value.(int)), 10)
}
