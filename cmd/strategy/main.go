package main

import (
	"fmt"
	"time"

	"main/internal/strategy"
)

func generator(lenArray int) []string {
	arr := make([]string, lenArray)

	for i := 0; i < lenArray; i++ {
		arr[i] = "ABCDEFG abcdefg"
	}

	return arr
}

func main() {
	lenArr := 100000000
	keyWord := "My word"
	arr := generator(lenArr)

	arr[(lenArr-1)-10000] = keyWord

	finder := strategy.NewFinder()
	multiFinder := strategy.NewMultiFinder()
	strategyRelease := strategy.NewStrategyRelease()

	start := time.Now()
	strategyRelease.SetStrategy(finder)
	result := strategyRelease.Search(arr[:], keyWord)
	elapsed := time.Now().Sub(start).Milliseconds()
	fmt.Println(fmt.Sprintf("Found: %t; Elapsed time: %vms", result, elapsed))

	start = time.Now()
	strategyRelease.SetStrategy(multiFinder)
	result = strategyRelease.Search(arr[:], keyWord)
	elapsed = time.Now().Sub(start).Milliseconds()
	fmt.Println(fmt.Sprintf("Found: %t; Elapsed time: %dms", result, elapsed))
}
