package main

import (
	DisplayResults "ParallelArraySummaryGo/internal/displayResults"
	LoadData "ParallelArraySummaryGo/internal/loadData"
	ProcessData "ParallelArraySummaryGo/internal/processData"
	ProcessResult "ParallelArraySummaryGo/internal/resultsProcess"
	Utils "ParallelArraySummaryGo/internal/utils"
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
	//DisplayResults.DisplayIdsWithValuesLessThanFive(finalResult)
	//DisplayResults.DisplayIdsWithValuesGreaterThanFive(finalResult)
}
