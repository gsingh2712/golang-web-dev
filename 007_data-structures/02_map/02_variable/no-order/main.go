package main

import (
	"fmt"
)

func main() {

	sages := map[string]string{
		"India":        "Krishna",
		"America":      "MLK",
		"Meditate":     "Buddha",
		"Love":         "Jesus",
		"Revoltionary": "Savarkar"}

	for k, v := range sages {
		fmt.Println(k, " - ", v)
	}

}
