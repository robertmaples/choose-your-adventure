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

	playChapter(adventure, adventure["intro"])
}

func playChapter(a map[string]chapter, c chapter) {
	fmt.Printf("\n\n%s\n\n", c.Title)
	for _, line := range c.Story {
		fmt.Printf("\t%s\n\n", line)
	}

	if len(c.Options) == 0 {
		os.Exit(0)
	}

	for i, o := range c.Options {
		fmt.Printf("%v.) %s\n", i+1, o.Text)
	}

	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		// TODO: make prompt repeating for incorrect input
		log.Fatalln(err)
	}

	index := input - 1
	if index < 0 || index >= len(c.Options) {
		log.Fatalln("Input must be one of the available options.")
	}

	arc := c.Options[index].Arc

	playChapter(a, a[arc])
}
