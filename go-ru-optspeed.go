package main

import (
	"fmt"
	"time"
)

func __pow(basic *float64, pow uint64, numthread uint64, chanel chan float64, printresult bool) {

	var result float64 = 0

	switch {
	case pow == 0:
		result = 1
	case pow == 1:
		result = *basic
	case pow > 1:

		var i uint64

		i = 2
		result = *basic

		for i <= pow {
			result *= *basic
			i++
		}

	}

	if printresult {
		fmt.Println("Basic:", *basic, "Pow:", pow, "Result:", result, ", Number thread:", numthread)
	}

	chanel <- result

}

func main() {

	var iVariant int = 1 // true(!=0) - slice(w/o delete chanel), false(==0) - slice(with delete chanel)

	BeginTime := time.Now()

	var (
		ChThreads   []chan float64
		basic       float64 = 2
		printresult         = false
		StrResult   string
		Result      = float64(0)
		Pow_Thread  = uint64(0)
		MaxPow      = uint64(100000)
		Iterator    = uint64(0)
	)

	for Pow_Thread <= MaxPow {

		ChThreads = append(ChThreads, make(chan float64))

		go __pow(&basic, Pow_Thread, Pow_Thread+1, ChThreads[Pow_Thread], printresult)

		Pow_Thread++
	}

	CountIteration := Pow_Thread

	Pow_Thread--

	fmt.Println("All threads runing.")

	if iVariant != 0 {

		DoneChThreads := uint64(0)

		for DoneChThreads != CountIteration {

			select {
			case TempResult := <-ChThreads[Iterator]:
				Result += TempResult
				DoneChThreads++
			default:
			}

			Iterator++
			if Iterator > Pow_Thread {
				Iterator = 0
			}
		}

		StrResult = "Result(sliceND):"

	} else {

		for len(ChThreads) > 0 {

			select {
			case TempResult := <-ChThreads[Iterator]:
				Result += TempResult
				ChThreads = append(ChThreads[:Iterator], ChThreads[Iterator+1:]...)
			default:
				Iterator++
			}

			if Iterator >= uint64(len(ChThreads)) {
				Iterator = 0
			}

		}

		StrResult = "Result(sliceWD):"
	}

	EndTime := time.Now()

	fmt.Println("----------------------------------\n", "MaxPow", MaxPow, StrResult, Result, "Timer:", EndTime.Sub(BeginTime), "\n----------------------------------")

}
