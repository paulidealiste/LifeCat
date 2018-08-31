package main

import (
	"fmt"

	"github.com/mkideal/cli"
	"github.com/paulidealiste/LifeCat/catalogueoflife"
	"github.com/paulidealiste/LifeCat/itis"
)

type argT struct {
	Help    bool   `cli:"h, help" usage"Shows gratuitous help"`
	Taxon   string `cli:"t,taxon" usage:"Genus or higher taxon latin name"`
	Species string `cli:"s,species" usage:"Species latin name"`
	CatLif  bool   `cli:"c, catlif" usage:"Use Catalogue of Life API" dft:"false"`
	Itis    bool   `cli:"i, itis" usage:"Use ITIS API" dft:"false"`
}

func (argv *argT) AutoHelp() bool {
	return argv.Help
}

func main() {

	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		fmt.Println()
		ctx.String("%s", ctx.Color().Cyan("Wonderful taxonomy via"))
		ctx.String("%s\n\n", ctx.Color().Magenta(" LifeCat"))
		if argv.Taxon == "" {
			ctx.String("%s\n", ctx.Color().Red("At the very least a taxon name should be provided via the -t flag."))
		} else {
			if argv.CatLif || argv.Itis {
				lifchan := make(chan *catalogueoflife.CollectionObject, 1)
				itischan := make(chan *itis.Container, 1)
				fmt.Println()
				ctx.String("%s %s %s\n", ctx.Color().Blue("Await for the info on"), argv.Taxon, argv.Species)
				fmt.Println()
				if argv.CatLif {
					ctx.String("%s\n", ctx.Color().Green("An API and a database from http://www.catalogueoflife.org will be used."))
					go musterCatLif(argv.Taxon, argv.Species, lifchan, ctx)
				}
				if argv.Itis {
					ctx.String("%s\n", ctx.Color().Green("An API and a database from https://www.itis.gov will be used."))
					go musterITIS(argv.Taxon, argv.Species, itischan, ctx)
				}
				select {
				case lifob := <-lifchan:
					catalogueoflife.PrintTaxon(lifob)
				case itsob := <-itischan:
					itis.PrintTaxon(itsob)
				}
			} else {
				ctx.String("%s\n", ctx.Color().Red("Neither of the available APIs was selected."))
			}
		}
		fmt.Println()
		ctx.String("%s\n", ctx.Color().Grey("Please press any key to conclude."))
		fmt.Scanln()
		return nil
	})
}

func musterCatLif(t1 string, t2 string, apichan chan *catalogueoflife.CollectionObject, ctx *cli.Context) {
	catlif := catalogueoflife.ReadAndUnmarsh(t1, t2)
	apichan <- &catlif
}

func musterITIS(t1 string, t2 string, apichan chan *itis.Container, ctx *cli.Context) {
	itislif := itis.ReadAndUnmarsh(t1, t2)
	apichan <- &itislif
}
