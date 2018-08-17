package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HigherTaxa struct {
	Name string `json:"name"`
	Rank string `json:"rank"`
}

// +gen slice:"Where"
type CollectionResult struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	Rank           string       `json:"rank"`
	Status         string       `json:"name_status"`
	AuthoredName   string       `json:"name_html"`
	Bib            string       `json:"bibliographic_citation"`
	Extinct        bool         `json:"is_extinct"`
	Classification []HigherTaxa `json:"classification"`
}

type CollectionObject struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Leng    int    `json:"total_number_of_results"`
	Results CollectionResultSlice
}

func main() {

	var todec CollectionObject
	var t1 string
	var t2 string

	flag.StringVar(&t1, "taxa1", "Rupicapra", "string")
	flag.StringVar(&t2, "taxa2", "Rupicapra", "string")

	flag.Parse()

	rupquer := "http://webservice.catalogueoflife.org/col/webservice?name=" + t1 + "+" + t2 + "&format=json&response=full"
	resp, err := http.Get(rupquer)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &todec)
	fmt.Println("Species found:")
	onlySpecies := todec.Results.Where(isSpecies)
	for _, sp := range onlySpecies {
		fmt.Printf("-%s\n", sp.Name)
	}
	fmt.Println("Classifications:")
	for _, tdc := range onlySpecies {
		for _, cls := range tdc.Classification {
			fmt.Printf("+r+:%s +n+:%s\n", cls.Rank, cls.Name)
		}
	}
}

func isSpecies(cr CollectionResult) bool {
	return cr.Rank == "Species"
}
