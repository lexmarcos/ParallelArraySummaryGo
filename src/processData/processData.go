package ProcessData

import (
	LoadData "ParallelArraySummaryGo/src/loadData"
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Total                        float64
	TotalsByGroup                map[int]float64
	ItemsWithValuesLessThanFive    int
	ItemsWithValuesGreaterThanFive int
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

func appendToItemsWithValuesLessThanFive(id int, result *Result, total float64) {
	if total < 5 {
		result.ItemsWithValuesLessThanFive++
	}
}

func appendToItemsWithValuesGreaterThanFive(id int, result *Result, total float64) {
	if total >= 5 {
		result.ItemsWithValuesGreaterThanFive++
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
		appendToItemsWithValuesLessThanFive(item.ID, &result, item.Total)
		appendToItemsWithValuesGreaterThanFive(item.ID, &result, item.Total)
	}
	//fmt.Printf("Thread %s Total: %.4f 💰\n", name, result.Total)
	resultsChannel <- result
	defer waitGroup.Done()
}
