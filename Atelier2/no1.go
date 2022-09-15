package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func numOne(s string) {

	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "'", " ")

	var dist = make(map[string]int)

	re, err := regexp.Compile(`[\.,\(\)]`)
	if err != nil {
		log.Fatalln(err)
	}

	s = re.ReplaceAllString(s, "")

	var words []string = strings.Fields(s)

	for _, word := range words {
		if count, ok := dist[word]; ok {
			dist[word] = count + 1
		} else {
			dist[word] = 1
		}
	}
	fmt.Printf("There u go : %v", dist)
}
