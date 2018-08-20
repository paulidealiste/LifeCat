// Package catalogueoflife sends and processes requestst to the http://www.catalogueoflife.org
package catalogueoflife

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HigherTaxa encompasses some of the fields from full catalogueoflife
type HigherTaxa struct {
	Name string `json:"name"`
	Rank string `json:"rank"`
}

// CollectionResult encompasses some of the fields from full catalogueoflife
// +gen slice:"Where,Select[string]"
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

// CollectionObject encompasses some of the fields from full catalogueoflife
type CollectionObject struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Leng    int    `json:"total_number_of_results"`
	Results CollectionResultSlice
}

// ReadAndUnmarsh returns the CollectionObject from catalogueoflife request
func ReadAndUnmarsh(t1 string, t2 string) CollectionObject {
	var todec CollectionObject
	rupquer := "http://webservice.catalogueoflife.org/col/webservice?name=" + t1 + "+" + t2 + "&format=json&response=full"
	resp, err := http.Get(rupquer)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &todec)
	return todec
}

// PrintTaxon prints the requested taxon and it classifications
func PrintTaxon(todex *CollectionObject) {
	ranks := todex.Results.SelectString(getRanks)
	fmt.Println()
	fmt.Printf("Taxa for request: %s\n", mockUniqe(ranks))
	for _, rnk := range ranks {
		rankfilter := func(cr CollectionResult) bool { return cr.Rank == rnk }
		onlytx := todex.Results.Where(rankfilter)
		fmt.Println()
		fmt.Println("Classifications:")
		for _, tdc := range onlytx {
			fmt.Println()
			fmt.Printf("For %s \n", tdc.Name)
			for _, cls := range tdc.Classification {
				fmt.Printf("%s: %s\n", cls.Rank, cls.Name)
			}
		}
	}
}

func getRanks(cr CollectionResult) string {
	return cr.Rank
}

func mockUniqe(forst []string) []string {
	u := make([]string, 0, len(forst))
	m := make(map[string]bool)
	for _, vava := range forst {
		if _, ok := m[vava]; !ok {
			m[vava] = true
			u = append(u, vava)
		}
	}
	return u
}
