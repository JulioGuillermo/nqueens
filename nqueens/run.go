package nqueens

import "fmt"

func RunGANQ(N, popsize, survivors int, mutrate float64, mc bool, threads int) {
	ga := NewGA(N, popsize, survivors, mutrate, mc, threads)

	fmt.Println(ga.Info())
	for ga.GetBestError() > 0 {
		ga.NextGen()
		fmt.Println(ga.Info())
	}

	ga.Save()
}
