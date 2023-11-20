package main

import (
	DisplayResults "ParallelArraySummaryGo/src/displayResults"
	LoadData "ParallelArraySummaryGo/src/loadData"
	ProcessData "ParallelArraySummaryGo/src/processData"
	ProcessResult "ParallelArraySummaryGo/src/resultsProcess"
	Utils "ParallelArraySummaryGo/src/utils"
	"sync"
)

func main() {
	var N, T int
	var waitGroup sync.WaitGroup

	Utils.LoadVariables(&N, &T)

	partitions := LoadData.PartitionList(LoadData.LoadData(N), T)
	resultsChannel := make(chan ProcessData.Result, T)
	ProcessData.ProcessData(partitions, resultsChannel, &waitGroup)

	finalResult := ProcessResult.ProcessResults(resultsChannel)

	DisplayResults.DisplayTotal(finalResult)
	DisplayResults.DisplayTotalsByGroup(finalResult)
	//DisplayResults.DisplayItemsWithValuesLessThanFive(finalResult)
	//DisplayResults.DisplayItemsWithValuesGreaterThanFive(finalResult)
}
