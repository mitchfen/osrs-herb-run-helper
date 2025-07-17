package helpers

import (
	"fmt"
	"sort"

	"github.com/Jeffail/gabs/v2"
)

const numberofHerbsPerSeed float64 = 9.423 // Assumes farming cape and magic secateurs
const numberofPatches float64 = 9.0

type Herb struct {
	Name           string
	Id             int
	SeedId         int
	HerbPrice      float64
	SeedPrice      float64
	ExpectedProfit float64
}

func (h Herb) FilterValue() string { return h.Name }
func (h Herb) Title() string       { return h.Name }
func (h Herb) Description() string { return fmt.Sprintf("Expected profit: %.2f", h.ExpectedProfit) }

func BuildHerbsSlice(parsedPriceJson *gabs.Container) []Herb {

	Ranarr := &Herb{
		Name:   "Ranarr",
		Id:     257,
		SeedId: 5295,
	}

	SnapDragon := &Herb{
		Name:   "Snapdragon",
		Id:     3000,
		SeedId: 5300,
	}

	Torstol := &Herb{
		Name:   "Torstol",
		Id:     269,
		SeedId: 5304,
	}

	Toadflax := &Herb{
		Name:   "Toadflax",
		Id:     2998,
		SeedId: 5296,
	}

	Cadantine := &Herb{
		Name:   "Cadantine",
		Id:     265,
		SeedId: 5301,
	}

	herbs := []Herb{*Ranarr, *SnapDragon, *Torstol, *Toadflax, *Cadantine}

	// Fill in the price values for each herb in the array
	for i := 0; i < len(herbs); i++ {
		herbHighPrice := parsedPriceJson.Search("data", fmt.Sprint(herbs[i].Id), "avgHighPrice").Data().(float64)
		herbLowPrice := parsedPriceJson.Search("data", fmt.Sprint(herbs[i].Id), "avgLowPrice").Data().(float64)
		herbs[i].HerbPrice = (herbHighPrice + herbLowPrice) / 2
		seedHighPrice := parsedPriceJson.Search("data", fmt.Sprint(herbs[i].SeedId), "avgHighPrice").Data().(float64)
		seedLowPrice := parsedPriceJson.Search("data", fmt.Sprint(herbs[i].SeedId), "avgLowPrice").Data().(float64)
		herbs[i].SeedPrice = (seedHighPrice + seedLowPrice) / 2
		expectedHerbsReturned := numberofHerbsPerSeed * numberofPatches
		costOfSeeds := float64(herbs[i].SeedPrice) * numberofPatches
		valueOfGatheredHerbs := float64(herbs[i].HerbPrice) * expectedHerbsReturned
		herbs[i].ExpectedProfit = valueOfGatheredHerbs - costOfSeeds
	}

	// Sort the array of herbs according to expected profit
	sort.Slice(herbs, func(i, j int) bool {
		return herbs[i].ExpectedProfit > herbs[j].ExpectedProfit
	})

	return herbs
}
