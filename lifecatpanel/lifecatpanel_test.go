package lifecatpanel

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paulidealiste/LifeCat/catalogueoflife"
	"github.com/paulidealiste/LifeCat/lifecatsink"
)

func TestSubdivideHierarchy(t *testing.T) {
	onecole := catalogueoflife.ReadAndUnmarsh("Alouatta", "belzebul")
	var lfct lifecatsink.LifeCatTaxonomy
	lfct.AnyToTaxonomy(onecole.Results[0].Classification)
	lfct.AnyToTaxonInfo(onecole.Results[0:1])
	fmt.Println(lfct)
	SubdivideHierarchy(lfct)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	hndlr := http.HandlerFunc(nestedFlexHandler)

	hndlr.ServeHTTP(rr, req)
	fmt.Println(rr.Body)
}
