package main

import (
	"fmt"

	"github.com/mkideal/cli"
)

type argT struct {
	Help    bool   `cli:"h, help" usage"Shows gratuitous help"`
	Taxon   string `cli:"t,taxon" usage:"Genus or higher taxon latin name"`
	Species string `cli:"s,species" usage:"Species latin name"`
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
			ctx.String("%s\n", ctx.Color().Red("At the very least a taxon name should be provided via the -t flag"))
		} else {
			ctx.String("%s%s %s\n", ctx.Color().Blue("Await for the info on "), argv.Taxon, argv.Species)
		}
		fmt.Println()
		return nil
	})

	// todex := catalogueoflife.ReadAndUnmarsh(t1, t2)
	// catalogueoflife.PrintTaxon(&todex)

	// godex := itis.ReadAndUnmarsh(t1, t2)

}
