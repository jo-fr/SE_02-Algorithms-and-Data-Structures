package main

import (
	"fmt"
	"math/rand"

	ds "github.com/jo-fr/SE_02-Algorithms-and-Data-Structures/examples/data_structures"
)

func main() {

	fmt.Println("############# LINKED LIST EXAMPLE #############")

	var nodes []*ds.Node

	for i := 0; i < 1000; i++ {
		nodes = append(nodes, ds.NewNode(rand.Intn(100)))

		if i != 0 {
			nodes[i-1].Next(nodes[i])

		}

	}

	ds.ShowCompleteList(nodes[0])

}
