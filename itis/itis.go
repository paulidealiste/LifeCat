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
type TaxonHierarchy struct {
	SciName       string      `json:"sciName"`
	RankName      string      `json:"rankName"`
	HierarchyList []Hierarchy `json:"hierarchyList"`
}

//TaxonInfo is the current info on the requested taxon
type TaxonInfo struct {
	Author         string
	ScientificName string
	CommonName     string
	TaxRank        string
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
	Info            []TaxonInfo
	Hierarchy       []TaxonHierarchy
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
	json.Unmarshal([]byte(body), &epona)
	go epona.fillTaxonInfo(chanchan)
	go epona.fillTaxonHierarchy(chanchan)
	<-chanchan
	<-chanchan
	return epona
}

func (ep *Container) fillTaxonInfo(chanchan chan bool) {
	ep.Info = ep.ScientificInfos.SelectTaxonInfo(tinfilMapper)
	chanchan <- true
}

func tinfilMapper(sni ScinamesInfo) TaxonInfo {
	var tinfo TaxonInfo
	var minkwo map[string]string
	rupquer := "https://www.itis.gov/ITISWebService/jsonservice/getTaxonomicRankNameFromTSN?tsn=" + sni.TSN
	resp, err := http.Get(rupquer)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, minkwo)
	fmt.Println(minkwo)
	return tinfo
}

func (ep *Container) fillTaxonHierarchy(chanchan chan bool) {
	ep.Hierarchy = ep.ScientificInfos.SelectTaxonHierarchy(hirfilMapper)
	chanchan <- true
}

func hirfilMapper(sni ScinamesInfo) TaxonHierarchy {
	var tihir TaxonHierarchy
	fmt.Println(sni.TSN)
	return tihir
}
