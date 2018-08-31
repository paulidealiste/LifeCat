// Package lifecatpanel offers hierarchical rendering of queried results
package lifecatpanel

import (
	"bytes"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/paulidealiste/LifeCat/lifecatsink"
)

var tpl *template.Template

var pageTemplate = `
<!DOCTYPE html>
<html class="is-clipped">

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Wonderful taxonomy via LifeCat</title>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.1/css/bulma.min.css">
  <link href="https://unpkg.com/basscss@8.0.2/css/basscss.min.css" rel="stylesheet">
  <script defer src="https://use.fontawesome.com/releases/v5.1.0/js/all.js"></script>
</head>

<body>
  <section class="hero is-fullheight is-unselectable is-light">
    <div class="hero-head">
      <div class="container">
        <h1 class="title">
          Hello LifeCat Farer!
        </h1>
        <p class="subtitle">
          My first taxonomic UI with <strong>Bulma</strong>!
        </p>
      </div>
    </div>
    <div class="hero-body">
      <div class="container has-text-centered">
        <div id="taxonomyholder" class="notification">
        </div>
      </div>
    </div>
  </section>
</body>
</html>
`

var brickTemplate = `
<div class="box boxcox">
  <span class="tag is-light mt0 mb1">{{.Taxon.Name}}</span>
  <span class="tag is-info mt0 mb1">{{.Name}}</span>
</div>
`

var firebrickTemplate = `
<div class="box boxcox">
  <span class="tag is-warning is-large mt0 mb1">{{.Taxon.Name}}</span>
  <span class="tag is-primary is-large mt0 mb1">{{.Name}}</span>
</div>
`

// SubdivideHierarchy creates a nested structure from the LifeCatTaxonomy suitable for eventual json output
func SubdivideHierarchy(lfct lifecatsink.LifeCatTaxonomy) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageTemplate))
	if err != nil {
		log.Fatal(err)
	}
	brickscontainer := doc.Find("#taxonomyholder").AppendHtml(brickToString(brickTemplate, lfct.Taxonomy[0]))
	var newbrick *goquery.Selection
	for _, txn := range lfct.Taxonomy[1:] {
		newbrick = brickscontainer.Find(".boxcox").Last()
		newbrick.AppendHtml(brickToString(brickTemplate, txn))
	}
	newbrick = brickscontainer.Find(".boxcox").Last()
	newbrick.AppendHtml(brickToString(firebrickTemplate, lfct.Teleos))
	newrend, _ := doc.Html()
	tpl = template.Must(template.New("").Parse(newrend))
}

func brickToString(btem string, txn lifecatsink.LifeCatOTU) string {
	t := template.Must(template.New("").Parse(btem))
	var topi bytes.Buffer
	if err := t.Execute(&topi, txn); err != nil {
		log.Fatal(err)
	}
	return topi.String()
}

func nestedFlexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, "")
}
