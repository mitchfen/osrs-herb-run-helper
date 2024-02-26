package main

import(
	"github.com/mitchfen/osrs_herb_run_helper/internal/helpers"
)

func main() {
	parsedPriceJson := helpers.GetParsedPriceJson()
	herbs := helpers.BuildHerbsSlice(parsedPriceJson)
	helpers.PrintResults(herbs)
}
