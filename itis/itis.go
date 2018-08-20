// package itis provides simple calls from ITIS https://www.itis.gov
package itis

import (
	"encoding/json"
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

//AcceptedName is the current scientific name of the requested taxon
type AcceptedName struct {
	AcceptedName string `json:"acceptedName"`
}

//CommonEnglighName is the common requested taxon name
type CommonEnglighName struct {
	CommonName string `json:"commonName"`
}

//TaxonInfo is the current info on the requested taxon
type TaxonInfo struct {
	AcceptedNameList []AcceptedName      `json:"acceptedNameList"`
	CommonNameList   []CommonEnglighName `json:"commonNameList"`
	TaxRank          string              `json:"taxRank.rankName"`
}

//ScinamesInfo is used for retrieval of the TSN number for the requested taxon
type ScinamesInfo struct {
	Author string `json:"author"`
	Name   string `json:"combinedName"`
	TSN    string `json:"tsn"`
}

//Container provides one all-encompassing data type for an ITIS object of the requested taxon
type Container struct {
	ScientificNames []ScinamesInfo `json:"scientificNames"`
}

//ReadAndUnmarsh constructs new ITISObject by evoking several ITIS API methods
func ReadAndUnmarsh(t1 string, t2 string) Container {
	querykey := t1 + "+" + t2
	rupquer := "http://www.itis.gov/ITISWebService/jsonservice/searchByScientificName?srchKey=" + querykey
	resp, err := http.Get(rupquer)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var epona Container
	json.Unmarshal([]byte(body), &epona)
	return epona
}
