package main

import "fmt"

type t struct {
	money int
	num   int
}

func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	ma := map[string]t{}
	for i := 0; i < n; i++ {
		var name string
		fmt.Scan(&name)
		a, b := 0, 0
		fmt.Scan(&a, &b)
		ma[name] = t{money: a, num: b}
	}
	sum := 0
	i := 0
	for i = 0; i < m; i++ {
		var name string
		fmt.Scan(&name)
		x := 0
		fmt.Scan(&x)
		if x <= ma[name].num {
			sum += ma[name].money * x

		} else {
			fmt.Println(-(i + 1))
			break
		}
	}
	if i == m {
		fmt.Println(sum)
	}

}
