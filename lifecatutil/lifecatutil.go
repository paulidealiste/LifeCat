// Package lifecatutil provides some utility functions common for all API-based results
package lifecatutil

import (
	"errors"
	"os"
	"path/filepath"
)

var taxotrans map[string]string

func init() {
	taxotrans = map[string]string{
		"Kingdom":      "Regnum",
		"Subkingdom":   "Subregnum",
		"Infrakingdom": "Infraregnum",
		"Phylum":       "Phylum",
		"Subphylum":    "Subphylum",
		"Infraphylum":  "Infraphylum",
		"Superclass":   "Superclassis",
		"Class":        "Classis",
		"Subclass":     "Subclassis",
		"Infraclass":   "Infraclassis",
		"Superorder":   "Superordo",
		"Order":        "Ordo",
		"Suborder":     "Subordo",
		"Superfamily":  "Superfamilia",
		"Family":       "Familia",
		"Subfamily":    "Subfamilia",
		"Genus":        "Genus",
		"Species":      "Species",
		"Subspecies":   "Subspecies",
		"Infraspecies": "Infraspecies",
	}
}

//MockUnique accepts a list of strings and prunes all of the duplicated ones
func MockUnique(forst []string) []string {
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

//GetCWD returns the current working directory of the golang executable
func GetCWD() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

//EngLatTaxon returns latin name for a taxon in English using the translation map (taxontrans)
func EngLatTaxon(entax string) (string, error) {
	cutax := taxotrans[entax]
	if cutax == "" {
		return "", errors.New("no suitable translation found for your querry")
	}
	return cutax, nil
}
