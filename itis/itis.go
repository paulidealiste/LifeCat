// Package itis provides simple calls from ITIS https://www.itis.gov
package itis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Hierarchy describes one OTU
type Hierarchy struct {
	Author     string `json:"author"`
	TaxonName  string `json:"taxonName"`
	RankName   string `json:"rankName"`
	ParentName string `json:"parentName"`
}

//TaxonHierarchy encompasses all of the taxonomic records for the requested taxon
// +gen slice:"Where"
type TaxonHierarchy struct {
	SciName       string      `json:"sciName"`
	RankName      string      `json:"rankName"`
	HierarchyList []Hierarchy `json:"hierarchyList"`
}

//CommonName depicts the taxon common name in a designated language
type CommonName struct {
	CommonName string `json:"commonName"`
	Languange  string `json:"language"`
}

//TaxonInfo is the current info on the requested taxon
// +gen slice:"Where, Select[string]"
type TaxonInfo struct {
	Author         string
	ScientificName string
	CommonNames    []CommonName `json:"commonNames"`
	RankName       string       `json:"rankName"`
}

//ScinamesInfo is used for retrieval of the TSN number for the requested taxon
// +gen slice:"Where, Select[TaxonInfo], Select[TaxonHierarchy]"
type ScinamesInfo struct {
	Author string `json:"author"`
	Name   string `json:"combinedName"`
	TSN    string `json:"tsn"`
}

//Container provides one all-encompassing data type for an ITIS object of the requested taxon
type Container struct {
	ScientificInfos ScinamesInfoSlice `json:"scientificNames"`
	TaxonInfos      TaxonInfoSlice
	Hierarchy       TaxonHierarchySlice
}

//ReadAndUnmarsh constructs new ITISObject by evoking several ITIS API methods
func ReadAndUnmarsh(t1 string, t2 string) Container {
	chanchan := make(chan bool, 2)
	querykey := t1
	if t2 != "" {
		querykey += "+" + t2
	}
	rupquer := "http://www.itis.gov/ITISWebService/jsonservice/searchByScientificName?srchKey=" + querykey
	resp, err := http.Get(rupquer)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var epona Container
	json.Unmarshal(body, &epona)
	go epona.fillTaxonInfo(chanchan)
	go epona.fillTaxonHierarchy(chanchan)
	<-chanchan
	<-chanchan
	return epona
}

// PrintTaxon prints the requested taxon, its author and its classification along with its common names
func PrintTaxon(godex *Container) {
	fmt.Printf("Taxa for ITIS request %s\n", godex.TaxonInfos.SelectString(getTaxas))
	for _, info := range godex.TaxonInfos {
		fmt.Println()
		fmt.Printf("-%s (%s)\n", info.ScientificName, info.Author)
		for _, cmnn := range info.CommonNames {
			fmt.Printf("-%s: %s\n", cmnn.Languange, cmnn.CommonName)
		}
	}
	fmt.Println()
	fmt.Println("Classification:")
	fmt.Println()
	for _, class := range godex.Hierarchy {
		for _, inclass := range class.HierarchyList {
			fmt.Printf("%s: %s \n", inclass.RankName, inclass.TaxonName)
		}
	}
}

func getTaxas(ti TaxonInfo) string {
	return ti.RankName
}

func (ep *Container) fillTaxonInfo(chanchan chan bool) {
	ep.TaxonInfos = ep.ScientificInfos.SelectTaxonInfo(tinfilMapper)
	chanchan <- true
}

func tinfilMapper(sni ScinamesInfo) TaxonInfo {
	tinfo := TaxonInfo{Author: sni.Author, ScientificName: sni.Name}
	rupquer := "https://www.itis.gov/ITISWebService/jsonservice/getTaxonomicRankNameFromTSN?tsn=" + sni.TSN
	resp, err := http.Get(rupquer)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &tinfo)
	rupquer = "https://www.itis.gov/ITISWebService/jsonservice/getCommonNamesFromTSN?tsn=" + sni.TSN
	resp, err = http.Get(rupquer)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body2, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body2, &tinfo)
	return tinfo
}

func (ep *Container) fillTaxonHierarchy(chanchan chan bool) {
	ep.Hierarchy = ep.ScientificInfos.SelectTaxonHierarchy(hirfilMapper)
	chanchan <- true
}

func hirfilMapper(sni ScinamesInfo) TaxonHierarchy {
	var tihir TaxonHierarchy
	rupquer := "https://www.itis.gov/ITISWebService/jsonservice/getFullHierarchyFromTSN?tsn=" + sni.TSN
	resp, err := http.Get(rupquer)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &tihir)
	return tihir
}
