package ProcessResult

import (
	ProcessData "ParallelArraySummaryGo/src/processData"
)

func calculateTotal(currentTotal float64, finalResult *ProcessData.Result) {
	finalResult.Total += currentTotal
}

func calculateGroupTotal(totalsByGroup *map[int]float64, result *ProcessData.Result) {
	for group, total := range *totalsByGroup {
		result.TotalsByGroup[group] += total
	}
}

func mergeItemsWithValuesLessThanFive(finalResult *ProcessData.Result) {
	finalResult.ItemsWithValuesLessThanFive++
}

func mergeItemsWithValuesGreaterThanFive(finalResult *ProcessData.Result) {
	finalResult.ItemsWithValuesGreaterThanFive++
}

func ProcessResults(resultsChannel chan ProcessData.Result) ProcessData.Result {
	finalResult := ProcessData.Result{
		TotalsByGroup: make(map[int]float64),
	}
	for result := range resultsChannel {
		calculateTotal(result.Total, &finalResult)
		calculateGroupTotal(&result.TotalsByGroup, &finalResult)
		mergeItemsWithValuesLessThanFive(&finalResult)
		mergeItemsWithValuesGreaterThanFive(&finalResult)
	}
	return finalResult
}
