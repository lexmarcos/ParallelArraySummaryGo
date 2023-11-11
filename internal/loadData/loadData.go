package LoadData

import (
	"fmt"
	"math"
	"math/rand"
)

type Item struct {
	ID    int
	Total float64
	Group int
}

func truncateNumber(num float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Trunc(num*factor) / factor
}

func calculateExpectedTotal(total float64, expectedTotal *float64) {
	*expectedTotal += total
}

func LoadData(N int) *[]Item {
	var expectedTotal float64
	itemCount := int(math.Pow(10, float64(N)))
	itemList := make([]Item, 0, itemCount)

	for i := 0; i < itemCount; i++ {
		total := truncateNumber(rand.Float64()*10, 4)
		calculateExpectedTotal(total, &expectedTotal)
		group := rand.Intn(5) + 1
		itemList = append(itemList, Item{i, total, group})
	}

	fmt.Printf("\nExpected total: %.4f ⚖️\n\n", expectedTotal)
	return &itemList
}

func PartitionList(items *[]Item, numberOfThreads int) *[]*[]Item {
	sizeOfList := len(*items)
	partSize := sizeOfList / numberOfThreads
	remainder := sizeOfList % numberOfThreads
	partitions := make([]*[]Item, 0, numberOfThreads)

	for i := 0; i < sizeOfList; {
		end := i + partSize
		if remainder > 0 {
			end++
			remainder--
		}
		if end > sizeOfList {
			end = sizeOfList
		}
		partition := (*items)[i:end]
		partitions = append(partitions, &partition)
		i = end
	}
	return &partitions
}
