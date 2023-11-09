package ProcessData

import (
	LoadData "ParallelArraySummaryGo/internal/loadData"
	"fmt"
)

type Result struct {
	Total                        float64
	TotalByGroup                 map[int]float64
	IdsWithValuesLessThanFive    []int
	IdsWithValuesGreaterThanFive []int
}

func ProcessData(partitions [][]LoadData.Item, sumChannel chan Result) {
	fmt.Println("Starting threads")
	for i, partition := range partitions {
		fmt.Println("Thread ", i, "started")
		go ItemProcessor(partition, fmt.Sprintf("T%d", i), sumChannel)
	}
}

func sumToGroup(groupId int, valueToSum float64, currentResult Result) {
	if currentResult.TotalByGroup != nil {
		currentResult.TotalByGroup[groupId] += valueToSum
	}
	currentResult.TotalByGroup = make(map[int]float64)
}

func ItemProcessor(items []LoadData.Item, name string, sumChannel chan Result) {
	var currentResult Result
	for _, item := range items {
		currentResult.Total += item.Total
		sumToGroup(item.Group, item.Total, currentResult)
		if item.Total < 5 {
			currentResult.IdsWithValuesLessThanFive = append(currentResult.IdsWithValuesLessThanFive, item.ID)
		} else {
			currentResult.IdsWithValuesGreaterThanFive = append(currentResult.IdsWithValuesGreaterThanFive, item.ID)
		}
	}
	fmt.Println("Thread ", name, "Total: ", currentResult.Total)
	sumChannel <- currentResult
}
