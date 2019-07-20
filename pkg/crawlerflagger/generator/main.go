package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Crawler struct {
	Pattern      string   `json:"pattern"`
	Instances    []string `json:"instances"`
	URL          *string  `json:"url"`
	Description  *string  `json:"description"`
	AdditionDate *string  `json:"additional_date"`
	DependsOn    *string  `json:"depends_on"`
}

func main() {
	var crawlers []Crawler
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get("https://raw.githubusercontent.com/monperrus/crawler-user-agents/master/crawler-user-agents.json")
	if err != nil {
		log.Fatalf("while retrieving the list of crawler, an error occurred: %v", err)
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&crawlers)
	if err != nil {
		log.Fatalf("an error occurred while decoding the JSON: %v", err)
	}
	var tplOutput bytes.Buffer
	f, err := ioutil.ReadFile("template.go.tpl")
	if err != nil {
		log.Fatalf("%v", err)
	}
	tpl := template.Must(template.New("main").Parse(string(f)))
	if err := tpl.Execute(&tplOutput, crawlers); err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(tplOutput.String())
}
