package ProcessData

import (
	LoadData "ParallelArraySummaryGo/internal/loadData"
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Total                        float64
	TotalsByGroup                map[int]float64
	IdsWithValuesLessThanFive    []int
	IdsWithValuesGreaterThanFive []int
}

func ProcessData(partitions *[]*[]LoadData.Item, resultsChannel chan<- Result, waitGroup *sync.WaitGroup) {
	fmt.Println("Starting threads 🧵")
	startTime := time.Now()

	for i := 0; i < len(*partitions); i++ {
		waitGroup.Add(1)
		go ItemProcessor((*partitions)[i], fmt.Sprintf("T%d", i), resultsChannel, waitGroup)
	}
	waitGroup.Wait()
	close(resultsChannel)
	duration := time.Since(startTime)
	fmt.Printf("Function execution time: %v ⏱️\n", duration)
}

func appendToIdsWithValuesLessThanFive(id int, result *Result, total float64) {
	if total < 5 {
		result.IdsWithValuesLessThanFive = append(result.IdsWithValuesLessThanFive, id)
	}
}

func appendToIdsWithValuesGreaterThanFive(id int, result *Result, total float64) {
	if total >= 5 {
		result.IdsWithValuesGreaterThanFive = append(result.IdsWithValuesGreaterThanFive, id)
	}
}

func sumToGroup(groupId int, valueToSum float64, result *Result) {
	result.TotalsByGroup[groupId] += valueToSum
}

func addToTotal(total float64, result *Result) {
	result.Total += total
}

func ItemProcessor(items *[]LoadData.Item, name string, resultsChannel chan<- Result, waitGroup *sync.WaitGroup) {
	result := Result{
		TotalsByGroup: make(map[int]float64),
	}
	for _, item := range *items {
		addToTotal(item.Total, &result)
		sumToGroup(item.Group, item.Total, &result)
		appendToIdsWithValuesLessThanFive(item.ID, &result, item.Total)
		appendToIdsWithValuesGreaterThanFive(item.ID, &result, item.Total)
	}
	//fmt.Printf("Thread %s Total: %.4f 💰\n", name, result.Total)
	resultsChannel <- result
	defer waitGroup.Done()
}
