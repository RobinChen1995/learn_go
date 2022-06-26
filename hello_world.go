package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	highlow(2, 0)
}

func highlow(high int, low int) {
	if high < low {
		fmt.Println("panic!!!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("deferred: highlow(", high, ",", low, ")")
	highlow(high, low+1)
}

func test3() {
	rand.Seed(time.Now().Unix())
	r := rand.Float64()
	switch {
	case r > 0.1:
		fmt.Println("common case")
	default:
		fmt.Println("10 percent of time. ")
	}
}
func test1() {
	rune := 'G'
	fmt.Println(rune)
}

func test2() {
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}

func sum(n1 string, n2 string) int {
	int1, _ := strconv.Atoi(n1)
	int2, _ := strconv.Atoi(n2)
	//fmt.Println(os.Args[0])
	return int1 + int2
}
