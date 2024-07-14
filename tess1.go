package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var ar []string
	scanner := bufio.NewScanner(os.Stdin)

	// อ่านค่า input จนกว่า EOF จะถูกพบ
	for scanner.Scan() {
		ar = append(ar, scanner.Text())
	}

	if len(ar) < 2 {
		fmt.Println("Not enough input")
		return
	}

	a := ar[0]
	b := ar[1]
	ar = ar[2:]
	n := len(ar)

	for i := 0; i < n; i++ {
		poit := strings.Fields(ar[i])
		if len(poit) < 3 {
			continue
		}
		start := poit[0]
		end := poit[1]
		result := poit[2]

		startInt := toInt(start)
		endInt := toInt(end)
		aInt := toInt(a)

		if startInt <= aInt && aInt <= endInt {
			fmt.Println(result)
		}
	}
}

func toInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}
