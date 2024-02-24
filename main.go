package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/Jeffail/gabs/v2"
)

type Herb struct {
	Name           string
	Id             int
	SeedId         int
	HerbPrice      float64
	SeedPrice      float64
	ExpectedProfit float64
}

func main() {
	priceJson, err := getPriceJson()
	if err != nil {
		fmt.Println("An error occurred while fetching the json data.")
		panic(err)
	}

	parsedJson, err := gabs.ParseJSON(priceJson)
	if err != nil {
		panic(err)
	}

	const numberofHerbsPerSeed float64 = 9.423
	const numberofPatches float64 = 9.0

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
		herbHighPrice := parsedJson.Search("data", fmt.Sprint(herbs[i].Id), "avgHighPrice").Data().(float64)
		herbLowPrice := parsedJson.Search("data", fmt.Sprint(herbs[i].Id), "avgLowPrice").Data().(float64)
		herbs[i].HerbPrice = (herbHighPrice + herbLowPrice) / 2
		seedHighPrice := parsedJson.Search("data", fmt.Sprint(herbs[i].SeedId), "avgHighPrice").Data().(float64)
		seedLowPrice := parsedJson.Search("data", fmt.Sprint(herbs[i].SeedId), "avgLowPrice").Data().(float64)
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

	// Print herb/seed prices to the console
	fmt.Printf("If you print %v seeds in %v patches, and assume %v herbs per patch. Then the expected profits are as follows:\n",
		numberofPatches, numberofPatches, numberofHerbsPerSeed)
	for i := 0; i < len(herbs); i++ {
		//fmt.Printf("%v herb price: %v\n", herbs[i].Name, herbs[i].HerbPrice)
		//fmt.Printf("%v seed price: %v\n", herbs[i].Name, herbs[i].SeedPrice)
		fmt.Printf("%v: %d\n", herbs[i].Name, int(herbs[i].ExpectedProfit))
	}
}

func getPriceJson() ([]byte, error) {
	url := "https://prices.runescape.wiki/api/v1/osrs/1h"
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "github.com/mitchfen/osrs_herb_run_helper")
	response, _ := client.Do(request)

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return responseBody, err
	}
	return responseBody, nil
}
