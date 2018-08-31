// Package lifecatsink provides collection data structures for visualization and utilities for their preparation
package lifecatsink

import (
	"github.com/paulidealiste/LifeCat/catalogueoflife"
	"github.com/paulidealiste/LifeCat/itis"
	"github.com/paulidealiste/LifeCat/lifecatutil"
)

// Taxon describes any available taxon
// +gen slice:"Where"
type Taxon struct {
	rank int
	Name string
}

// Taxa is a list of all possible/allowed taxons
var Taxa TaxonSlice

func init() {
	Taxa = TaxonSlice{
		Taxon{rank: 0, Name: "Regnum"},
		Taxon{rank: 1, Name: "Subregnum"},
		Taxon{rank: 2, Name: "Infraregnum"},
		Taxon{rank: 3, Name: "Phylum"},
		Taxon{rank: 4, Name: "Subphylum"},
		Taxon{rank: 5, Name: "Infraphylum"},
		Taxon{rank: 6, Name: "Superclassis"},
		Taxon{rank: 7, Name: "Classis"},
		Taxon{rank: 8, Name: "Subclassis"},
		Taxon{rank: 9, Name: "Infraclassis"},
		Taxon{rank: 10, Name: "Superordo"},
		Taxon{rank: 11, Name: "Ordo"},
		Taxon{rank: 12, Name: "Subordo"},
		Taxon{rank: 13, Name: "Superfamilia"},
		Taxon{rank: 14, Name: "Familia"},
		Taxon{rank: 15, Name: "Subfamilia"},
		Taxon{rank: 16, Name: "Genus"},
		Taxon{rank: 17, Name: "Species"},
		Taxon{rank: 18, Name: "Subspecies"},
		Taxon{rank: 19, Name: "Infraspecies"},
	}
}

// LifeCatTaxonomy defines universal taxonomy for any of the possible inputs
type LifeCatTaxonomy struct {
	Teleos   LifeCatOTU
	Taxonomy LifeCatOTUSlice
}

// LifeCatOTU defines one opertational taxonomic unit
// +gen slice:"Where, SortBy, GroupBy[string]"
type LifeCatOTU struct {
	Taxon Taxon
	Name  string
}

// AnyToTaxonInfo contructs new general taxon info from any of the possible linputs
func (lfc *LifeCatTaxonomy) AnyToTaxonInfo(tir interface{}) {
	if ccol, ok := tir.(catalogueoflife.CollectionResultSlice); ok {
		for _, co := range ccol {
			ts, err := lifecatutil.EngLatTaxon(co.Rank)
			if err != nil {
				panic(err)
			}
			lfc.Teleos = fromRankName(ts, co.Name)
		}
	}
	if icol, ok := tir.(itis.TaxonHierarchySlice); ok {
		for _, co := range icol {
			ts, err := lifecatutil.EngLatTaxon(co.RankName)
			if err != nil {
				panic(err)
			}
			lfc.Teleos = fromRankName(ts, co.SciName)
		}
	}
}

// AnyToTaxonomy contructs new general taxonomy from any of the possible inputs
func (lfc *LifeCatTaxonomy) AnyToTaxonomy(hir interface{}) {
	lfc.Taxonomy = []LifeCatOTU{}
	if hft, ok := hir.([]catalogueoflife.HigherTaxa); ok {
		for _, ht := range hft {
			ts, err := lifecatutil.EngLatTaxon(ht.Rank)
			if err != nil {
				panic(err)
			}
			newotu := fromRankName(ts, ht.Name)
			lfc.Taxonomy = append(lfc.Taxonomy, newotu)
		}
	}
	if ift, ok := hir.([]itis.Hierarchy); ok {
		for _, ht := range ift {
			ts, err := lifecatutil.EngLatTaxon(ht.RankName)
			if err != nil {
				panic(err)
			}
			newotu := fromRankName(ts, ht.TaxonName)
			lfc.Taxonomy = append(lfc.Taxonomy, newotu)
		}
	}
	lfc.Taxonomy = lfc.Taxonomy.SortBy(func(arg1 LifeCatOTU, arg2 LifeCatOTU) bool {
		return arg1.Taxon.rank < arg2.Taxon.rank
	})
}

func fromRankName(rank string, name string) LifeCatOTU {
	cutax := Taxa.Where(func(arg1 Taxon) bool {
		return arg1.Name == rank
	})
	return LifeCatOTU{Taxon: cutax[0], Name: name}
}
