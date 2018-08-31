package lifecatsink

import (
	"testing"

	"github.com/paulidealiste/LifeCat/catalogueoflife"
	"github.com/paulidealiste/LifeCat/itis"
)

func TestAnyToTaxonomy(t *testing.T) {
	var mainbumb LifeCatTaxonomy
	testtaxa := []catalogueoflife.HigherTaxa{
		catalogueoflife.HigherTaxa{Name: "Lepomis gibbosus", Rank: "Species"},
		catalogueoflife.HigherTaxa{Name: "Lepomis", Rank: "Genus"},
	}
	mainbumb.AnyToTaxonomy(testtaxa)
	if mainbumb.Taxonomy[0].Name != "Lepomis" {
		t.Error("Ooops, either not sorted or not filled in at all.")
	}
	tusttaxa := []itis.Hierarchy{
		itis.Hierarchy{Author: "Lepomir", RankName: "Species", TaxonName: "Salmo trutta", ParentName: "Kakanic"},
		itis.Hierarchy{Author: "Lepomir", RankName: "Genus", TaxonName: "Salmo", ParentName: "Kakanic"},
	}
	mainbumb.AnyToTaxonomy(tusttaxa)
	if mainbumb.Taxonomy[0].Name != "Salmo" {
		t.Error("Ooops, either not sorted or not filled in at all.")
	}
}

func TestAnyToTaxonInfo(t *testing.T) {
	var mainbumb LifeCatTaxonomy
	testinfo := catalogueoflife.CollectionResultSlice{
		catalogueoflife.CollectionResult{Name: "Bumborosni", Rank: "Class"},
	}
	mainbumb.AnyToTaxonInfo(testinfo)
	if mainbumb.Teleos.Taxon.Name != "Classis" {
		t.Error("Ooops, neither teleos infor was set properly nor was it sorted adequately.")
	}
	tetsxinfo := itis.TaxonHierarchySlice{
		itis.TaxonHierarchy{SciName: "Semanjovz", RankName: "Order"},
	}
	mainbumb.AnyToTaxonInfo(tetsxinfo)
	if mainbumb.Teleos.Taxon.Name != "Ordo" {
		t.Error("Ooops, neither teleos info was set properly not info could be read")
	}
}
