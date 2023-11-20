package DisplayResults

import (
	ProcessData "ParallelArraySummaryGo/src/processData"
	"fmt"
)

func DisplayTotal(result ProcessData.Result) {
	fmt.Printf("\nTotal after proccess: %.4f 💰\n\n", result.Total)
}

func DisplayTotalsByGroup(result ProcessData.Result) {
	fmt.Println("\nTotals by Group:")
	for group, total := range result.TotalsByGroup {
		fmt.Printf("Group %d: %.4f 📊\n", group, total)
	}
	fmt.Println()
}

func DisplayItemsWithValuesLessThanFive(result ProcessData.Result) {
	fmt.Println("IDs with Values Less Than 5: ", result.ItemsWithValuesLessThanFive)
	fmt.Println()
}

func DisplayItemsWithValuesGreaterThanFive(result ProcessData.Result) {
	fmt.Println("\nIDs with Values Greater Than 5: ", result.ItemsWithValuesGreaterThanFive)
	fmt.Println()
}
