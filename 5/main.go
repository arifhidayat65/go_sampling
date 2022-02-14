package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	m := make(map[string]int)
	steps := 7
	m["A"] = 0
	m["B"] = 0
	m["C"] = 0
	m["D"] = 0
	m["E"] = 0

	for i := 1; i <= steps; i++ {
		for i := range m {
			m[i] = 0
		}
		insertRandomObj(m)
	}
	fmt.Println(m)
	objectInMap := findObj(m)
	fmt.Println(objectInMap)
}

func findObj(m map[string]int) string {
	keys := []string{}
	for k := range m {
		if m[k] != 0 {
			keys = append(keys, k)
		}
	}
	return keys[0]
}

func insertRandomObj(m map[string]int) {
	rand.Seed(time.Now().UnixNano())
	randINT := rand.Intn(len(m))
	switch randINT {
	case 0:
		m["A"] = 1
	case 1:
		m["B"] = 1
	case 2:
		m["C"] = 1
	case 3:
		m["D"] = 1
	default:
		m["E"] = 1
	}
}
