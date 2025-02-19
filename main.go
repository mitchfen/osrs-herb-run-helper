package main

import (
	"github.com/mitchfen/osrs-herb-run-helper/internal/helpers"
)

func main() {
	parsedPriceJson := helpers.GetParsedPriceJson()
	herbs := helpers.BuildHerbsSlice(parsedPriceJson)
	helpers.PrintResults(herbs)
}
