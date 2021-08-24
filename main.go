package main

import (
	"log"
	"os"
)

const (
	in  = "example.txt"
	out = "treated_example.txt"
)

func main() {
	s, err := os.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			s[i] = '{'
		} else if s[i] == ']' {
			s[i] = '}'
		}
	}

	treated, _ := os.OpenFile(out, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	treated.Write(s)
}
