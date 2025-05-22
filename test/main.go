package main

import "log"

type Numbers interface {
	int64 | float64 | float32
}

func main() {
	a := []int64{1, 2, 3, 4}
	b := []float64{1.1, 2.2, 3.3, 4.4}

	log.Println(a)
	log.Println(b)

}

func sumV[v Numbers](a []v) v {
	var sum v
	for _, val := range a {
		sum += val
	}
	return sum
}
