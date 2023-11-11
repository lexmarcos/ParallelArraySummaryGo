package DisplayResults

import (
	ProcessData "ParallelArraySummaryGo/internal/processData"
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

func DisplayIdsWithValuesLessThanFive(result ProcessData.Result) {
	fmt.Println("IDs with Values Less Than 5:")
	for _, id := range result.IdsWithValuesLessThanFive {
		fmt.Printf("%d ", id)
	}
	fmt.Println()
}

func DisplayIdsWithValuesGreaterThanFive(result ProcessData.Result) {
	fmt.Println("\nIDs with Values Greater Than 5:")
	for _, id := range result.IdsWithValuesGreaterThanFive {
		fmt.Printf("%d ", id)
	}
	fmt.Println()
}
