package ProcessResult

import (
	ProcessData "ParallelArraySummaryGo/internal/processData"
)

func calculateTotal(currentTotal float64, finalResult *ProcessData.Result) {
	finalResult.Total += currentTotal
}

func calculateGroupTotal(totalsByGroup *map[int]float64, result *ProcessData.Result) {
	for group, total := range *totalsByGroup {
		result.TotalsByGroup[group] += total
	}
}

func mergeIdsWithValuesLessThanFive(idsWithValuesLessThanFive []int, finalResult *ProcessData.Result) {
	finalResult.IdsWithValuesLessThanFive = append(finalResult.IdsWithValuesLessThanFive, idsWithValuesLessThanFive...)
}

func mergeIdsWithValuesGreaterThanFive(idsWithValuesGreaterThanFive []int, finalResult *ProcessData.Result) {
	finalResult.IdsWithValuesGreaterThanFive = append(finalResult.IdsWithValuesGreaterThanFive, idsWithValuesGreaterThanFive...)
}

func ProcessResults(resultsChannel chan ProcessData.Result) ProcessData.Result {
	finalResult := ProcessData.Result{
		TotalsByGroup: make(map[int]float64),
	}
	for result := range resultsChannel {
		calculateTotal(result.Total, &finalResult)
		calculateGroupTotal(&result.TotalsByGroup, &finalResult)
		mergeIdsWithValuesLessThanFive(result.IdsWithValuesLessThanFive, &finalResult)
		mergeIdsWithValuesGreaterThanFive(result.IdsWithValuesGreaterThanFive, &finalResult)
	}
	return finalResult
}
