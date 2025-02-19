package helpers

import (
	"fmt"
)

func PrintResults(herbs []Herb) {
	fmt.Printf("Assuming you plant %v seeds in %v patches, and assuming %v herbs per patch.\nThen the expected profits are as follows:\n",
		numberofPatches, numberofPatches, numberofHerbsPerSeed)
	for i := 0; i < len(herbs); i++ {
		fmt.Printf("%v: %dk\n", herbs[i].Name, int(herbs[i].ExpectedProfit/1000.0))
	}
}
