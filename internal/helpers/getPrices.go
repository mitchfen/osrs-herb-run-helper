package helpers

import (
	"io"
	"net/http"

	"github.com/Jeffail/gabs/v2"
)

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

func GetParsedPriceJson() *gabs.Container {
	priceJson, err := getPriceJson()
	if err != nil {
		panic(err)
	}

	parsedJson, err := gabs.ParseJSON(priceJson)
	if err != nil {
		panic(err)
	}

	return parsedJson
}
