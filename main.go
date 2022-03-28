package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
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

func parseJson(filename string) (map[string]chapter, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var adventure map[string]chapter
	if err = json.Unmarshal(data, &adventure); err != nil {
		return nil, err
	}

	return adventure, nil
}

func main() {
	adventure, err := parseJson("gopher.json")
	if err != nil {
		log.Fatalln(err)
	}

	templates := template.Must(template.ParseGlob("tmpl/*"))

	handler := &ChooseAdventureHandler{
		adventure: adventure,
		templates: templates,
	}
	http.Handle("/", handler)

	//playChapter(adventure, adventure["intro"])

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type ChooseAdventureHandler struct {
	adventure map[string]chapter
	templates *template.Template
}

func (h *ChooseAdventureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	err = h.templates.ExecuteTemplate(w, "index.html", h.adventure["intro"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

//func renderTemplate(w http.ResponseWriter, tmpl string, c chapter) {
//	err := templates.ExecuteTemplate(w, tmpl+".html", p)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
