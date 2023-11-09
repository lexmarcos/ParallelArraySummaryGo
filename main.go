package main

import (
	LoadData "ParallelArraySummaryGo/internal/loadData"
	ProcessData "ParallelArraySummaryGo/internal/processData"
)

func main() {
	N := 2
	T := 4
	partitions := LoadData.PartitionList(LoadData.LoadData(N), T)
	channel := make(chan ProcessData.Result)
	ProcessData.ProcessData(partitions, channel)

}
