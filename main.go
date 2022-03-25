package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type chapter struct {
	Title   string
	Story   []string
	Options []option
}
type option struct {
	Text string
	Arc  string
}

func main() {
	data, err := os.ReadFile("gopher.json")
	if err != nil {
		log.Fatalln(err)
	}

	var adventure map[string]chapter
	if err = json.Unmarshal(data, &adventure); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(adventure)
}
