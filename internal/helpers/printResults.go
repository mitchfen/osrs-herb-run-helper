package helpers

import (
	"fmt"
)

func PrintResults(herbs []Herb) {
	fmt.Printf("If you print %v seeds in %v patches, and assume %v herbs per patch. Then the expected profits are as follows:\n",
		numberofPatches, numberofPatches, numberofHerbsPerSeed)
	for i := 0; i < len(herbs); i++ {
		//fmt.Printf("%v herb price: %v\n", herbs[i].Name, herbs[i].HerbPrice)
		//fmt.Printf("%v seed price: %v\n", herbs[i].Name, herbs[i].SeedPrice)
		fmt.Printf("%v: %d\n", herbs[i].Name, int(herbs[i].ExpectedProfit))
	}
}
