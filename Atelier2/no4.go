package main

import (
	"fmt"
	"log"
	"regexp"
)

func numFour(s string) {

	re, err := regexp.Compile(`^\s*(?:\+?(\d{1,3}))?[-. (]*(\d{3})[-. )]*(\d{3})[-. ]*(\d{4})(?: *x(\d+))?\s*$`)
	if err != nil {
		log.Fatalln(err)
	}

	s = re.ReplaceAllString(s, "")
	fmt.Println(s)
}
