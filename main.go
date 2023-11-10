package main

import (
	LoadData "ParallelArraySummaryGo/internal/loadData"
	ProcessData "ParallelArraySummaryGo/internal/processData"
	"fmt"
	"sync"
)

func main() {
	N := 2
	T := 4
	var waitGroup sync.WaitGroup

	partitions := LoadData.PartitionList(LoadData.LoadData(N), T)
	resultsChannel := make(chan ProcessData.Result, T)
	ProcessData.ProcessData(partitions, resultsChannel, &waitGroup)
	total := 0.0
	for result := range resultsChannel {
		total += result.Total
	}
	fmt.Println("Total: ", total)
}
