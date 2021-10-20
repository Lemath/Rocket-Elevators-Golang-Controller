package main

import (
	"os"
	"strconv"
)

func main() {
    scenarioNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		runScenario(scenarioNumber)
	}
}

func Contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
